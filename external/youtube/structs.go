package external

import (
	"PayFam/internal/models/repository"
	"time"
)

type VideoService struct {
	Repo           *repository.VideoRepository
	ApiKeys        []APIKey
	Query          string
	FetchInterval  time.Duration
	PublishedAfter time.Time
}

type YouTubeConfig struct {
	APIKeys []APIKey
}

type Config struct {
	YouTube YouTubeConfig
	Query   string
}

type YouTubeVideo struct {
	VideoID     string
	Title       string
	Description string
	PublishedAt time.Time
	Thumbnail   string
}

type APIResponse struct {
	NextPageToken string      `json:"nextPageToken"`
	Items         []VideoItem `json:"items"`
}

type VideoItem struct {
	ID      VideoID `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type VideoID struct {
	VideoID string `json:"videoId"`
}

type Snippet struct {
	PublishedAt string     `json:"publishedAt"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Thumbnails  Thumbnails `json:"thumbnails"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type APIKey struct {
	Key        string
	ErrorCount int64
	Enabled    bool
}
