package controllers

import (
	external "PayFam/external/youtube"
	"PayFam/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type VideoController struct {
	Service *service.VideoService
}

func NewVideoController(service *service.VideoService, logger *zap.Logger) *VideoController {
	return &VideoController{Service: service}
}

func (vc *VideoController) GetVideos(ctx *gin.Context, logger *zap.Logger) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", external.Page))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", external.Limit))

	videos, err := vc.Service.GetVideos(ctx.Request.Context(), page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, videos)
}

func (vc *VideoController) SearchVideos(ctx *gin.Context, logger *zap.Logger) {
	query := ctx.Query("query")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", external.Page))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", external.Limit))

	videos, err := vc.Service.SearchVideos(ctx.Request.Context(), query, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, videos)
}

func (vc *VideoController) FetchVideos(ctx *gin.Context, logger *zap.Logger) {
	query := ctx.DefaultQuery("query", external.Query)
	publishedAfter := ctx.DefaultQuery("publishedAfter", time.Now().Add(-24*time.Hour).UTC().Format(time.RFC3339))
	maxResults := external.MaxResults

	videos, err := vc.Service.FetchVideos(query, publishedAfter, maxResults, logger)
	if err != nil {
		logger.Error("Error fetching YouTube videos: %v", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos"})
		return
	}

	ctx.JSON(http.StatusOK, videos)
}
