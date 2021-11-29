package runner

import "flag"

type Options struct {
	Concurrency int
	JSON        bool
	Retries     int
	Verbose     bool
}

func ParseOptions() *Options {
	options := &Options{}
	flag.IntVar(&options.Concurrency, "c", 100, "Tool go brrrr")
	flag.IntVar(&options.Retries, "retries", 2, "Number of retries for name resolution")
	flag.BoolVar(&options.JSON, "json", false, "Display JSON output")
	flag.BoolVar(&options.Verbose, "verbose", false, "Display verbose output")
	flag.Parse()
	return options
}
