package model

import (
	"time"
)

type CategoryCreateRequest struct {
	CategoryName string `json:"category_name"`
}

type CategoryUpdateRequest struct {
	CategoryName string `json:"category_name"`
}

type CategoryResponse struct {
	ID           int64     `json:"id"`
	CategoryName string    `json:"category_name"`
	CategorySlug string    `json:"category_slug"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
