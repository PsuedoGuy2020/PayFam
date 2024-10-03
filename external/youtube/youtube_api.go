package external

import (
	"PayFam/internal/models/entity"
	"PayFam/internal/models/repository"
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func LoadConfig() *Config {
	return &Config{
		YouTube: YouTubeConfig{
			APIKeys: []APIKey{
				{Key: API_KEY1, Enabled: true},
				{Key: API_KEY2, Enabled: true},
				{Key: API_KEY3, Enabled: true},
			},
		},
		Query: Query,
	}
}

func YoutubeConfig(repo *repository.VideoRepository, apiKeys []APIKey, query string, fetchInterval time.Duration, publishedAfter time.Time) *VideoService {
	return &VideoService{
		Repo:           repo,
		ApiKeys:        apiKeys,
		Query:          query,
		FetchInterval:  fetchInterval,
		PublishedAfter: publishedAfter,
	}
}

func (vs *VideoService) FetchYouTubeVideosRoundRobin(ctx context.Context, logger *zap.Logger) {
	keyCount := len(vs.ApiKeys)
	if keyCount == 0 {
		return
	}

	ticker := time.NewTicker(vs.FetchInterval)
	defer ticker.Stop()

	keyIndex := 0
	var totalPageCount int64 = 0

	for {
		select {
		case <-ticker.C:
			apiKey := &vs.ApiKeys[keyIndex]
			if !apiKey.Enabled {
				keyIndex = (keyIndex + 1) % keyCount
				continue
			}
			publishedAfterStr := vs.PublishedAfter.UTC().Format(time.RFC3339)
			videos, pagesFetched, err := vs.FetchVideos(apiKey.Key, vs.Query, publishedAfterStr, MaxResults, logger)
			if err != nil {
				logger.Error("Error fetching videos", zap.String("api_key", apiKey.Key), zap.Error(err))
				apiKey.ErrorCount++
				if apiKey.ErrorCount >= MaxErrors {
					apiKey.Enabled = false
					logger.Error("Disabling due to consecutive errors\n", zap.String("api_key", apiKey.Key))
				}
			} else {
				apiKey.ErrorCount = 0
				for _, video := range videos {
					ytVideoToEntityVideo := entity.Video{
						VideoId:      video.VideoID,
						Title:        video.Title,
						Description:  video.Description,
						PublishedAt:  video.PublishedAt,
						ThumbnailUrl: video.Thumbnail,
					}
					if dbErr := vs.Repo.SaveVideoIfNotExists(ctx, &ytVideoToEntityVideo); dbErr != nil {
						logger.Error("Failed to save video: %v", zap.Error(dbErr))
					}
				}
			}
			keyIndex = (keyIndex + 1) % keyCount
			totalPageCount += pagesFetched
			if totalPageCount >= MaxPages {
				logger.Info("Max pages fetched. Stopping background task")
				return
			}
		}
	}
}

func (vs *VideoService) FetchVideos(apiKey string, query string, publishedAfter string, maxResults int64, logger *zap.Logger) ([]YouTubeVideo, int64, error) {
	var videos []YouTubeVideo
	var nextPageToken = ""
	var pageCount int64 = 0

	for {
		url := fmt.Sprintf("%v?key=%v&q=%v&type=video&part=snippet&order=date&publishedAfter=%v&maxResults=%v&pageToken=%v",
			youtubeApiUrl, apiKey, query, publishedAfter, maxResults, nextPageToken)

		resp, apiErr := http.Get(url)
		if apiErr != nil {
			return nil, pageCount, apiErr
		}

		respBody, respErr := ioutil.ReadAll(resp.Body)
		err := resp.Body.Close()
		if err != nil {
			return nil, pageCount, err
		}
		if respErr != nil {
			return nil, pageCount, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, pageCount, fmt.Errorf("HTTP error: %s, Response: %s", resp.Status, string(respBody))
		}

		var result APIResponse
		if resErr := json.Unmarshal(respBody, &result); resErr != nil {
			return nil, pageCount, resErr
		}

		for _, item := range result.Items {
			publishedAt, dateErr := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			if dateErr != nil {
				logger.Error("Error parsing time: %v", zap.Error(dateErr))
				continue
			}

			video := YouTubeVideo{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				PublishedAt: publishedAt,
				Thumbnail:   item.Snippet.Thumbnails.Default.URL,
				VideoID:     item.ID.VideoID,
			}
			videos = append(videos, video)
		}

		pageCount++

		if result.NextPageToken == "" || pageCount >= MaxPages {
			break
		}

		nextPageToken = result.NextPageToken
	}

	return videos, pageCount, nil
}
