package service

import (
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/dao"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"go-go-golang/GoAuth/util"
)

//Interface for Services Available for Auth
type AuthServiceInterface interface {
	//Sign Up New User. Encrypt password and save into datastore.
	SignUp(credentials model.Credentials) (bool, error)

	//Log In For a User, validate if the password is same as encrypted password and return session details
	LogIn(credentials model.Credentials) (model.UserSession, error)

	//Log Out a User by terminating his session
	LogOut(userSession model.UserSession) (bool, error)

	//Verify if provided session details are valid. If valid return username
	IsAuthorized(userSession model.UserSession) (string, error)
}

//Implementation for AuthServiceInterface
type AuthService struct {
	UserDAO        dao.UserDAOInterface
	UserSessionDAO dao.UserSessionDAOInterface
}

//Sign Up New User. Encrypt password and save into datastore.
func (t *AuthService) SignUp(credentials model.Credentials) (bool, error) {

	if credentials.Username == "" || credentials.Password == "" {
		return false, errors.New("username/password cannot be empty")
	}

	userNameExists, err := t.UserDAO.UserNameExists(credentials.Username)
	if err != nil {
		return false, err
	} else if userNameExists {
		return false, errors.New("username already exists")
	} else {
		// Salt and hash the password using the bcrypt algorithm
		// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), 8)
		credentials.Password = string(hashedPassword)
		err = t.UserDAO.AddCredentials(credentials)
		if err != nil {
			return false, err
		}
	}
	return true, err
}


//Log In For a User, validate if the password is same as encrypted password and return session details
func (t *AuthService) LogIn(credentials model.Credentials) (model.UserSession, error) {

	if credentials.Username == "" || credentials.Password == "" {
		return model.UserSession{}, errors.New("username/password cannot be empty")
	}

	savedPassword, err := t.UserDAO.GetPassword(credentials.Username)
	if err != nil {
		return model.UserSession{}, err
	} else if err = bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(credentials.Password)); err != nil {
		return model.UserSession{}, errors.New("incorrect password")
	} else {
		userSession := model.UserSession{Username:credentials.Username, DeviceId:util.GetUUID(), UserSecureKey:util.GetUUID()}
		success, err := t.UserSessionDAO.AddUserSession(userSession)
		if err == nil && success{
			return userSession, nil
		} else {
			return model.UserSession{}, err
		}
	}
}

//Log Out a User by terminating his session
func (t *AuthService) LogOut(userSession model.UserSession) (bool, error) {

	if userSession.DeviceId == "" || userSession.UserSecureKey == "" {
		return false, errors.New("invalid credentials")
	}

	return t.UserSessionDAO.DeleteUserSession(userSession)
}

//Verify if provided session details are valid. If valid return username
func (t *AuthService) IsAuthorized(userSession model.UserSession) (string, error) {

	if userSession.DeviceId == "" || userSession.UserSecureKey == "" {
		return "", errors.New("invalid credentials")
	}

	return t.UserSessionDAO.IsAuthorized(userSession)
}