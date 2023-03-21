package main

import (
	"flag"
	"log"
	"reader-adviser/clients/telegram"
)

const (
	tgBotHost = "api.telegram.arg"
)

func main() {

	tgClient = telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor )
}

func mustToken() string {
	token := flag.String( //                   I DON'T KNOW THE REASON WHY USAGE OF FLAGS ARE PREFERABLE OVER RAW STRINGS
		"token-bot-token",                  // name token-bot-token
		"",                                 // value
		"token for access to telegram bot", // usage	->	Use -h or --help flags to get automatically generated help text for the command-line program.
	)

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
