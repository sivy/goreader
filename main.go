package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser(
		"goreader",
		"Goodreads client and mastodon reposter",
	)

	verbose := parser.Flag(
		"v", "verbose", &argparse.Options{
			Required: false, Help: "Verbose flag"})

	updates := parser.NewCommand(
		"updates", "Get updates from GoodReads")

	latest := updates.Flag(
		"l", "latest", &argparse.Options{
			Required: false,
			Help:     "Get only the latest updates (noop for now)",
		})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
	}

	if updates.Happened() {
		client := Client{}
		// get_updates()
	}
}
