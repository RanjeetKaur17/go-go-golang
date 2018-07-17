package database

import (
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/dao/connectors"
)

//Implementation for UserSessionDAO
type UserSessionDAO struct {}

//Add User Session on Login
func (t *UserSessionDAO) AddUserSession(userSession model.UserSession) (bool, error) {
	sqlQuery := "INSERT INTO user_login(username, device_id, user_secure_key) VALUES ($1, $2, $3)"
	res, err := connectors.GetDatabaseConnection().Exec(sqlQuery, userSession.Username, userSession.DeviceId, userSession.UserSecureKey)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected > 0, err
}

//Delete User Session on Logout
func (t *UserSessionDAO) DeleteUserSession(userSession model.UserSession) (bool, error) {
	sqlQuery := "DELETE FROM user_login WHERE device_id = $1 AND user_secure_key = $2"
	res, err := connectors.GetDatabaseConnection().Exec(sqlQuery, userSession.DeviceId, userSession.UserSecureKey)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	return rowsAffected > 0, err
}

//Check if Provided session details are correct, for Authorizing
func (t *UserSessionDAO) IsAuthorized(userSession model.UserSession) (string, error) {
	sqlQuery := "SELECT username FROM user_login WHERE device_id = $1 AND user_secure_key = $2"
	rows, err := connectors.GetDatabaseConnection().Query(sqlQuery, userSession.DeviceId, userSession.UserSecureKey)
	if rows.Next() {
		var username string
		rows.Scan(&username)
		return username, err
	}
	return "", err
}