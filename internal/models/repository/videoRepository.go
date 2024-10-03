package repository

import (
	"PayFam/internal/models/dao"
	"PayFam/internal/models/entity"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{db: db}
}

func (repo *VideoRepository) GetVideoByVideoID(ctx context.Context, videoId string) (*entity.Video, error) {
	var videoRow dao.Video
	err := repo.db.WithContext(ctx).
		Where("video_id = ?", videoId).
		First(&videoRow).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	videoObj := videoRow.ToEntity()
	return &videoObj, nil
}

func (repo *VideoRepository) RecordVideo(ctx context.Context, videoObj *entity.Video) error {
	videoRow := &dao.Video{}
	videoRow.EntityToModel(*videoObj)

	dbErr := repo.db.WithContext(ctx).Create(videoRow).Error
	if dbErr != nil {
		return dbErr
	}

	videoObj.ID = videoRow.ID
	return nil
}

func (repo *VideoRepository) SaveVideoIfNotExists(ctx context.Context, video *entity.Video) error {
	existing, err := repo.GetVideoByVideoID(ctx, video.VideoId)
	if err != nil {
		return err
	}

	if existing != nil {
		log.Printf("Video with ID %s already exists", video.VideoId)
		return nil
	}

	return repo.RecordVideo(ctx, video)
}

func (repo *VideoRepository) GetVideos(ctx context.Context, limit, offset int) ([]entity.Video, error) {
	var videoRows []dao.Video
	err := repo.db.WithContext(ctx).Limit(limit).Offset(offset).Order("published_at DESC").Find(&videoRows).Error
	if err != nil {
		return nil, err
	}

	videos := make([]entity.Video, len(videoRows))
	for i, videoRow := range videoRows {
		videos[i] = videoRow.ToEntity()
	}

	return videos, nil
}

func (repo *VideoRepository) SearchVideos(ctx context.Context, query string, limit, offset int) ([]entity.Video, error) {
	var videoRows []dao.Video
	err := repo.db.WithContext(ctx).
		Where("title ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%").
		Limit(limit).
		Offset(offset).
		Order("published_at DESC").
		Find(&videoRows).Error

	if err != nil {
		return nil, err
	}

	videos := make([]entity.Video, len(videoRows))
	for i, videoRow := range videoRows {
		videos[i] = videoRow.ToEntity()
	}

	return videos, nil
}
