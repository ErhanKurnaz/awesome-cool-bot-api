package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ErhanKurnaz/awesome-cool-bot/api/constants"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/entities"
	"github.com/ErhanKurnaz/awesome-cool-bot/api/utils"
	"github.com/gin-gonic/gin"
)

func SetBodyId(field string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
        if utils.IsReqHandled(ctx) {
            return
        }

		body, exists := ctx.Get(constants.BodyField)

		if !exists {
			utils.AbortWithErr(ctx, http.StatusInternalServerError, errors.New("Body ID not set since the body doesn't exists"))
			return
		}

		entity, ok := body.(entities.IBaseModel)

		if !ok {
			msg := fmt.Sprintf("Body ID not set since body with type %T doesn't implement entities.IBaseModel\n", entity)
			utils.AbortWithErr(ctx, http.StatusInternalServerError, errors.New(msg))
			return
		}

		id, e := strconv.ParseUint(ctx.Param(field), 10, 0)

		if e != nil {
			msg := fmt.Sprintf("Body ID not set since there was no ID parameter found with key: %s\n", field)
			utils.AbortWithErr(ctx, http.StatusInternalServerError, errors.New(msg))
			return
		}

		entity.SetId(id)
	}
}
