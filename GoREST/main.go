package main

import (
	"go-go-golang/GoREST/dao"
	"github.com/gin-gonic/gin"
	"go-go-golang/GoREST/controller"
)

func main(){

	dao.InitializeMySQL()

	router := gin.Default()
	router.Use(LoggingMiddleware)

	studentRouterGroup := router.Group("student/v1")
	{
		studentRouterGroup.GET("/student/:studentID", controller.GetStudent)
		studentRouterGroup.POST("/student", controller.AddStudent)
		studentRouterGroup.PUT("/student", controller.UpdateStudent)
		studentRouterGroup.DELETE("/student/:studentID", controller.DeleteStudent)
	}

	router.Run(":8080")

}


func LoggingMiddleware(c *gin.Context) {

	reqElements := make(map[string]interface{})
	reqElements["path"] = c.Request.URL.Path
	println(reqElements)
	c.Next()
}
