package middlewares

import (
	"net/http"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/constants"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/models"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/utils"
	"github.com/gin-gonic/gin"
)

func ResponseMiddleware() gin.HandlerFunc {
    return func (ctx *gin.Context) {
        ctx.Next()
        if utils.IsReqHandled(ctx) {
            return
        }

        res, res_exists := ctx.Get(constants.ResponseField)

        if !res_exists {
            res = nil
        }

        status, status_exists := ctx.Get(constants.StatusField)
        status_int, cast_ok := status.(int)

        if !status_exists || !cast_ok {
            status_int = http.StatusOK
        }

        utils.RespondJSON(ctx, status_int, models.NewSuccessResponse(res))
    }
}

