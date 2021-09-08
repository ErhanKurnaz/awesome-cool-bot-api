package main

import (
	"github.com/ErhanKurnaz/awesome-cool-bot/api/controllers"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/repositories"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterControllers(r *gin.RouterGroup, db *gorm.DB) {
	videoRepository := repositories.NewVideoRepository(db)
	videoService := services.NewVideoService(videoRepository)
	controllers.RegisterVideoController(r, videoService)
}
