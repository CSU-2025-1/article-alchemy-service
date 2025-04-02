package model

type WorkerEvent struct {
	ContentID uint64 `json:"content_id"`
	URL       string `json:"url"`
	Email     string `json:"email"`
	Body      string `json:"body"`
	Error     string `json:"error"`
}
