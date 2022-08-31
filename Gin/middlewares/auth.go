package middlewares

import (
	"esm-backend/utilities"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) < 2 {
			context.JSON(http.StatusForbidden, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		token := bearerToken[1]
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		_, err := utilities.VerifyJwtToken(token)
		if err != nil {
			context.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
