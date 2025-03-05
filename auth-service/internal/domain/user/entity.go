package user

import "time"

type User struct {
	UserID    uint64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}
