package dao

import "go-go-golang/GoAuth/model"

//Interface for methods available related to user credentials
type UserDAOInterface interface {

	//Check if username already exists
	UserNameExists(username string) (bool, error)

	//Add Credentials for new users, on signup
	AddCredentials(credentials model.Credentials) (error)

	//Get Encrypted password for given user name, to validate login
	GetPassword(username string) (string, error)
}

//Interface for methods available related to user session
type UserSessionDAOInterface interface {

	//Add User Session on Login
	AddUserSession(userSession model.UserSession) (bool, error)

	//Delete User Session on Logout
	DeleteUserSession(userSession model.UserSession) (bool, error)

	//Check if Provided session details are correct, for Authorizing
	IsAuthorized(userSession model.UserSession) (string, error)
}
