package runner

import (
	"github.com/projectdiscovery/gologger"
	"os"
)

func (o *Options) validate() {
	if o.Token == "" {
		gologger.Fatal().Msgf("Github access token required")
		os.Exit(1)
	}
	if o.Saveto == "" {
		gologger.Fatal().Msgf("No output dir provided")
		os.Exit(1)
	}
	if err := os.MkdirAll(o.Saveto, 0750); err != nil {
		gologger.Fatal().Msgf("Failed to create output directory: %s", err.Error())
		os.Exit(1)
	}
}