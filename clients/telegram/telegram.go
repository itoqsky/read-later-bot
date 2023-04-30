package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"reader-adviser-bot/lib/e"
	"strconv"
)

const (
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

type Client struct {
	host     string
	basePath string      // API key
	client   http.Client // performs http.Client operations like Do, Post, Get, Head ...
}

func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {

	q := url.Values{} // preparing parameters for a request (preparing a query or URL)
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var res UpdateResponse

	if err := json.Unmarshal(data, &res); err != nil { // parses according to the properties of UpdateResponse
		return nil, err
	}

	return res.Result, err

}

func (c *Client) SendMessage(chatId int, text string) error {
	q := url.Values{}
	q.Add("chatId", strconv.Itoa(chatId)) // add key value query for url
	q.Add("text", text)

	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return e.Wrap("can't send a message", err)
	}

	return nil
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) { // we should specify return variables for defer, becuase defer does not know what variable do you mean
	defer func() { err = e.WrapIfErr("can't do request", err) }()

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil) // creating a GET request (METHOD, URL, body)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode() // prepare query by encoding into URL, click on Encode() method to see info

	resp, err := c.client.Do(req) // Requesting
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }() // defer MUST

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
