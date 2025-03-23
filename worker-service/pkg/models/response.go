package models

type YandexResponse struct {
	Status     string `json:"status"`
	SharingURL string `json:"sharing_url"`
}

type SummaryResponse struct {
	Summary []Chapter `json:"summary"`
}

type Chapter struct {
	Title  string   `json:"title"`
	Points []string `json:"points"`
}
