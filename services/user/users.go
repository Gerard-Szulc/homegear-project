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
	user := structs.ResponseUser{}
	result := db.DB.Model(&structs.User{}).Where("id = ? ", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]interface{}{"message": "error.user_not_found"}})
		return
	}
	c.JSON(http.StatusOK, result)
}
func GetUsers(c *gin.Context) {
	// Find and return user
	users := []structs.ResponseUser{}
	db.DB.Model(&structs.User{}).Select("id, username, email").Find(&users)

	c.JSON(http.StatusOK, users)
}
