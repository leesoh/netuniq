package main

import "github.com/leesoh/netuniq/internal/runner"

func main() {
	options := runner.ParseOptions()
	runner := runner.NewRunner(options)
	runner.Run()
}
