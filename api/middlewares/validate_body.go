package middlewares

import (
	"net/http"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/constants"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/errors"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/models"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/utils"
	"github.com/gin-gonic/gin"
	validator2 "github.com/go-playground/validator/v10"
)

func ValidateBodyMiddleware(validate *validator2.Validate) gin.HandlerFunc {
	return func (ctx *gin.Context) {
        if utils.IsReqHandled(ctx) {
            return
        }

		body, exists := ctx.Get(constants.BodyField)
        if !exists {
            utils.AbortWithErr(ctx, http.StatusInternalServerError, errors.NewNoBodyError())
            return
        }

        e := validate.Struct(body)
        if e != nil {
            utils.AbortWithJSON(ctx, http.StatusBadRequest, models.NewErrorResponse(e))
            return
        }
	}
}
