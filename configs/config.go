package configs

import (
	external "PayFam/external/youtube"
	"PayFam/internal/models/repository"
	"gorm.io/gorm"
	"time"
)

func Load(db *gorm.DB) (*YtConfig, error) {
	ytApiConfig := external.LoadConfig()
	videoRepo := repository.NewVideoRepository(db)
	publishedTime := time.Now().Add(-24 * time.Hour)
	fetchIntervalTime := FetchIntervalTime

	return &YtConfig{
		YTApiConfig:       *ytApiConfig,
		VideoRepo:         videoRepo,
		PublishedTime:     publishedTime,
		FetchIntervalTime: fetchIntervalTime,
	}, nil
}
