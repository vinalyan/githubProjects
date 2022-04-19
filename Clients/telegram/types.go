//тут должны быть все типы с которыми работает клиент

package telegram

type UpdatesResponse struct {
	Ok     bool     `json: "OK"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int    `json:"update_id"`
	Messege string `json:"messege"`
}
