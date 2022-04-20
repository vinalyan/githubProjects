package event_consumer

import (
	"log"
	"read-adviser-bot/events"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

// работает в цикле
// если ошибка то бросает ее в лог
// если эвентов нет, что через секунду продолжает работать
// если
func (c Consumer) Start() error {
	for {
		gotEvent, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consmer: %s", err.Error())
			continue
		}

		if len(gotEvent) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}
		if err := c.hadleEvents(gotEvent); err != nil {
			log.Printf(err.Error())

			continue
		}
	}
}

/*
Проблемы с фукнцией ниже
1. потеря событий: ретраи, возращение в хранилише, фолбэкб, подтверждение для фетчера,
2. ОБработка всей пачки. Остановка поле ошибки,
3. Праллельная обработка.
TODO посмотреть функцию waitGroup
*/
func (c *Consumer) hadleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("получил новыое событие %s", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("упал тут hadleEvents: %s", err.Error())

			continue
		}
	}
	return nil
}
