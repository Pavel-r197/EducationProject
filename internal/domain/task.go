package domain

import "time"

type Task struct {
	Id          int64
	Title       string
	Description string
	UserId      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
