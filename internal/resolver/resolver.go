package resolver

import (
	"net"

	"github.com/Masterminds/log-go/impl/cli"
	"github.com/projectdiscovery/dnsx/libs/dnsx"
)

type DNSClient struct {
	Client *dnsx.DNSX
	Logger *cli.Logger
}

func NewDNSClient(retries int, logger *cli.Logger) (*DNSClient, error) {
	options := dnsx.DefaultOptions
	options.MaxRetries = retries
	options.Hostsfile = true
	client, err := dnsx.New(options)
	if err != nil {
		return &DNSClient{}, err
	}
	dnsclient := &DNSClient{
		Client: client,
		Logger: logger,
	}
	return dnsclient, nil
}

func (d *DNSClient) Lookup(h string) ([]net.IP, error) {
	var targetIPs []net.IP
	ips, err := d.Client.Lookup(h)
	if err != nil {
		return targetIPs, err
	}
	for _, ii := range ips {
		ip := net.ParseIP(ii)
		if ip != nil {
			targetIPs = append(targetIPs, ip)
			d.Logger.Debugf("found ip %v\n", ii)
		}
	}
	return targetIPs, nil
}
