package main

import (
	"PayFam/cmd/Db"
	"PayFam/cmd/common"
	"PayFam/cmd/server/urls"
	"PayFam/configs"
	external "PayFam/external/youtube"
	"PayFam/internal/controllers"
	"PayFam/internal/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("../../views/*.html")

	logger, err := common.NewLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logger.Sync()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configData, err := ioutil.ReadFile("../../configs/dev/config.yaml")

	if err != nil {
		panic(err)
	}

	var AppConfig configs.AppConfig
	err = yaml.Unmarshal(configData, &AppConfig)
	if err != nil {
		panic(err)
	}

	db, err := Db.InitDatabase(
		AppConfig.Database.Host,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Name,
		AppConfig.Database.Port,
	)
	if err != nil {
		panic(err)
	}
	appConfig, err := configs.Load(db)
	if err != nil {
		logger.Error("Could not load application configuration: %v", zap.Error(err))
	}

	ytConfig := external.YoutubeConfig(appConfig.VideoRepo, appConfig.YTApiConfig.YouTube.APIKeys, appConfig.YTApiConfig.Query, appConfig.FetchIntervalTime, appConfig.PublishedTime)
	videoService := service.NewVideoService(appConfig.VideoRepo, appConfig.YTApiConfig.YouTube.APIKeys, appConfig.YTApiConfig.Query)
	videoController := controllers.NewVideoController(videoService, logger)

	urls.AddRoutes(r, videoController, logger)

	go ytConfig.FetchYouTubeVideosRoundRobin(ctx, logger)

	serverPort := AppConfig.Server.Port

	if servErr := r.Run(fmt.Sprintf(":%s", serverPort)); servErr != nil {
		logger.Error("Server failed to start: %v", zap.Error(servErr))
	}
}
