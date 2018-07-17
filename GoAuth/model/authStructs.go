package model

//Struct for User Credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Struct for User Session
type UserSession struct {
	Username 		string `json:"username"`
	DeviceId 		string `json:"deviceId"`
	UserSecureKey 	string `json:"userSecureKey"`
}

//Struct for Response Object
type Response struct {
	Data 		interface{} `json:"data"`
	Exception	string		`json:"exception"`
}