package model

import "time"

const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
)

type Content struct {
	ContentID uint64    `db:"content_id"`
	UserID    uint64    `db:"user_id"`
	Status    string    `db:"status"`
	Data      string    `db:"data"`
	Error     string    `db:"error"`
	CreatedAt time.Time `db:"created_at"`
}

type ExtractContentDTO struct {
	UserID uint64
	Email  string
	Url    string
}

type ExtractContentPreviewDTO struct {
	Email string
	Url   string
}
