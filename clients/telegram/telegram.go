package telegram

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"read-adviser-bot/lib/e"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

//получаем сообщения
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	//пасрим JSON
	var res UpdatesResponse

	//TODO json.Unmarshal(data,&res) как это устроено?
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	log.Printf("Updates 1 data:  %s", data)
	log.Printf("Updates 2 es.Ok:  %s", res.Ok)
	log.Printf("Updates 3 res.Result:  %s", res.Result)

	return res.Result, nil
}

//отправляем сообщения
func (c *Client) SendMessage(chatId int, text string, replymarkup ReplyMarkup) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", text)

	//TODO убрать от сюда эту дичь
	//q.Add("reply_markup", `{ "keyboard": [ [{"text": "/rnd"}], [{ "text": "/start"},{ "text": "/help"}] ] }`)

	q.Add("reply_markup", `{"inline_keyboard": [[{"text": "рандом","callback_data": "/rnd"}, {"text": "помощь","callback_data": "/help"}]]}`)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}
	return nil
}

// отправка запроса

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {

	defer func() { err = e.WrapIfErr("doRequest не могу выполнить запрос ", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	//TODO почитать про errors.Is() и errors.As()

	req.URL.RawQuery = query.Encode()

	//log.Printf("doRequest req = %s", req.URL)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
