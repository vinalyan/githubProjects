package telegram

import (
	"errors"
	"log"
	"net/url"
	"read-adviser-bot/storage"
	"strings"

	"read-adviser-bot/lib/e"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

//TODO: Клавиатуры вывести отдельно

const (
	replyMarkupKeyboard = `{ "keyboard": [ [{"text": "/rnd"}], [{ "text": "/start"},{ "text": "/help"}] ], "one_time_keyboard": true}`
	NoReplyMarkup       = ""
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("Получена новая команда %s от %s", text, username)

	//проверяем, что это комнада добавления
	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}
	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand, NoReplyMarkup)

	}
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func (p *Processor) savePage(chatID int, pageURL string, username string) (err error) {
	defer func() { err = e.WrapIfErr("не выполнил events/telegram/commands.savePage", err) }()
	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExist, err := p.storage.IsExist(page)
	if err != nil {
		return err
	}
	//если уже сохранено, то пишем сообщение
	if isExist {
		return p.tg.SendMessage(chatID, msgAlreadyExists, NoReplyMarkup)
	}
	//пытаемся сохранить страницу
	if err := p.storage.Save(page); err != nil { //TODO разобраться  с обработкой ошибок и этой записью
		return err
	}
	if err := p.tg.SendMessage(chatID, msgSaved, NoReplyMarkup); err != nil {
		return err
	}
	return nil
}

//получить рандомную ссылку
// rnd page: /rnd
func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() {
		err = e.WrapIfErr("sendRandom не вышло получить случайную ссылку", err)
	}()
	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	}
	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages, NoReplyMarkup)
	}
	if err := p.tg.SendMessage(chatID, page.URL, NoReplyMarkup); err != nil {
		return err
	}
	return p.storage.Remove(page)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp, NoReplyMarkup)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMessage(chatID, msgHello, replyMarkupKeyboard)
}

//проверяем является ли ссылка ссылкой
func isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != "" // TODO потыкать в парсилку урлов
}
