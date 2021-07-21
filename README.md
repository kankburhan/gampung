# Gampung

[![made-with-Go](https://img.shields.io/badge/made%20with-Go-blue.svg)](http://golang.org)
[![issues](https://img.shields.io/github/issues/kankburhan/gampung?color=blue)](https://github.com/kankburhan/gampung/issues)

```txt
░██████╗░░█████╗░███╗░░░███╗██████╗░██╗░░░██╗███╗░░██╗░██████╗░
██╔════╝░██╔══██╗████╗░████║██╔══██╗██║░░░██║████╗░██║██╔════╝░
██║░░██╗░███████║██╔████╔██║██████╔╝██║░░░██║██╔██╗██║██║░░██╗░
██║░░╚██╗██╔══██║██║╚██╔╝██║██╔═══╝░██║░░░██║██║╚████║██║░░╚██╗
╚██████╔╝██║░░██║██║░╚═╝░██║██║░░░░░╚██████╔╝██║░╚███║╚██████╔╝
░╚═════╝░╚═╝░░╚═╝╚═╝░░░░░╚═╝╚═╝░░░░░░╚═════╝░╚═╝░░╚══╝░╚═════╝░
______________________ by @kankburhan ________________________

Searching Template From github...
Found 100 Templates from github
[INF] Success Download tor-socks-proxy.yaml
```

Tools to find nuclei template from github.

---

## Resources

- [Installation](#installation)
    - [from Source](#from-source)
    - [from GitHub](#from-github)
- [Usage](#usage)
    - [Basic Usage](#basic-usage)
    - [Flags](#flags)
- [TODOs](#todos)
- [Help & Bugs](#help--bugs)
- [License](#license)
- [Version](#version)

## Installation

### from Source

go compiler:

```bash
▶ GO111MODULE=on go get github.com/kankburhan/gampung
```

### from GitHub

```bash
▶ git clone https://github.com/kankburhan/gampung
▶ cd gampung
▶ go build .
▶ (sudo) mv gampung /usr/local/bin
```

## Usage

### Basic Usage

Run gampung with:

```bash
▶ gampung -t "token github" -o "/path/output"
```

### Flags

```bash
▶ gampung -h
```

This will display help for the tool. 

```txt


░██████╗░░█████╗░███╗░░░███╗██████╗░██╗░░░██╗███╗░░██╗░██████╗░
██╔════╝░██╔══██╗████╗░████║██╔══██╗██║░░░██║████╗░██║██╔════╝░
██║░░██╗░███████║██╔████╔██║██████╔╝██║░░░██║██╔██╗██║██║░░██╗░
██║░░╚██╗██╔══██║██║╚██╔╝██║██╔═══╝░██║░░░██║██║╚████║██║░░╚██╗
╚██████╔╝██║░░██║██║░╚═╝░██║██║░░░░░╚██████╔╝██║░╚███║╚██████╔╝
░╚═════╝░╚═╝░░╚═╝╚═╝░░░░░╚═╝╚═╝░░░░░░╚═════╝░╚═╝░░╚══╝░╚═════╝░
______________________ by @kankburhan ________________________


Usage:
  gampung [options]

Options:
  -t, --token <Token>       User access token Github
  -c, --concurrent <i>      Set the concurrency level (default: 20)
  -q, --query <keyword>     Keyword search template default "extension:yaml  matchers-condition"
  -o, --output <Path>       Path output files
  -s, --silent              Silent mode
  -h, --help                Display its help


```
## TODOs

- [ ] rate limit github
- [ ] Remove duplicate templates
- [ ] Validate template
## Help & Bugs

[![contributions welcome](https://img.shields.io/badge/contributions-welcome-blue.svg)](https://github.com/kankburhan/gampung/issues)

If you are still confused or found a bug, please [open the issue](https://github.com/kankburhan/gampung/issues). 

## License

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

**gampung** released under MIT. See `LICENSE` for more details.

## Version

**Current version is 0.0.3** and still development.
