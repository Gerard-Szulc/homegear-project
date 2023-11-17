package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"homegear/db"
	"homegear/structs"
	"net/http"
)

func GetUser(id string, c *gin.Context) {
	// Find and return user
	user := &structs.User{}

	result := db.DB.Where("id = ? ", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]interface{}{"message": "error.user_not_found"}})
		return
	}
	c.JSON(http.StatusOK, result)
}
func GetUsers(c *gin.Context) {
	// Find and return user
	var users []structs.User
	db.DB.Select("id, username, email").Find(&users)

	responseUsers := []structs.ResponseUser{}
	for _, user := range users {
		responseUser := structs.ResponseUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
		responseUsers = append(responseUsers, responseUser)
	}

	c.JSON(http.StatusOK, responseUsers)
}
