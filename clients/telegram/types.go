package telegram

type UpdateResponse struct {
	Ok     bool     `json: "ok"`
	Result []Update `json: "result"`
}

type Update struct {
	ID      int              `json: "update_id"` // for the parser
	Message *IncomingMessage `json: "message"`   // pointer because message actually optional and can be omitted
}

type IncomingMessage struct {
	Text string `json:"text"`
	From From   `json:"from"`
	Chat Chat   `json:"chat"`
}

type Chat struct {
	ID int `json:"id"`
}

type From struct {
	Username string `json:"username"`
}
