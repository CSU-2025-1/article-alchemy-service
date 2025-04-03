package models

type YandexResponse struct {
	Status     string `json:"status"`
	SharingURL string `json:"sharing_url"`
}

type SummaryResponse struct {
	MainTitle string    `json:"main_title"`
	Data      []Chapter `json:"data"`
}

type Chapter struct {
	Title  string   `json:"title"`
	Points []string `json:"points"`
}
