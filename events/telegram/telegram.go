package telegram

import (
	"errors"
	"read-adviser-bot/clients/telegram"
	"read-adviser-bot/events"
	"read-adviser-bot/lib/e"
	"read-adviser-bot/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	Username string
}

var (
	ErrUnknownEventType = errors.New("неизвестный тип события")
	ErrUnknownMetaType  = errors.New("не известная мета")
)

func New(clien *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      clien,
		offset:  0,
		storage: storage,
	}

}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	default:
		return e.Wrap("Process не разобрался в сообщении", ErrUnknownEventType)

	}

}

func (p *Processor) processMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return e.Wrap("не разобрал мету meta events/telegram/telegram.processMessage", err)
	}
	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err != nil {
		return e.Wrap("не выполнил команду p.doCmd events/telegram/telegram.processMessage", err)
	}

	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta) //TODO разобраться с проверкой типов данных
	if !ok {
		return Meta{}, e.Wrap("ошибка с определением меты", ErrUnknownMetaType)
	}
	return res, nil
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("Fetch не могу получить событие", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}
	p.offset = updates[len(updates)-1].ID + 1 //делаем смещение офсета

	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type: updType,
		Text: fethchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknow
	}
	return events.Message

}

func fethchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}
