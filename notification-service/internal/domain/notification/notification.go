package notification

type EmailNotification struct {
    Email string          `json:"email"`
    URL   string          `json:"url"`
    Data  []Chapter       `json:"data"`
}

type Chapter struct {
    Title  string         `json:"title"`
    Points []string       `json:"points"`
}