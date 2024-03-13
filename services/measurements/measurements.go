package measurementsService

import (
	"homegear/db"
	"homegear/structs"
)

func GetMeasurements(deviceId string, startDate string, endDate string) (measurements []structs.ResponseMeasurement, err error) {
	measurements = []structs.ResponseMeasurement{}

	if startDate != "" && endDate != "" {
		db.DB.Model(&structs.Measurement{}).Select("id, created_at, type, value, device_id").Where("device_id = ? AND created_at BETWEEN ? AND ?", deviceId, startDate, endDate).Find(&measurements)
	} else {
		db.DB.Model(&structs.Measurement{}).Select("id, created_at, type, value, device_id").Where("device_id = ?", deviceId).Find(&measurements)
	}

	return measurements, nil
}

func UserOwnsDevice(userId float64, deviceId string) bool {
	device := &structs.Device{}
	db.DB.Raw("SELECT * FROM devices JOIN user_devices on devices.id = user_devices.device_id WHERE (id = ?) AND (user_devices.user_id = ?)", deviceId, userId).First(&device)
	return device.ID != 0
}
