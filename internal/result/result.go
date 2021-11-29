package result

import (
	"net"
	"sync"

	"github.com/Masterminds/log-go/impl/cli"
)

type Result struct {
	sync.RWMutex
	Hosts  []Host
	Logger *cli.Logger
}

type Host struct {
	IPs    []net.IP   `json:"ips,omitempty"`
	CIDR   *net.IPNet `json:"-"` // Subnet is more readable
	Subnet string     `json:"subnet,omitempty"`
	Name   string     `json:"hostname,omitempty"`
}

func (h *Host) HasIP(i net.IP) bool {
	for _, ii := range h.IPs {
		if ii.Equal(i) {
			return true
		}
	}
	return false
}

func NewResult(logger *cli.Logger) *Result {
	return &Result{Logger: logger}
}

func (r *Result) AddHost(h *Host) {
	r.Lock()
	defer r.Unlock()
	if r.ContainsHost(h) {
		return
	}
	// This bit of ugliness is so we can get nice CIDR outputs
	if h.CIDR != nil {
		h.Subnet = h.CIDR.String()
	}
	r.Hosts = append(r.Hosts, *h)
	r.Logger.Debugf("added host %v", h)
}

func (r *Result) ContainsHost(h *Host) bool {
	for _, hh := range r.Hosts {
		if hh.Name != "" && hh.Name == h.Name {
			return true
		}
	}
	return false
}

// AddIP adds a given IP to the result
func (r *Result) AddIP(i net.IP) {
	r.Lock()
	defer r.Unlock()
	for _, hh := range r.Hosts {
		// IP is in existing CIDR range, don't add
		if hh.CIDR != nil && hh.CIDR.Contains(i) {
			return
		}
		// Host contains IP, don't add
		if hh.HasIP(i) {
			return
		}
	}
	// IP is unique, add
	ips := []net.IP{i}
	r.Hosts = append(r.Hosts, Host{IPs: ips})
}
