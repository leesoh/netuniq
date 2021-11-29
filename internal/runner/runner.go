package runner

import (
	"net"
	"sync"

	"github.com/Masterminds/log-go"
	"github.com/Masterminds/log-go/impl/cli"
	"github.com/leesoh/netuniq/internal/resolver"
	"github.com/leesoh/netuniq/internal/result"
)

type Runner struct {
	Options   *Options
	Result    *result.Result
	Logger    *cli.Logger
	DNSClient *resolver.DNSClient
	FQDN      []string
	IP        []net.IP
	CIDR      []*net.IPNet
}

func NewRunner(options *Options) *Runner {
	logger := cli.NewStandard()
	res := result.NewResult(logger)
	dnsclient, err := resolver.NewDNSClient(options.Retries, logger)
	if err != nil {
		logger.Fatal(err)
	}
	if options.Verbose {
		logger.Level = log.DebugLevel
	} else {
		logger.Level = log.FatalLevel
	}
	runner := &Runner{
		Options:   options,
		Result:    res,
		Logger:    logger,
		DNSClient: dnsclient,
	}
	return runner
}

func (r *Runner) Run() {
	r.Logger.Debugf("processing targets with concurrency %d", r.Options.Concurrency)
	r.IngestTargets()
	jobs := make(chan Job, len(r.FQDN))
	results := make(chan Result, len(r.FQDN))
	var wg sync.WaitGroup
	// One worker for every level of concurrency
	for w := 1; w <= r.Options.Concurrency; w++ {
		wg.Add(1)
		worker := &Worker{
			DNSClient: r.DNSClient,
			Logger:    r.Logger,
		}
		go func(w int) {
			defer wg.Done()
			worker.Start(w, jobs)
		}(w)
	}
	// One job for every item in the queue
	for j := 0; j < len(r.FQDN); j++ {
		jobs <- Job{r.FQDN[j], results}
	}
	close(jobs)
	r.Logger.Debug("waiting to exit")
	wg.Wait()
	r.ProcessResults(results)
}

func (r *Runner) ProcessResults(results chan Result) {
	// All FQDNs should be added
	for i := 0; i < len(r.FQDN); i++ {
		record := <-results
		if record.Err != nil {
			r.Logger.Errorf("error during name resolution of %v: %v", record.Hostname, record.Err)
			continue
		}
		r.Result.AddHost(&result.Host{
			Name: record.Hostname,
			IPs:  record.IPs,
		})
	}
	// FQDN will not contain a CIDR, so we can add all CIDRs
	for _, cidr := range r.CIDR {
		r.Result.AddHost(&result.Host{CIDR: cidr})
		r.Logger.Debugf("added cidr %v to result", cidr)
	}
	// IPs can be contained in FQDN or CIDR, so we must check
	for _, ip := range r.IP {
		r.Result.AddIP(ip)
		r.Logger.Debugf("added ip %v to result", ip)
	}
	// Print results
	if r.Options.JSON {
		r.Result.PrintJSON()
	} else {
		r.Result.PrintTargets()
	}
}
