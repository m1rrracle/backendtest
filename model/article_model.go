package model

import (
	"time"
)

type ArticleCreateRequest struct {
	Title      string `json:"title"`
	CategoryID int64  `json:"category_id"`
	Content    string `json:"content"`
}

type ArticleUpdateRequest struct {
	Title      string `json:"title"`
	CategoryID int64  `json:"category_id"`
	Content    string `json:"content"`
}

type ArticleResponse struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	Slug         string    `json:"slug"`
	CategoryID   int64     `json:"category_id"`
	CategoryName string    `json:"category_name"`
	CategorySlug string    `json:"category_slug"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}
