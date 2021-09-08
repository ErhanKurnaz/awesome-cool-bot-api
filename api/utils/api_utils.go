package utils

import (
	"github.com/ErhanKurnaz/awesome-cool-bot/api/constants"
	"github.com/gin-gonic/gin"
)

func IsReqHandled(ctx *gin.Context) bool {
    is_handled, exists := ctx.Get(constants.RequestHandled)

    if !exists {
        return false
    }

    handled, ok := is_handled.(bool)

    if !ok {
        return false
    }

    return handled
}

func AbortWithJSON(ctx *gin.Context, status int, body interface{}) {
    ctx.AbortWithStatusJSON(status, body)
    ctx.Set(constants.RequestHandled, true)
}

func AbortWithErr(ctx *gin.Context, status int, err error) {
    ctx.AbortWithError(status, err)
    ctx.Set(constants.RequestHandled, true)
}

func RespondJSON(ctx *gin.Context, status int, body interface{}) {
    ctx.JSON(status, body)
    ctx.Set(constants.RequestHandled, true)
}

