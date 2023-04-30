package main

import (
	"flag"
	"log"
	event_consumer "reader-adviser-bot/consumer/event-consumer"
	"reader-adviser-bot/events/telegram"
	"reader-adviser-bot/storage/files"
	tgClient "reader-adviser/clients/telegram"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

// 6098437192:AAES4UcLv12IioNcHlCQ4urEJI7HSmYJUTw

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
