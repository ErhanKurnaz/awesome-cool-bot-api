package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/constants"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/models"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/utils"
	"github.com/gin-gonic/gin"
)

func ErrorResponseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

        if utils.IsReqHandled(ctx) {
            return
        }

		err, err_exists := ctx.Get(constants.ErrorField)

		if !err_exists {
			return
		}

		casted_err, err_ok := err.(error)

		if !err_ok {
			utils.AbortWithErr(ctx, http.StatusInternalServerError,
				errors.New(fmt.Sprintf("an error was specified %T didn't comply to the error interface", err)))

			return
		}

		status, status_exists := ctx.Get(constants.StatusField)
		status_int, cast_ok := status.(int)

		if !status_exists || !cast_ok {
			status = http.StatusBadRequest
		}

		utils.RespondJSON(ctx, status_int, models.NewErrorResponse(casted_err))
	}
}
