package main

import (
	"flag"
	"log"
	tgClient "read-adviser-bot/clients/telegram"
	event_consumer "read-adviser-bot/consumer/event-consumer"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(tgClient.New(tgBotHost, mustToken()), files.New(storagePath))

	log.Printf("Cервис запущен")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("Cервис дико остановился", err)
	}

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)

}

//TODO: надо разобраться как работает этот flag.String
// мы обрабатываем токен

func mustToken() string {

	token := flag.String("tg-bot-token", "", "тут должен быть токен")

	flag.Parse()

	if *token == "" {
		log.Fatal("токена нет ")

	}
	return *token

}
