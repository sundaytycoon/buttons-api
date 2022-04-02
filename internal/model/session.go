package model

type Session struct {
	Token    string    `json:"session_token"`
	UserData *UserData `json:"user_data"`
}

type UserData struct {
	Id       string `json:"id"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	SignUp   bool   `json:"sign_up"`
	Username string `json:"username"`

	// user_metas
	Profile string `json:"profile"`

	// user_devices
	DeviceId       string `json:"user_device_id"`
	DeviceType     string `json:"device_type"`
	DeviceBrowser  string `json:"device_os"`
	DevicePlatform string `json:"device_platform"`
}
