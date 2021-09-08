package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
    // TODO: change this with actual auth
	return gin.BasicAuth(gin.Accounts{
		"erhan": "kurnaz",
	})
}
