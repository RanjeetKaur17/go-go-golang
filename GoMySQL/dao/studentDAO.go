package dao

import (
	"database/sql"
	"go-go-golang/GoMySQL/model"
)

func AddStudent(student model.Student) (int64, int64, error) {
	sqlQuery := "INSERT student SET name = ?, age = ?"
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(student.Name, student.Age)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	return rowsAffected, lastInsertedId, err
}

func UpdateStudent(student model.Student) (int64, error) {
	sqlQuery := "UPDATE student SET name = ?, age = ? WHERE id = ?"
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(student.Name, student.Age, student.StudentID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

func DeleteStudent(studentID int64) (int64, error) {
	sqlQuery := "DELETE FROM student WHERE id = ?"
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(studentID)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

func GetStudent(studentID int64) (model.Student, error) {
	sqlQuery := "SELECT id, name, age FROM student WHERE id = ?"
	stmt, err := GetMySQLConnection().Prepare(sqlQuery)
	defer closeStmt(stmt)
	var student model.Student
	if err != nil {
		return student, err
	}
	res, err := stmt.Query(studentID)
	defer closeRows(res)
	if err != nil {
		return student, err
	}
	if res.Next() {
		res.Scan(&student.StudentID, &student.Name, &student.Age)
	}
	return student, err
}

func closeRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}
