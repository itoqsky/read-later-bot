package telegram

type UpdateResponse struct {
	Ok     bool     `json: "ok"`
	Result []Update `json: "result"`
}

type Update struct {
	ID      int    `json: "update_id"` // for the parser
	Message string `json: "message"`
}
