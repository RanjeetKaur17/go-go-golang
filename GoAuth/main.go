package main

import (
	"go-go-golang/GoAuth/controller"
	"go-go-golang/GoAuth/service"
	"go-go-golang/GoAuth/dao/database"
	"go-go-golang/GoAuth/dao/cache"
	"go-go-golang/GoAuth/model"
)

func main() {

	controller.SetUp(model.CONFIG_PATH)
	controller.StartRoutesV1(&controller.AuthController{
		Service:&service.AuthService{
			UserDAO:&database.UserDAO{},
			UserSessionDAO:&cache.UserSessionCacheDAO{}}})

}