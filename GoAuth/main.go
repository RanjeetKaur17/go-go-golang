package main

import (
	"go-go-golang/GoAuth/controller"
	"go-go-golang/GoAuth/service"
	"go-go-golang/GoAuth/dao/database"
	"go-go-golang/GoAuth/dao/cache"
	"go-go-golang/GoAuth/model"
	"github.com/gin-gonic/gin"
	"go-go-golang/GoAuth/util"
)

func main() {

	router := gin.Default()

	controller.SetUp(model.CONFIG_PATH)
	controller.StartRoutesV1(router, &controller.AuthController{
		Service:&service.AuthService{
			UserDAO:&database.UserDAO{},
			UserSessionDAO:&cache.UserSessionCacheDAO{}}})


	router.Run(":" + util.GetConfiguration().ServerPort)
}