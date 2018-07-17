package controller

import (
	"github.com/gin-gonic/gin"
	"go-go-golang/GoAuth/util"
	"go-go-golang/GoAuth/dao/connectors"
	_ "go-go-golang/GoAuth/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/gin-swagger"
)

// @APIVersion 1.0.0
// @APITitle Swagger Example API
// @APIDescription Swagger Example API
// @BasePath http://127.0.0.1:8080/
// @Contact ranjeet.17may@gmail.com
func StartRoutesV1(auth AuthControllerInterface) {

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	authGroup := router.Group("auth")
	{
		authGroup.POST("/signup", auth.SignUp)
		authGroup.POST("/login", auth.LogIn)
		authGroup.POST("/logout", auth.LogOut)
	}

	activityGroup := router.Group("").Use(auth.CheckAuthorization)
	{
		activityGroup.GET("/hello", auth.GetMessage)
	}

	router.Run(":" + util.GetConfiguration().ServerPort)
}

func SetUp(root string) {
	util.LoadConfigurations(root)
	connectors.InitializeDataBase(util.GetConfiguration().Database["Host"],
		util.GetConfiguration().Database["Port"],
		util.GetConfiguration().Database["Username"],
		util.GetConfiguration().Database["Password"],
		util.GetConfiguration().Database["Dbname"])

	connectors.InitializeCache("localhost:6379")

}