package authentication

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
	"regexp"
	"time"
)

func PrepareToken(user *structs.User) string {
	jwtKey, exists := os.LookupEnv("JWTKEY")
	if !exists {
		panic("JWTKEY not provided")
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
func PrepareDeviceToken(user *structs.User, deviceId uint) string {
	jwtKey, exists := os.LookupEnv("JWTKEY")
	if !exists {
		panic("JWTKEY not provided")
	}
	tokenContent := jwt.MapClaims{
		"user_id":   user.ID,
		"expiry":    time.Now().Add(time.Hour * 24 * 180).Unix(),
		"device_id": deviceId,
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte(jwtKey))
	utils.HandleErr(err)
	return token
}

func Login(username string, pass string, c *gin.Context) {
	fmt.Printf("ClientIP: %s\n", c.ClientIP())

	valid := Validation(
		[]structs.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		user := &structs.User{}
		result := db.DB.Where("username = ? ", username).First(&user)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.user_not_found"}})
			return
		}
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.wrong_password"}})
			return
		}
		if !user.Active {
			c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.account_not_active"}})
			return
		}
		prepareAuthResponse(user, c)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]interface{}{"message": "not valid values"}})
	}
}
func validateDeviceTempToken(token string) (valid bool) {
	// todo add temp token validation
	fmt.Println(token)
	return true
}
func LoginDevice(deviceId uint, token string, c *gin.Context) {
	fmt.Printf("ClientIP: %s\n", c.ClientIP())
	userId, exists := c.Get("userId")
	if exists != true {
		c.JSON(http.StatusUnauthorized, gin.H{"error": map[string]interface{}{"message": "error.no_user"}})
		return
	}
	valid := validateDeviceTempToken(token)
	if valid != true {
		c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.invalid_token"}})
		return
	}

	user := &structs.User{}
	result := db.DB.Where("id = ? ", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.user_not_found"}})
		return
	}
	if !user.Active {
		c.JSON(http.StatusForbidden, gin.H{"error": map[string]interface{}{"message": "error.account_not_active"}})
		return
	}
	prepareAuthDeviceResponse(user, deviceId, c)
}

// Create registration function
func Register(username string, email string, pass string, c *gin.Context) {
	// Add validation to registration
	valid := Validation(
		[]structs.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		// Create registration logic
		// Connect DB
		generatedPassword := HashAndSalt([]byte(pass))
		user := &structs.User{Username: username, Email: email, Password: generatedPassword}
		db.DB.Create(&user)
		prepareAuthResponse(user, c)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": map[string]interface{}{"message": "error.values_not_valid"}})
	}
}

func prepareAuthResponse(user *structs.User, c *gin.Context) {
	responseUser := &structs.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	var response structs.ResponseUserWithToken
	response.Jwt = PrepareToken(user)
	response.Data = responseUser
	c.JSON(http.StatusOK, response)
}

func prepareAuthDeviceResponse(user *structs.User, deviceId uint, c *gin.Context) {
	var response structs.LoginDeviceResponseViewModel
	response.Jwt = PrepareDeviceToken(user, deviceId)
	response.DeviceId = deviceId
	c.JSON(http.StatusOK, response)
}

func Validation(values []structs.Validation) bool {
	username := regexp.MustCompile("^([A-Za-z0-9]{5,})+$")
	email := regexp.MustCompile("^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$")
	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	utils.HandleErr(err)
	return string(hashed)
}
