package router

import (
	"PayFam/internal/controllers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func backendRoutes(apiRouter *gin.RouterGroup, videoController *controllers.VideoController, logger *zap.Logger) {
	apiRouter.GET("/youtubeApi", func(ctx *gin.Context) {
		videoController.FetchVideos(ctx, logger)
	})
	apiRouter.GET("/videos", func(ctx *gin.Context) {
		videoController.GetVideos(ctx, logger)
	})
	apiRouter.GET("/videos/search", func(ctx *gin.Context) {
		videoController.SearchVideos(ctx, logger)
	})
}

func frontendRoutes(apiRouter *gin.RouterGroup) {
	apiRouter.GET("/videos/view", func(c *gin.Context) {
		c.HTML(http.StatusOK, "youtubeSearch.html", nil)
	})
}

func SetupRoutes(apiRouter *gin.RouterGroup, videoController *controllers.VideoController, logger *zap.Logger) {
	backendRoutes(apiRouter, videoController, logger)
	frontendRoutes(apiRouter)
}
