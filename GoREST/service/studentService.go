package service

import (
	"go-go-golang/GoREST/model"
	"go-go-golang/GoREST/dao"
	"github.com/pkg/errors"
)

func AddStudent(name string, age int8) (model.Student, error) {
	student := model.Student{Name:name, Age:age}
	rowsAffected, lastInsertedId, err := dao.AddStudent(student)
	if err == nil && rowsAffected > 0 {
		student.StudentID = lastInsertedId
	}
	return student, err
}

func UpdateStudent(id int64, name string, age int8) (model.Student, error) {
	student := model.Student{StudentID:id, Name:name, Age:age}
	rowsAffected, err := dao.UpdateStudent(student)
	if err == nil && rowsAffected == 0 {
		err = errors.New("No Data Found")
	}
	return student, err
}

func DeleteStudent(studentID int64) (bool, error) {
	rowsAffected, err := dao.DeleteStudent(studentID)
	if err == nil && rowsAffected == 0 {
		err = errors.New("No Data Found")
	}
	if err != nil {
		return false, err
	}
	return true, err
}

func GetStudent(studentID int64) (model.Student, error) {
	return dao.GetStudent(studentID)
}
