package main

import (
	"flag"
	"fmt"
	"os"
)

type cliOptions struct {
	telegram bool
	discord  bool
	slack    bool
	version  bool
	stdin    bool
	message  string
}

func processArgs() cliOptions {

	opts := cliOptions{}
	flag.BoolVar(&opts.telegram, "telegram", false, "Send via telegram")
	flag.BoolVar(&opts.telegram, "t", false, "Send via telegram")
	flag.BoolVar(&opts.slack, "slack", false, "Send via slack")
	flag.BoolVar(&opts.slack, "s", false, "Send via slack")
	flag.BoolVar(&opts.version, "version", false, "Show version number")
	flag.BoolVar(&opts.version, "v", false, "Show version number")
	flag.BoolVar(&opts.stdin, "stdin", false, "Take input from stdin")
	flag.BoolVar(&opts.stdin, "si", false, "Take input from stdin")
	flag.StringVar(&opts.message, "message", "", "The message you want to send")
	flag.StringVar(&opts.message, "m", "", "The message you want to send")
	flag.Parse()

	return opts

}

func init() {
	flag.Usage = func() {
		h := "\nSend data through chat channels. Made by @dubs3c.\n\n"

		h += "Usage:\n"
		h += "  emissary [channel] [message]\n\n"

		h += "Options:\n"
		h += "  -s,  --slack        Send via Slack\n"
		h += "  -t,  --telegram     Send via Telegram\n"
		h += "  -si, --stdin        Get message from stdin\n"
		h += "  -m,  --message      Message to send\n"
		h += "  -v,  --version      Show version\n"

		h += "\nExamples:\n"
		h += "  emissary -telegram --message \"Hello telegram\"\n"
		h += "  cat domins.txt | emissary --slack --stdin \n\n"

		fmt.Fprintf(os.Stderr, h)
	}
}