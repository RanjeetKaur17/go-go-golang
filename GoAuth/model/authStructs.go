package model

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSession struct {
	Username 		string `json:"username"`
	DeviceId 		string `json:"deviceId"`
	UserSecureKey 	string `json:"userSecureKey"`
}

type Response struct {
	Data 		interface{} `json:"data"`
	Exception	string		`json:"exception"`
}