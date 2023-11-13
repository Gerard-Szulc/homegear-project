package structs

import (
	"gorm.io/gorm"
)

type DeviceDto struct {
	Id    int        `json:"id"`
	Name  string     `json:"name"`
	Label string     `json:"label"`
	Type  DeviceType `json:"type"`
	IP    string     `json:"ip"`
}

type DevicesDto []DeviceDto

type DeviceType string

const (
	BulbTopic        DeviceType = "Bulb"
	TemperatureTopic DeviceType = "Temperature"
	HumidityTopic    DeviceType = "Humidity"
	PollutionTopic   DeviceType = "Pollution"
)

type Device struct {
	gorm.Model
	Name  string
	Label string
	IP    string
	Users []*User `gorm:"many2many:user_devices;"`
	Type  DeviceType
}
