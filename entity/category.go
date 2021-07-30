package entity

import (
	"time"
)

type Category struct {
	ID           int64
	CategoryName string
	CategorySlug string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
