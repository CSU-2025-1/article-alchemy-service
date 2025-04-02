package notification

import(
    "encoding/json"
)

type Chapter struct {
    Title  string      `json:"title"`
    Points []string    `json:"points"`
}

type EmailNotification struct {
    ContentID uint64            `json:"content_id"`
    URL       string            `json:"url"`
    Email     string            `json:"email"`
    Body      json.RawMessage   `json:"body"`
    Error     string            `json:"error"`
  }
  