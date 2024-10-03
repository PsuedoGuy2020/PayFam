package configs

import (
	external "PayFam/external/youtube"
	"PayFam/internal/models/repository"
	"time"
)

type AppConfig struct {
	Server struct {
		Port        string `yaml:"PORT"`
		LogFilePath string `yaml:"LOG_FILE_PATH"`
	} `yaml:"SERVER"`
	Database struct {
		Host     string `yaml:"HOST"`
		Port     string `yaml:"PORT"`
		User     string `yaml:"USER"`
		Password string `yaml:"PASSWORD"`
		Name     string `yaml:"NAME"`
	} `yaml:"DATABASE"`
}

type YtConfig struct {
	YTApiConfig       external.Config
	VideoRepo         *repository.VideoRepository
	PublishedTime     time.Time
	FetchIntervalTime time.Duration
}
