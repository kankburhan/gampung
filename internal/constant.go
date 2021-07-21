package runner

const (
	author  = "kankburhan"
	banner  = `

░██████╗░░█████╗░███╗░░░███╗██████╗░██╗░░░██╗███╗░░██╗░██████╗░
██╔════╝░██╔══██╗████╗░████║██╔══██╗██║░░░██║████╗░██║██╔════╝░
██║░░██╗░███████║██╔████╔██║██████╔╝██║░░░██║██╔██╗██║██║░░██╗░
██║░░╚██╗██╔══██║██║╚██╔╝██║██╔═══╝░██║░░░██║██║╚████║██║░░╚██╗
╚██████╔╝██║░░██║██║░╚═╝░██║██║░░░░░╚██████╔╝██║░╚███║╚██████╔╝
░╚═════╝░╚═╝░░╚═╝╚═╝░░░░░╚═╝╚═╝░░░░░░╚═════╝░╚═╝░░╚══╝░╚═════╝░
______________________ by @`+ author +` ________________________`
	usage = `
  gampung [options]`
	options = `
  -t, --token <Token>       User access token Github
  -c, --concurrent <i>      Set the concurrency level (default: 20)
  -q, --query <keyword>     Keyword search template default "extension:yaml  matchers-condition"
  -o, --output <Path>       Path output files
  -s, --silent              Silent mode
  -h, --help                Display its help
`
)