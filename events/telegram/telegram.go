package telegram

import (
	"read-adviser-bot/lib/e"
	"read-adviser-bot/clients/telegram"
	"read-adviser-bot/events"
	"read-adviser-bot/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID int
	Username string
}

var (
	ErrUnknownEventType = error.New("неизвестный тип события")
	ErrUnknownMetaType = error.New("не известная мета")
)

func New(clien *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      clien,
		offset:  0,
		storage: storage,
	}

}

func (p *Processor) Process(event events.Event) error{
	switch event.Type {
	case events.Message:
		return p.processMeaasge(event)
	default:
		return Wrap(msg: "не разобрался в сообщении", ErrUnknownEventType)

	}

}

func (p *Processor) processMessage(event events.Event){
	meta, err:=meta(event)
	if err != nil {
		return e.Wrap(msg: "не вышло в processMessage")
	}
	if err := p.doCmd(event.Text, meta.ChatID, meta.Username); err!=nil(
		retun e.Wrap("не вышло в processMessage",err)
	)
	return nil
}

func meta(event events.Event)(Meta, error){
	res,ok := event.Meta.(Meta)   //TODO разобраться с проверкой типов данных
	if !ok {
		return Meta{}, e.Wrap("ошибка с определением меты", ErrUnknownMetaType)
	}
	return res, nil
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap(msg: "не могу получить событие", err)
	}

	if len(updatse) == 0 {
		return nil,nil
	}

	res := make([]events.Event,0,len(updates))

	for _, u := range updates {
		res = append(res, events(u))
	}
	p.offset = updates[len(updates)-1].ID + 1 //делаем смещение офсета

	return res, nil
}



func event (upd telegram.Update) events.Event {
	updType := fetchType(upd)


	res := events.Event{
		Type: updType,
		Text: fethchText(upd)
	}

	if updType==events.Message {
		res.Meta = Meta{
			ChatID: upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchType(upd telegram.Update) events.Type {
	if upd == nil {
		return events.Unknow
	}
	return events.Message 

}

func fethchText(upd telegram.Update) string {
	if upd == nil {
		return ""
	}
	return upd.Message.Text 
}