package user

import "time"

type User struct {
	UserID       uint64
	Username     string
	Email        string
	IsActive     bool
	PasswordHash string
	CreatedAt    time.Time
}
