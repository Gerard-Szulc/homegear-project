package structs

type LoginDto struct {
	Username string
	Password string
}
type LoginDeviceDto struct {
	DeviceId uint   `json:"deviceId"`
	Token    string `json:"token"`
}

type LoginDeviceResponseViewModel struct {
	DeviceId uint   `json:"deviceId"`
	Jwt      string `json:"jwt"`
}
