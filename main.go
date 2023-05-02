package main

import (
	"flag"
	"log"

	tgClient "github.com/itoqsky/reader-adviser-bot/clients/telegram"
	event_consumer "github.com/itoqsky/reader-adviser-bot/consumer/event-consumer"

	"github.com/itoqsky/reader-adviser-bot/events/telegram"
	"github.com/itoqsky/reader-adviser-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("... server started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
