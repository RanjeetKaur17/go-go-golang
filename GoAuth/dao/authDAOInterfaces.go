package dao

import "go-go-golang/GoAuth/model"

type UserDAOInterface interface {
	UserNameExists(username string) (bool, error)
	AddCredentials(credentials model.Credentials) (error)
	GetPassword(username string) (string, error)
}

type UserSessionDAOInterface interface {
	AddUserSession(userSession model.UserSession) (bool, error)
	DeleteUserSession(userSession model.UserSession) (bool, error)
	IsAuthorized(userSession model.UserSession) (string, error)
}
