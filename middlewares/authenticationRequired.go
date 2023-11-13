package middlewares

import (
	"dustData/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.ValidateRequestToken(c.Request) {
			c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{
				"message": "error:token_not_valid",
			}})
			c.Abort()
			return
		}
		c.Next()
	}
}
