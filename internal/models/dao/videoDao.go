package dao

import (
	"PayFam/internal/models/entity"
	"time"
)

type Video struct {
	ID           uint      `gorm:"column:id;primaryKey;autoIncrement"`
	VideoId      string    `json:"video_id"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Title        string    `gorm:"column:title;index"`
	Description  string    `gorm:"column:description"`
	PublishedAt  time.Time `gorm:"column:published_at"`
	ThumbnailUrl string    `gorm:"column:thumbnail_url"`
}

func (v *Video) TableName() string {
	return "videos"
}

func (v *Video) ToEntity() entity.Video {
	return entity.Video{
		ID:           v.ID,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		Title:        v.Title,
		Description:  v.Description,
		PublishedAt:  v.PublishedAt,
		ThumbnailUrl: v.ThumbnailUrl,
	}
}

func (v *Video) EntityToModel(videoEntity entity.Video) {
	v.ID = videoEntity.ID
	v.VideoId = videoEntity.VideoId
	v.CreatedAt = videoEntity.CreatedAt
	v.UpdatedAt = videoEntity.UpdatedAt
	v.Title = videoEntity.Title
	v.Description = videoEntity.Description
	v.PublishedAt = videoEntity.PublishedAt
	v.ThumbnailUrl = videoEntity.ThumbnailUrl
}
