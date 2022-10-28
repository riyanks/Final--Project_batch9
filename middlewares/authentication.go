package middlewares

import (
	"final-project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(a *gin.Context) {
		verifyToken, err := helpers.VerifyToken(a)

		if err != nil {
			a.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		a.Set("userData", verifyToken)
		a.Next()
	}
}
