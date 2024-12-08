package models

import (
	"time"
)

// ArticleStatus represents the publication status of an article
type ArticleStatus string

const (
	StatusDraft     ArticleStatus = "draft"
	StatusPublished ArticleStatus = "published"
	StatusArchived  ArticleStatus = "archived"
)

type Article struct {
	ID          int64         `json:"id"`
	AuthorId    int64         `json:"author_id"`
	Title       string        `json:"title"`
	Content     string        `json:"content"`
	Slug        string        `json:"slug"`
	PublishedAt *time.Time    `json:"published_at"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Status      ArticleStatus `json:"status"`
	ViewCount   int64         `json:"view_count"`
	Author      *User         `json:"author"`
}
