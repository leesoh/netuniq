package runner

import (
	"net"

	"github.com/Masterminds/log-go/impl/cli"
	"github.com/leesoh/netuniq/internal/resolver"
)

type Job struct {
	FQDN    string
	Results chan<- Result
}

type Result struct {
	Hostname string
	IPs      []net.IP
	Err      error
}

type Worker struct {
	DNSClient *resolver.DNSClient
	Logger    *cli.Logger
}

func (w *Worker) Start(id int, jobs <-chan Job) {
	for jj := range jobs {
		w.Logger.Debugf("worker %d processing %s", id, jj.FQDN)
		ips, err := w.DNSClient.Lookup(jj.FQDN)
		if err != nil {
			jj.Results <- Result{jj.FQDN, ips, err}
			w.Logger.Errorf("error during DNS lookup of %v: %v", jj.FQDN, err)
			continue
		}
		jj.Results <- Result{jj.FQDN, ips, nil}
	}
}
