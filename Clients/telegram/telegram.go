package telegram

import (
	"encoding/json"
	"fmt"
	"go/constant"
	"go/format"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"path"
	
	"read-adviser-bot/lib/e"

)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod = "getUpdates"
	sendMessageMethod = "sendMessage"

)

func New(host string, token string) Client {
	return Client{
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
	q.Add(key: "offset", strconv.Itoa(offset))
	q.Add(key: "limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return updates: nil, err
	}

	//пасрим JSON
	var res UpdatesResponse 
	//TODO json.Unmarshal(data,&res) как это устроено? 
	if err := json.Unmarshal(data,&res); err != nil { 
		return updates: nil, err
	}
}

//отправляем сообщения
func (c *Client) SendMessage(chatId int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send message", err)
	}
	return nil
}

// отправка запроса

func (c *Client)  doRequest(method string, querry url.Values)(data []byte, err error) {

	//TODO разбираемся как устроены деструкторы
	defer func() {err = e.WrapIfErr(msg: "не могу выполнить запрос", err) }

	u := url.URL{
		Scheme: "https",
		Host: 	c.host,
		Path:	path.Join(c.basePath, method) //удобная приблуда склеивать путь, что бы со всякими слешами проблем не было.
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), body: nil)
	if err != nil {
		return nil, err
	}
	//TODO почитать про errors.Is() и errors.As()

	req.URL.RawQuerry = query.Encode()

	resp,err := c.client.Do(req)
	if err != nil {
		return nil, err
	}	
	defer func() {_=resp.Body.Close()}()

	body, err:=io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, err: nil
}



