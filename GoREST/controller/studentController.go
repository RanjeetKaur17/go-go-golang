package controller

import (
	"go-go-golang/GoREST/model"
	"go-go-golang/GoREST/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddStudent(c *gin.Context){
	student := model.Student{}
	err := c.BindJSON(&student)
	student, err = service.AddStudent(student.Name, student.Age)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": student})
}

func UpdateStudent(c *gin.Context){
	student := model.Student{}
	err := c.BindJSON(&student)
	student, err = service.UpdateStudent(student.StudentID, student.Name,  student.Age)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": student})
}

func DeleteStudent(c *gin.Context){
	studentID, _ := strconv.ParseInt(c.Params.ByName("studentID"), 10, 64)
	isDeleted, err := service.DeleteStudent(studentID)
	exception := ""
	if err != nil {
		exception = err.Error()
	}
	if !isDeleted {
		exception = "Student Could not be deleted"
	}
	c.JSON(200, gin.H{"exception": exception})
}

func GetStudent(c *gin.Context){
	studentID, _ := strconv.ParseInt(c.Params.ByName("studentID"), 10, 64)
	student, err := service.GetStudent(studentID)
	exception := ""
	if err != nil{
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": student})
}
