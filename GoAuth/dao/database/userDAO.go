package database

import (
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/dao/connectors"
)

//Implementation for UserDAOInterface
type UserDAO struct {}

//Check if username already exists
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

//Add Credentials for new users, on signup
func (t *UserDAO) AddCredentials(credentials model.Credentials) (error) {
	sqlQuery := "INSERT INTO users(username, password) VALUES ($1, $2)"
	_, err := connectors.GetDatabaseConnection().Exec(sqlQuery, credentials.Username, credentials.Password)
	return err
}

//Get Encrypted password for given user name, to validate login
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