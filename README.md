# Signal Sciences CLI

CLI tool that interacts with the Signal Sciences API using the go-sigsci
library. Build with a simple `go build` or `make build`.

## Usage

    $ ./sigcli
    Interact with the Signal Sciences API

    Usage:
      sigcli [command]

    Available Commands:
      bash-completion     Generates bash completion scripts
      get-corp            Get info about a corp
      help                Help about any command
      list-agents         List sigsci agents
      list-corps          List sigsci corps
      list-suspicious-ips List suspicious IPS for a specific site
      list-top-attacks    List top attacks for a specific site

    Flags:
          --config string   config file (default is $HOME/.sigcli.yaml)
      -d, --debug           enable debug mode for verbose output
          --email string    SigSci email/username
      -h, --help            help for sigcli
          --token string    SigSci API token

    Use "sigcli [command] --help" for more information about a command.

## Config

The following order of precedence is used for obtaining config values:

- CLI flags
- Env vars
- Config file

All CLI commands have built-in help. Use the -h flag on any part of the CLI to
view usage docs.

You can set the following env vars:

- SIGSCI_EMAIL
- SIGSCI_TOKEN
- SIGSCI_CORP
- SIGSCI_SITE

You can also use a YAML or JSON config file:

    $ cat ~/.sigsci.yaml
    email: "YOUR_SIGSCI_EMAIL"
    token: "YOUR_SIGSCI_API_TOKEN"
    corp: "YOUR_CORP"
    site: "YOUR_SITE"
