package controller

import (
	"github.com/gin-gonic/gin"
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/service"
	"encoding/json"
	"errors"
)

type AuthControllerInterface interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	LogOut(c *gin.Context)
	CheckAuthorization(c *gin.Context)
	GetMessage(c *gin.Context)
}

type AuthController struct {
	Service service.AuthServiceInterface
}

// Sign Up New User
// @Tags Auth
// Sign Up New User godoc
// @Summary Sign Up New User
// @Description Sign Up New Users
// @Accept  json
// @Produce  json
// @Param credentials body model.Credentials true "Credentials"
// @Router /auth/signup [post]
// @Success 200 {object} model.Response
func (t *AuthController) SignUp(c *gin.Context){
	var credentials model.Credentials
	err := c.BindJSON(&credentials)
	exception := ""
	signUpSuccessful := false
	if err != nil {
		exception = err.Error()
	} else {
		signUpSuccessful, err = t.Service.SignUp(credentials)
		if err != nil {
			exception = err.Error()
		}
	}
	c.JSON(200, gin.H{"exception": exception, "data": signUpSuccessful})
}

// Log In For a User
// @Tags Auth
// Log In For a User godoc
// @Summary Log In For a User
// @Description Log In For a User and Create Session Details
// @Accept  json
// @Produce  json
// @Param credentials body model.Credentials true "Credentials"
// @Router /auth/login [post]
// @Success 200 {object} model.UserSession
func (t *AuthController) LogIn(c *gin.Context){
	var credentials model.Credentials
	err := c.BindJSON(&credentials)
	exception := ""
	var userSession model.UserSession
	if err != nil {
		exception = err.Error()
	} else {
		userSession, err = t.Service.LogIn(credentials)
		if err != nil {
			exception = err.Error()
		}
	}
	c.JSON(200, gin.H{"exception": exception, "data": userSession})
}

// Log Out a User
// @Tags Auth
// Log In For a User godoc
// @Summary Log In For a User
// @Description Log Out a User and End his Session
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Router /auth/logout [post]
// @Success 200 {object} model.Response
func (t *AuthController) LogOut(c *gin.Context){
	userSession := getUserSession(c)
	signUpSuccessful, err := t.Service.LogOut(userSession)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": signUpSuccessful})
}

func (t *AuthController) CheckAuthorization(c *gin.Context) {
	userSession := getUserSession(c)
	username, err := t.Service.IsAuthorized(userSession)
	if username == "" || err != nil {
		c.AbortWithError(401, errors.New("access denied"))
	}
}

// Verify User Session
// @Tags Auth
// Verify User Session
// @Summary Verify User Session
// @Description Verify User Session
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization"
// @Router /hello [get]
// @Success 200 {object} model.Response
func (t *AuthController) GetMessage(c *gin.Context) {
	userSession := getUserSession(c)
	text := "Hello "
	username, err := t.Service.IsAuthorized(userSession)
	if err != nil || username == "" {
		text += "Guest!"
	} else {
		text += username + "!"
	}
	c.JSON(200, gin.H{"data" : text})
}

func getUserSession(c *gin.Context) (model.UserSession){
	userSessionJsonStr := c.GetHeader("Authorization")
	var userSession model.UserSession
	json.Unmarshal([]byte(userSessionJsonStr), &userSession)
	return userSession
}
