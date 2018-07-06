package model

type Student struct {
	StudentID int64		`json:"id"`
	Name      string	`json:"name"`
	Age       int8		`json:"age"`
}
