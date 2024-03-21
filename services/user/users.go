package users

import (
	"github.com/gin-gonic/gin"
	"homegear/db"
	"homegear/structs"
	"net/http"
)

func GetUser(id string, c *gin.Context) {
	// Find and return user
	user := structs.ResponseUser{}
	db.DB.Model(&structs.User{}).Select("id, created_at, email, username").Where("id = ? ", id).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]interface{}{"message": "error.user_not_found"}})
		return
	}
	c.JSON(http.StatusOK, user)
}
func GetUsers(c *gin.Context) {
	// Find and return user
	users := []structs.ResponseUser{}
	db.DB.Model(&structs.User{}).Select("id, username, email").Find(&users)

	c.JSON(http.StatusOK, users)
}
