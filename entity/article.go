package entity

import (
	"time"
)

// Article is an entity.
type Article struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    int64     `json:"userId"`
}
