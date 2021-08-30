package main

import (
	"log"
	"net/http"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/controllers"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/database"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/repositories"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/services"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	db = database.SetupDBConnection()
	videoRepository = repositories.NewVideoRepository(db)
	videoService = services.NewVideoService(videoRepository)
)


func setupRouter() *gin.Engine {
	engine := gin.Default()
	r := engine.Group("api/v1")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "data": "pong!" })
	})

	controllers.RegisterVideoController(r, videoService)

	return engine
}

func main() {
	r := setupRouter()
	defer database.Close(db)
	// Listen and Server in 0.0.0.0:4200
	log.Println(r.Run(":4200"))
}
