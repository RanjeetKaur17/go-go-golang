package model

type Configuration struct {
	Database	map[string]string
	Redis		string
	ServerPort	string
}