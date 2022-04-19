package main

import (
	"flag"
	"go/token"
	"log"
)

const (
	tgBotHost = "api.telegram.org" //TODO вынести хост во флаг. Флаг это походу ключ
)

func main() {
	
	tgClient = telegram.New(tgBotHost, mustToken())

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
