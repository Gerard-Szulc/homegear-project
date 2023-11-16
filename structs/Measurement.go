package structs

import "gorm.io/gorm"

type MeasurementDto struct {
	Type  string  `json:"type"`
	Value float32 `json:"value"`
}

type Measurement struct {
	gorm.Model
	Type     string
	Value    float32
	DeviceID int
	Device   Device
}
