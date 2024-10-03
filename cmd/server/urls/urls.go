package urls

import (
	"PayFam/cmd/router"
	"PayFam/internal/controllers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AddRoutes(r *gin.Engine, videoController *controllers.VideoController, logger *zap.Logger) {
	addHealthCheck(r)
	externalRouterGroup := r.Group("/internal/v1")
	router.SetupRoutes(externalRouterGroup, videoController, logger)
}

func addHealthCheck(r *gin.Engine) {
	r.GET("/knockknock", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Service is up and running",
		})
	})
}
