package structs

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Active   bool      `gorm:"default:false"`
	Devices  []*Device `gorm:"many2many:user_devices;"`
}

type ResponseUser struct {
	ID       uint
	Username string
	Email    string
}

type ResponseUserWithToken struct {
	Data *ResponseUser
	Jwt  string `json:"jwt"`
}

type Validation struct {
	Value string
	Valid string
}

type Register struct {
	Username string
	Email    string
	Password string
}

type ErrResponse struct {
	Message string
}
