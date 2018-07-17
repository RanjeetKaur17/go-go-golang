package database

import (
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/dao/connectors"
)

type UserDAO struct {}

func (t *UserDAO) UserNameExists(username string) (bool, error) {
	sqlQuery := "SELECT username FROM users WHERE username = $1"
	rows, err := connectors.GetDatabaseConnection().Query(sqlQuery, username)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (t *UserDAO) AddCredentials(credentials model.Credentials) (error) {
	sqlQuery := "INSERT INTO users(username, password) VALUES ($1, $2)"
	_, err := connectors.GetDatabaseConnection().Exec(sqlQuery, credentials.Username, credentials.Password)
	return err
}

func (t *UserDAO) GetPassword(username string) (string, error) {
	sqlQuery := "SELECT password FROM users WHERE username = $1"
	rows, err := connectors.GetDatabaseConnection().Query(sqlQuery, username)
	if rows.Next() {
		var password string
		err = rows.Scan(&password)
		return password, err
	}
	return "", err
}