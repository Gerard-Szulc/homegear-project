package middlewares

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"strings"
	"time"
)

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !ValidateRequestToken(c) {
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

func ValidateRequestToken(context *gin.Context) bool {
	r := context.Request
	jwtKey, exists := os.LookupEnv("JWTKEY")
	if !exists {
		fmt.Println(exists)
		return false
	}
	jwtToken := r.Header.Get("Authorization")
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	if !strings.Contains(cleanJWT, ".") {
		return false
	}
	cleanJWTHeader := strings.Split(cleanJWT, ".")[0]
	cleanJWTPayload := strings.Split(cleanJWT, ".")[1]
	cleanJWTSecret := strings.Split(cleanJWT, ".")[2]
	_, err := jwt.DecodeSegment(cleanJWTHeader)
	if err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			fmt.Println("base64 input is corrupt, check service Key")
			return false
		}
		return false
	}
	_, err = jwt.DecodeSegment(cleanJWTPayload)
	if err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			panic("\nbase64 input is corrupt, check service Key")
		}
		fmt.Println(err)
	}
	_, err = jwt.DecodeSegment(cleanJWTSecret)
	if err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			panic("\nbase64 input is corrupt, check service Key")
		}
		fmt.Println(err)
	}

	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return false
	}
	//HandleErrRequest(err)

	now := time.Now()
	expiry := tokenData["expiry"].(float64)
	fmt.Println(token.Valid)
	fmt.Println(tokenData["userId"])
	context.Set("userId", tokenData["userId"])
	if (tokenData["deviceId"]) != nil {
		fmt.Println(tokenData["deviceId"])
		context.Set("deviceId", tokenData["deviceId"])
	}

	expired := now.After(time.Unix(int64(expiry), 0))
	if expired {
		return false
	}
	if !token.Valid {
		return false
	}

	return true
}
