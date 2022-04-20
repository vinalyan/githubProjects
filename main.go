package main

import (
	"flag"
	"go/token"
	"log"
	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/events"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/consumer/event-consumer"
)

const (
	tgBotHost = "api.telegram.org" //TODO вынести хост во флаг. Флаг это походу ключ
	storagePath = "storage"
	batchSize = 100
)

func main() {
	
	eventsProcessor := telegram.New(
					tgClient.New(tgBotHost, 
					mustToken()), 
					filles.New(storagePath))

	log.Printf("Cсервис запущен")

	consumer := event_consumer.New(
					eventsProcessor, 
					eventsProcessor, 
					batchSize);

	if err:= consumer.Start();err!=nil{
		log.Fatal("сервис дико остановился", err)
	}

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)


}

//TODO: надо не забыть разобраться, что такое приставка must
// мы обрабатываем токен

func	mustToken() string {
	
	token := flag.String(name: "tg-bot-token", value: "", usage: "тут должен быть токен" )

	flag .Parse()

	if *token == "" {
		log.Fatal(v..."токена нет ")
		
	}
	return *token

}
