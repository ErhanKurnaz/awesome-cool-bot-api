package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/api_errors"
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

		casted_err, err_ok := err.(api_errors.ApiError)

		if !err_ok {
			utils.AbortWithErr(ctx, http.StatusInternalServerError,
				errors.New(fmt.Sprintf("an error was specified %T didn't comply to the ApiError interface", err)))

			return
		}

		utils.RespondJSON(ctx, casted_err.Code(), models.NewErrorResponse(casted_err))
	}
}
