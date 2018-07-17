package database

import (
	"go-go-golang/GoAuth/dao"
	"testing"
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/util"
	"go-go-golang/GoAuth/dao/database"
	"go-go-golang/GoAuth/controller"
)

var userDAO dao.UserDAOInterface

func init() {
	controller.SetUp("../../../" + model.CONFIG_PATH)
	userDAO = &database.UserDAO{}
}

func TestAddCredentials(t *testing.T) {
	err := userDAO.AddCredentials(model.Credentials{Username: util.GetUUID(), Password:"password"})
	if err != nil {
		println("TestAddCredntials Failed")
		t.Fail()
	}
	println("TestAddCredntials Successful")
}

func TestGetPasswordForValidUser(t *testing.T) {
	credentials := model.Credentials{Username: util.GetUUID(), Password:"password"}
	err := userDAO.AddCredentials(credentials)
	password, err := userDAO.GetPassword(credentials.Username)
	if err != nil || password != credentials.Password {
		println("TestGetPasswordForValidUser Failed")
		t.Fail()
	}
	println("TestGetPasswordForValidUser Successful")
}

func TestGetPasswordForInValidUser(t *testing.T) {
	password, err := userDAO.GetPassword(util.GetUUID())
	if err != nil || password != "" {
		println("TestGetPasswordForInValidUser Failed")
		t.Fail()
	}
	println("TestGetPasswordForInValidUser Successful")
}

func TestUserNameExistsForValidUser(t *testing.T) {
	credentials := model.Credentials{Username: util.GetUUID(), Password:"password"}
	err := userDAO.AddCredentials(credentials)
	flag, err := userDAO.UserNameExists(credentials.Username)
	if err != nil || !flag {
		println("TestUserNameExistsForValidUser Failed")
		t.Fail()
	}
	println("TestUserNameExistsForValidUser Successful")
}

func TestUserNameExistsForInValidUser(t *testing.T) {
	flag, err := userDAO.UserNameExists(util.GetUUID())
	if err != nil || flag {
		println("TestUserNameExistsForInValidUser Failed")
		t.Fail()
	}
	println("TestUserNameExistsForInValidUser Successful")
}

