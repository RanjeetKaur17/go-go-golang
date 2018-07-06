package controller

import (
	"go-go-golang/GoREST/model"
	"go-go-golang/GoREST/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddStudent(c *gin.Context){
	student := model.Student{}
	err := c.BindJSON(&student)
	rowsAffected, lastInsertedId, err := dao.AddStudent(student)
	exception := ""
	if err == nil && rowsAffected > 0 {
		student.StudentID = lastInsertedId
	} else if err == nil {
		exception = "Some Error Occured"
	}
	c.JSON(200, gin.H{"exception": exception, "data": student})
}

func UpdateStudent(c *gin.Context){
	student := model.Student{}
	err := c.BindJSON(&student)
	rowsAffected, lastInsertedId, err := dao.AddStudent(student)
	exception := ""
	if err == nil && rowsAffected > 0 {
		student.StudentID = lastInsertedId
	} else if err == nil {
		exception = "Some Error Occured"
	}
	c.JSON(200, gin.H{"exception": exception, "data": student})
}

func DeleteStudent(c *gin.Context){
	studentID, _ := strconv.ParseInt(c.Params.ByName("studentID"), 10, 64)
	rowsAffected, err := dao.DeleteStudent(studentID)
	exception := ""
	if err == nil && rowsAffected == 0 {
		exception = "No Data Found"
	} else if err == nil {
		exception = "Some Error Occured"
	} else if err != nil{
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception})
}

func GetStudent(c *gin.Context){
	studentID, _ := strconv.ParseInt(c.Params.ByName("studentID"), 10, 64)
	student, err := dao.GetStudent(studentID)
	exception := ""
	if err != nil{
		exception = err.Error()
	}
	c.JSON(200, gin.H{"exception": exception, "data": student})
}
