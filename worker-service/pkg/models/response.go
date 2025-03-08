package models

type YandexResponse struct {
	Status     string `json:"status"`
	SharingURL string `json:"sharing_url"`
}

type SummaryResponse struct {
	Summary string `json:"summary"`
}
