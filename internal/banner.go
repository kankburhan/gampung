package runner

import "github.com/projectdiscovery/gologger"

func showInit()  {
	gologger.Print().Msgf("%s\n\n", banner)
}