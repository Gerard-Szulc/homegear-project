package structs

import "gorm.io/gorm"

type MeasurementDto struct {
	Type  string  `json:"type"`
	Value float32 `json:"value"`
}

type MeasurementsDto []MeasurementDto

type Measurement struct {
	gorm.Model
	Type     string
	Value    float32
	DeviceID int
	Device   Device
}
type Measurements []Measurement
type ResponseMeasurement struct {
	ID        uint    `json:"id"`
	Type      string  `json:"type"`
	Value     float32 `json:"value"`
	CreatedAt string  `json:"createdAt"`
	DeviceID  int     `json:"deviceId"`
}
