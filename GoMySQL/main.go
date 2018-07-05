package main

import (
	"go-go-golang/GoMySQL/service"
	"fmt"
	"go-go-golang/GoMySQL/dao"
)

func main(){

	dao.InitializeMySQL()

	name := "JohnDoe"
	age := int8(12)

	student, err := service.AddStudent(name, age)
	if err != nil {
		fmt.Println("Adding Student Failed With Error : ", err.Error())
	} else {
		fmt.Println("Added Student Successfully : ", student)
	}

	student, err = service.GetStudent(student.StudentID)
	if err != nil {
		fmt.Println("Got Student Failed With Error : ", err.Error())
	} else {
		fmt.Println("Got Student Successfully : ", student)
	}

	name = "JaneDoe"
	age = 13
	student, err = service.UpdateStudent(student.StudentID, name, age)
	if err != nil {
		fmt.Println("Updating Student Failed With Error : ", err.Error())
	} else {
		fmt.Println("Updated Student Successfully : ", student)
	}

	_, err = service.DeleteStudent(student.StudentID)
	if err != nil {
		fmt.Println("Deleting Student Failed With Error : ", err.Error())
	} else {
		fmt.Println("Deleted Student Successfully : ", student)
	}

}

