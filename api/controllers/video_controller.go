package controllers

import (
	"net/http"
	"strconv"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/entities"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/constants"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/middlewares"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/services"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/validators"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
)

// FindAll godoc
// @Summary Finds all videos in the system
// @ID FindAll videos
// @Tags videos,list
// @Produce json
// @Success 200 {array} entities.Video
// @Failure 400 {object} interface{}
// @Router /videos [get]
func findAll(service services.VideoService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Set(constants.ResponseField, service.FindAll())
	}
}

// CreateVideo godoc
// @Summary Creates a new video in the system
// @ID Create videos
// @Tags videos,create
// @Accept json
// @Produce json
// @Param video body entities.Video true "Create video"
// @Success 201 {object} entities.Video
// @Failure 422 {object} interface{}
// @Failure 400 {object} interface{}
// @Router /videos [post]
func createVideo(service services.VideoService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		body, _ := ctx.Get(constants.BodyField)
		video := body.(*entities.Video)

		service.Save(video)
		ctx.Set(constants.ResponseField, video)
		ctx.Set(constants.StatusField, http.StatusCreated)
	}
}

// UpdateVideo godoc
// @Summary Update a video in the system
// @ID Update videos
// @Tags videos,update
// @Accept json
// @Produce json
// @Param video body entities.Video true "Update video"
// @Param  id path int true "Video ID"
// @Success 200 {object} entities.Video
// @Failure 422 {object} interface{}
// @Failure 400 {object} interface{}
// @Router /videos/{id} [put]
func updateVideo(service services.VideoService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		body, _ := ctx.Get(constants.BodyField)
		video := body.(*entities.Video)

		service.Update(video)
		ctx.Set(constants.ResponseField, video)
	}
}

// DeleteVideo godoc
// @Summary Delete a video in the system
// @ID Delete videos
// @Tags videos,delete
// @Produce json
// @Param  id path int true "Video ID"
// @Success 200 {object} entities.Video
// @Failure 422 {object} interface{}
// @Failure 400 {object} interface{}
// @Router /videos/{id} [delete]
func deleteVideo(service services.VideoService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var video entities.Video
        id, e := strconv.ParseUint(ctx.Param("id"), 10, 64)

        if e != nil {
            ctx.Set(constants.ErrorField, e)
        }

		video.ID = id
		service.Delete(&video)

		ctx.Set(constants.ResponseField, "Video was deleted")
	}
}

func RegisterVideoController(r *gin.RouterGroup, service services.VideoService) {
	g := r.Group("/videos")

	validate := validator2.New()
	_ = validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	g.PUT("/:id",
		middlewares.BodyParserMiddleware(&entities.Video{}),
		middlewares.SetBodyId("id"),
		middlewares.ValidateBodyMiddleware(validate),
		updateVideo(service))

	g.DELETE("/:id", deleteVideo(service))
	g.GET("/", findAll(service))

	g.POST("/",
		middlewares.BodyParserMiddleware(&entities.Video{}),
		middlewares.ValidateBodyMiddleware(validate),
		createVideo(service))
}
