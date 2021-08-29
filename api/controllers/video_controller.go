package controllers

import (
	"net/http"
	"strconv"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/entities"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/middlewares"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/services"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/validators"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
)

var validate *validator2.Validate

func findAll(service services.VideoService) func (ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{ "videos": service.FindAll() })
	}
}

func createVideo(service services.VideoService) func (ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var video entities.Video
		e := ctx.ShouldBindJSON(&video)

		if e != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H { "error": e.Error() })
			return
		}

		e = validate.Struct(video)
		if e != nil {
			ctx.JSON(http.StatusBadRequest, gin.H { "error": e.Error() })
			return
		}

		service.Save(video)
		ctx.JSON(http.StatusOK, gin.H{ "video": video })
	}
}

func updateVideo(service services.VideoService) func (ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var video entities.Video
		e := ctx.ShouldBindJSON(&video)

		if e != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H { "error": e.Error() })
			return
		}

		id, e := strconv.ParseUint(ctx.Param("id"), 10, 0)

		if e != nil {
			ctx.JSON(http.StatusBadRequest, gin.H { "error": "no ID parameter provided" })
			return
		}

		video.ID = id
		e = validate.Struct(video)
		if e != nil {
			ctx.JSON(http.StatusBadRequest, gin.H { "error": e.Error() })
			return
		}

		service.Save(video)
		ctx.JSON(http.StatusOK, gin.H{ "video": video })
	}
}

func deleteVideo(service services.VideoService) func (ctx *gin.Context) {
	return func (ctx *gin.Context) {
		var video entities.Video
		id, e := strconv.ParseUint(ctx.Param("id"), 10, 0)

		if e != nil {
			ctx.JSON(http.StatusBadRequest, gin.H { "error": "no ID parameter provided" })
			return
		}

		video.ID = id
		service.Delete(video)

		ctx.JSON(http.StatusOK, gin.H { "message": "Video was deleted" })
	}
}

func RegisterVideoController(r *gin.RouterGroup, service services.VideoService) {
	g := r.Group("/videos")

	validate = validator2.New()
	_ = validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	g.PUT("/:id", updateVideo(service))
	g.DELETE("/:id", deleteVideo(service))
	g.GET("/", findAll(service))
	g.POST("/", middlewares.BasicAuth(), createVideo(service))
}
