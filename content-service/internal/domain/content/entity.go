package content

import "time"

const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
)

type Content struct {
	ContentID uint64
	UserID    uint64
	Status    string
	Data      string
	Error     string
	CreatedAt time.Time
}
