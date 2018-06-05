# tt

tt is command line twitter client

## How to install

Fetch from GitHub and install.

```zsh
$ go get github.com/nasum/tt
$ go install
```

create `ttrc.json` in your `$HOME`

```json
{
  "CONSUMER_KEY":"your consumer key",
  "CONSUMER_SECRET":"your consumer secret",
  "ACCESS_TOKEN":"your access token",
  "ACCESS_SECRET":"your access secret"
}

```

## How to use

please see help

```zsh
$ tt help
Twitter Client

Usage:
  tt [command]

Available Commands:
  help        Help about any command
  timeline    get your timeline
  tweet       post your tweet

Flags:
  -h, --help   help for tt

Use "tt [command] --help" for more information about a command.

```
