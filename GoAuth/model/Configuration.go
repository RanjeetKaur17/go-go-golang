package model

//Struct for Configurations
type Configuration struct {
	Database	map[string]string
	Redis		string
	ServerPort	string
}