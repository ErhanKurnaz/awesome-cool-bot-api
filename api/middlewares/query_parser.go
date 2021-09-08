package middlewares

import (
	"net/http"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/constants"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/models"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/utils"
	"github.com/gin-gonic/gin"
)

func QueryParserMiddleware(entity interface{}) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        if utils.IsReqHandled(ctx) {
            return
        }

        e := ctx.ShouldBindQuery(&entity)

        if e != nil {
            utils.AbortWithJSON(ctx, http.StatusUnprocessableEntity, models.NewErrorResponse(e))
            return
        }

        ctx.Set(constants.QueryField, entity)
    }
}

