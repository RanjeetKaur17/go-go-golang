package service

import (
	"testing"
	"go-go-golang/GoAuth/service"
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/util"
	"go-go-golang/GoAuth/dao/database"
	"go-go-golang/GoAuth/dao/cache"
	"go-go-golang/GoAuth/controller"
)

var authService service.AuthServiceInterface

func init() {
	controller.SetUp("../../" + model.CONFIG_PATH)
	authService = &service.AuthService{UserDAO:&database.UserDAO{}, UserSessionDAO:&cache.UserSessionCacheDAO{}}
}

func TestSignUpWithEmptyUsername(t *testing.T) {
	flag, err := authService.SignUp(model.Credentials{Password:"password"})
	if err == nil || flag {
		println("TestSignUpWithEmptyUsername Failed")
		t.Fail()
	}
	println("TestSignUpWithEmptyUsername Successful")
}

func TestSignUpWithEmptyPassword(t *testing.T) {
	flag, err := authService.SignUp(model.Credentials{Username:util.GetUUID()})
	if err == nil || flag{
		println("TestSignUpWithEmptyPassword Failed")
		t.Fail()
	}
	println("TestSignUpWithEmptyPassword Successful")
}

func TestSignUp(t *testing.T) {
	flag, err := authService.SignUp(model.Credentials{Username:util.GetUUID(), Password:"password"})
	if err != nil || !flag{
		println("TestSignUp Failed")
		t.Fail()
	}
	println("TestSignUp Successful")
}

func TestLogInWithEmptyUserName(t *testing.T) {
	_, err := authService.LogIn(model.Credentials{Password:"password"})
	if err == nil{
		println("TestLogInWithEmptyUserName Failed")
		t.Fail()
	}
	println("TestLogInWithEmptyUserName Successful")
}

func TestLogInWithEmptyPassword(t *testing.T) {
	_, err := authService.LogIn(model.Credentials{Username:util.GetUUID()})
	if err == nil{
		println("TestLogInWithEmptyPassword Failed")
		t.Fail()
	}
	println("TestLogInWithEmptyPassword Successful")
}

func TestLogInWithoutSignUp(t *testing.T) {
	_, err := authService.LogIn(model.Credentials{Username:util.GetUUID(), Password:"password"})
	if err == nil{
		println("TestLogInWithoutSignUp Failed")
		t.Fail()
	}
	println("TestLogInWithoutSignUp Successful")
}

func TestLogInWithSignUp(t *testing.T) {
	credentials := model.Credentials{Username:util.GetUUID(), Password:"password"}
	authService.SignUp(credentials)
	_, err := authService.LogIn(credentials)
	if err != nil {
		println("TestLogInWithSignUp Failed")
		t.Fail()
	}
	println("TestLogInWithSignUp Successful")
}

func TestLogInUserSessionGeneration(t *testing.T) {
	credentials := model.Credentials{Username:util.GetUUID(), Password:"password"}
	authService.SignUp(credentials)
	userSession, err := authService.LogIn(credentials)
	if err != nil || userSession.DeviceId == "" || userSession.UserSecureKey == ""{
		println("TestLogInUserSessionGeneration Failed")
		t.Fail()
	}
	println("TestLogInUserSessionGeneration Successful")
}

func TestLogOutWithEmptyCredentials(t *testing.T) {
	flag, err := authService.LogOut(model.UserSession{})
	if err == nil || flag {
		println("TestLogOutWithEmptyCredentials Failed")
		t.Fail()
	}
	println("TestLogOutWithEmptyCredentials Successful")
}

func TestLogOutWithoutLogin(t *testing.T) {
	flag, err := authService.LogOut(model.UserSession{DeviceId:util.GetUUID(), UserSecureKey:util.GetUUID()})
	if err != nil && flag {
		println("TestLogOutWithoutLogin Failed")
		t.Fail()
	}
	println("TestLogOutWithoutLogin Successful")
}

func TestLogOutWithLogin(t *testing.T) {
	credentials := model.Credentials{Username:util.GetUUID(), Password:"password"}
	authService.SignUp(credentials)
	userSession, err := authService.LogIn(credentials)
	flag, err := authService.LogOut(userSession)
	if err != nil || !flag {
		println("TestLogOutWithoutLogin Failed")
		t.Fail()
	}
	println("TestLogOutWithoutLogin Successful")
}

func TestIsAuthorizedWithoutLogin(t *testing.T) {
	username, err := authService.IsAuthorized(model.UserSession{DeviceId:util.GetUUID(), UserSecureKey:util.GetUUID()})
	if err == nil && username == "" {
		println("TestIsAuthorizedWIthoutLogin Failed")
		t.Fail()
	}
	println("TestIsAuthorizedWIthoutLogin Successful")
}

func TestIsAuthorizedWithLogin(t *testing.T) {
	credentials := model.Credentials{Username:util.GetUUID(), Password:"password"}
	authService.SignUp(credentials)
	userSession, err := authService.LogIn(credentials)
	username, err := authService.IsAuthorized(userSession)
	if err != nil || username == "" {
		println("TestIsAuthorizedWithLogin Failed")
		t.Fail()
	}
	println("TestIsAuthorizedWithLogin Successful")
}