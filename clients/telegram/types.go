//тут должны быть все типы с которыми работает клиент

package telegram

type UpdatesResponse struct {
	Ok     bool     `json: "OK"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

//TODO все это взято из описания API телеги
type IncomingMessage struct {
	Text string `json: "text"`
	From From   `json: "from"`
	Chat Chat   `json: "chat"`
}

type From struct {
	Username string `json: "username"`
}

type Chat struct {
	ID int `json: "id"`
}

// Тут основная идея в том, что есть тип ReplyMarkup в который мы будем запиливать разные данные
// ДЛя начала это буедт ReplyKeyboardMarkup

//TODO ReplyMarkup сделать интерфесом с возможностью получить JSON.
// reply_markup = { "keyboard": [ [{"text": "FIRST_BUTTON"}], [{ "text": "SECOND_BUTTON"}], [{ "text": "THIRD_BUTTON"}] ] }

type ReplyMarkup struct {
	//	Meta interface{}
	// TODO добавть интерфес.
}
