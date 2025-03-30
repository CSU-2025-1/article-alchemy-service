package models

type Event struct {
	UserID string `json:"user_id"`
	URL    string `json:"url"`
	Email  string `json:"email"`
}

type EventResponse struct {
	UserID string    `json:"user_id"`
	URL    string    `json:"url"`
	Email  string    `json:"email"`
	Data   []Chapter `json:"data"`
}
