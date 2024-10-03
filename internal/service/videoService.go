package service

import (
	external "PayFam/external/youtube"
	"PayFam/internal/models/entity"
	"PayFam/internal/models/repository"
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type VideoService struct {
	Repo           *repository.VideoRepository
	ApiKeys        []external.APIKey
	Query          string
	fetchInterval  time.Duration
	publishedAfter time.Time
}

func NewVideoService(repo *repository.VideoRepository, apiKeys []external.APIKey, query string) *VideoService {
	return &VideoService{
		Repo:    repo,
		ApiKeys: apiKeys,
		Query:   query,
	}
}

func (vs *VideoService) GetVideos(ctx context.Context, page, limit int) ([]entity.Video, error) {
	offset := (page - 1) * limit
	return vs.Repo.GetVideos(ctx, limit, offset)
}

func (vs *VideoService) SearchVideos(ctx context.Context, query string, page, limit int) ([]entity.Video, error) {
	offset := (page - 1) * limit
	return vs.Repo.SearchVideos(ctx, query, limit, offset)
}

func (vs *VideoService) FetchVideos(query string, publishedAfter string, maxResults int64, logger *zap.Logger) ([]external.YouTubeVideo, error) {
	if len(vs.ApiKeys) == 0 {
		return nil, fmt.Errorf("no API keys available")
	}

	apiKey := vs.ApiKeys[0].Key

	ytVs := external.VideoService{
		Repo:           vs.Repo,
		ApiKeys:        vs.ApiKeys,
		Query:          vs.Query,
		FetchInterval:  vs.fetchInterval,
		PublishedAfter: vs.publishedAfter,
	}

	videos, _, err := ytVs.FetchVideos(apiKey, query, publishedAfter, maxResults, logger)
	if err != nil {
		logger.Error("Error fetching YouTube videos: %v", zap.Error(err))
		return nil, err
	}

	return videos, nil
}
