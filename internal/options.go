package runner

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Options args
type Options struct {
	Token, Saveto, Query string
	Concurrency int
	Silent, Verbose bool
}

var o *Options

func init(){
	o = &Options{}

	flag.StringVar(&o.Token, "t", "", "")
	flag.StringVar(&o.Token, "token", "", "")

	flag.StringVar(&o.Saveto, "o", "", "")
	flag.StringVar(&o.Saveto, "output", "", "")

	flag.StringVar(&o.Query, "q", "extension:yaml  matchers-condition", "")
	flag.StringVar(&o.Query, "query", "extension:yaml  matchers-condition", "")

	flag.IntVar(&o.Concurrency, "c", 25, "")
	flag.IntVar(&o.Concurrency, "concurrent", 25, "")

	flag.BoolVar(&o.Silent, "s", false, "")
	flag.BoolVar(&o.Silent, "silent", false, "")

	flag.BoolVar(&o.Verbose, "v", false, "")
	flag.BoolVar(&o.Verbose, "Verbose", false, "")

	flag.Usage = func() {
		showInit()
		h := []string{
			"",
			"Usage:" + usage,
			"",
			"Options:" + options,
			"",
		}

		fmt.Fprint(os.Stderr, strings.Join(h, "\n"))
	}
	flag.Parse()
}

func ParseOptions() *Options {
	if !o.Silent {
		showInit()
	}

	o.validate()

	return o
}