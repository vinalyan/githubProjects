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

//TODO Реализовать нормальную структуру ReplyMarkup что бы можно было в нее пихать разные типы клавиатур.
