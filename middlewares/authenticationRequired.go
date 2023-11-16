package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homegear/utils"
	"net/http"
)

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !utils.ValidateRequestToken(c) {
			c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{
				"message": "error:token_not_valid",
			}})
			c.Abort()
			return
		}
		fmt.Println(c.Get("userId"))
		c.Next()
	}
}
