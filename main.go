package main

import (
	"log"
	"net/http"

	"github.com/ErhanKurnaz/awesome-cool-bot-api/controllers"
	"github.com/ErhanKurnaz/awesome-cool-bot-api/services"
	"github.com/gin-gonic/gin"
)

var (
	videoService = services.NewVideoService()
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
	// Listen and Server in 0.0.0.0:8080
	log.Fatal(r.Run(":4200"))
}
