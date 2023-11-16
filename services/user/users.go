package users

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"homegear/db"
	"homegear/structs"
	"homegear/utils"
	"net/http"
	"os"
	"time"
)

func prepareToken(user *structs.User) string {
	jwtKey, exists := os.LookupEnv("JWTKEY")
	if !exists {
		fmt.Println(exists)
	}
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(jwtKey))
	utils.HandleErr(err)
	return token
}

func prepareResponse(user *structs.User, withToken bool, c *gin.Context) {
	responseUser := &structs.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	var response structs.ResponseUserWithToken
	// Add withToken feature to prepare response
	if withToken {
		var token = prepareToken(user)
		response.Jwt = token
	}
	response.Data = responseUser
	c.JSON(http.StatusOK, response)
}

func Login(username string, pass string, c *gin.Context) {
	fmt.Printf("ClientIP: %s\n", c.ClientIP())

	// Add validation to login
	valid := utils.Validation(
		[]structs.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		// Connect DB
		user := &structs.User{}
		result := db.DB.Where("username = ? ", username).First(&user)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.user_not_found"}})
			return
		}
		// Verify password
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.wrong_password"}})
			return
		}
		if !user.Active {
			c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.account_not_active"}})
			return
		}
		prepareResponse(user, true, c)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]interface{}{"message": "not valid values"}})
	}
}

// Create registration function
func Register(username string, email string, pass string, c *gin.Context) {
	// Add validation to registration
	valid := utils.Validation(
		[]structs.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		// Create registration logic
		// Connect DB
		generatedPassword := utils.HashAndSalt([]byte(pass))
		user := &structs.User{Username: username, Email: email, Password: generatedPassword}
		db.DB.Create(&user)
		prepareResponse(user, true, c)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]interface{}{"message": "error.values_not_valid"}})
	}
}

func GetUser(id string, c *gin.Context) {
	// Find and return user
	user := &structs.User{}

	result := db.DB.Where("id = ? ", id).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]interface{}{"message": "error.user_not_found"}})
		return
	}
	prepareResponse(user, false, c)
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
