package entity

import (
	"time"
)

type Article struct {
	ID         int64
	Title      string
	Slug       string
	CategoryID int64
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
