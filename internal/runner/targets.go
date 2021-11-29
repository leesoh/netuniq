package runner

import (
	"bufio"
	"net"
	"os"
	"strings"

	"github.com/leesoh/netuniq/internal/iputils"
)

func (r *Runner) IngestTargets() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		target := strings.TrimSpace(sc.Text())
		if target == "" {
			continue
		}
		r.Logger.Debugf("processing %s", target)
		if iputils.IsIP(target) {
			r.Logger.Debugf("%s is an IP", target)
			r.AddIP(target)
		} else if iputils.IsCIDR(target) {
			r.Logger.Debugf("%s is a CIDR", target)
			r.AddCIDR(target)
		} else {
			r.Logger.Debugf("%s must be a hostname", target)
			r.AddFQDN(target)
		}
	}
}

func (r *Runner) AddIP(target string) {
	ip := net.ParseIP(target)
	r.IP = append(r.IP, ip)
	r.Logger.Debugf("added IP %s", target)
}

func (r *Runner) AddCIDR(target string) {
	_, cidr, _ := net.ParseCIDR(target)
	r.CIDR = append(r.CIDR, cidr)
	r.Logger.Debugf("added CIDR %s", target)
}

func (r *Runner) AddFQDN(target string) {
	r.FQDN = append(r.FQDN, target)
	r.Logger.Debugf("added FQDN %s", target)
}
