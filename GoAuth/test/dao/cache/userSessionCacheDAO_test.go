package cache

import (
	"go-go-golang/GoAuth/dao"
	"testing"
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/util"
	"go-go-golang/GoAuth/dao/cache"
	"go-go-golang/GoAuth/controller"
)

var userSessionDAO dao.UserSessionDAOInterface

func init() {
	controller.SetUp("../../../" + model.CONFIG_PATH)
	userSessionDAO = &cache.UserSessionCacheDAO{}
}

func TestAddUserSession(t *testing.T) {
	flag, err := userSessionDAO.AddUserSession(model.UserSession{Username: "username", UserSecureKey: util.GetUUID(), DeviceId:util.GetUUID()})
	if err != nil || !flag{
		println("TestAddUserSession Failed")
		t.Fail()
	}
	println("TestAddUserSession Successful")
}

func TestAddDuplicateUserSession(t *testing.T) {
	userSession := model.UserSession{Username: "username", UserSecureKey: util.GetUUID(), DeviceId:util.GetUUID()}
	flag, err := userSessionDAO.AddUserSession(userSession)
	flag, err = userSessionDAO.AddUserSession(userSession)
	if err == nil && flag {
		println("TestAddDuplicateUserSession Failed")
		t.Fail()
	}
	println("TestAddDuplicateUserSession Successful")
}

func TestDeleteUserSessionForValidKey(t *testing.T) {
	userSession := model.UserSession{Username: "username", UserSecureKey: util.GetUUID(), DeviceId:util.GetUUID()}
	flag, err := userSessionDAO.AddUserSession(userSession)
	flag, err = userSessionDAO.DeleteUserSession(userSession)
	if err != nil || !flag {
		println("TestDeleteUserSessionForValidKey Failed")
		t.Fail()
	}
	println("TestDeleteUserSessionForValidKey Successful")
}

func TestDeleteUserSessionForInValidKey(t *testing.T) {
	userSession := model.UserSession{Username: "username", UserSecureKey: util.GetUUID(), DeviceId:util.GetUUID()}
	flag, err := userSessionDAO.DeleteUserSession(userSession)
	if err == nil && flag {
		println("TestDeleteUserSessionForInValidKey Failed")
		t.Fail()
	}
	println("TestDeleteUserSessionForInValidKey Successful")
}

func TestIsAuthorizedForValidKey(t *testing.T) {
	userSession := model.UserSession{Username: "username", UserSecureKey: util.GetUUID(), DeviceId:util.GetUUID()}
	_, err := userSessionDAO.AddUserSession(userSession)
	username, err := userSessionDAO.IsAuthorized(userSession)
	if err != nil && username == "" {
		println("TestIsAuthorizedForValidKey Failed")
		t.Fail()
	}
	println("TestIsAuthorizedForValidKey Successful")
}

func TestIsAuthorizedForInValidKey(t *testing.T) {
	userSession := model.UserSession{Username: "username", UserSecureKey: util.GetUUID(), DeviceId:util.GetUUID()}
	username, err := userSessionDAO.IsAuthorized(userSession)
	if err == nil && username != "" {
		println("TestIsAuthorizedForInValidKey Failed")
		t.Fail()
	}
	println("TestIsAuthorizedForInValidKey Successful")
}

