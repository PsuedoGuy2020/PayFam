package entity

import (
	"time"
)

type Video struct {
	ID           uint      `json:"id"`
	VideoId      string    `json:"video_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	PublishedAt  time.Time `json:"published_at"`
	ThumbnailUrl string    `json:"thumbnail_url"`
}
