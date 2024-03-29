package main

import (
	"log"
	"net/http"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/controllers"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/database"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/docs"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/middlewares"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/repositories"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/services"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

var (
	db = database.SetupDBConnection()
	videoRepository = repositories.NewVideoRepository(db)
	videoService = services.NewVideoService(videoRepository)
)


func setupRouter() *gin.Engine {
	engine := gin.Default()

	url := ginSwagger.URL("http://localhost:4200/swagger/doc.json") // The url pointing to API definition
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

    engine.Use(middlewares.ResponseMiddleware(), middlewares.ErrorResponseMiddleware())
	r := engine.Group("api/v1")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{ "data": "pong!" })
	})

	controllers.RegisterVideoController(r, videoService)

	return engine
}

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "AwesomeCoolBots API"
	docs.SwaggerInfo.Description = "The api for a little Twitch bot called 'AwesomeCoolBots'"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:4200"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string {"http"}

	r := setupRouter()
	defer database.Close(db)
	// Listen and Server in 0.0.0.0:4200
	log.Println(r.Run(":4200"))
}
