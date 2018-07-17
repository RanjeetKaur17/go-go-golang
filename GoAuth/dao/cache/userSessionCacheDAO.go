package cache

import (
	"go-go-golang/GoAuth/model"
	"go-go-golang/GoAuth/dao/connectors"
	"time"
)

//Implementation for UserSessionDAOInterface
type UserSessionCacheDAO struct {}

//Add User Session on Login
func (t *UserSessionCacheDAO) AddUserSession(userSession model.UserSession) (bool, error) {
	key := model.USER_SESSION_CACHE_KEY + model.UNDERSCORE + userSession.DeviceId + model.UNDERSCORE + userSession.UserSecureKey
	var username string
	err := connectors.GetCacheClient().Get(key).Scan(&username)
	if err != nil || username == "" {
		err = connectors.GetCacheClient().Set(key, userSession.Username, model.SESSION_TTL_HOUR * time.Hour).Err()
		if err == nil {
			return true, nil
		}
	}
	return false, err
}

//Delete User Session on Logout
func (t *UserSessionCacheDAO) DeleteUserSession(userSession model.UserSession) (bool, error) {
	key := model.USER_SESSION_CACHE_KEY + model.UNDERSCORE + userSession.DeviceId + model.UNDERSCORE + userSession.UserSecureKey
	var username string
	err := connectors.GetCacheClient().Get(key).Scan(&username)
	if err == nil && username != "" {
		err := connectors.GetCacheClient().Del(key).Err()
		if err == nil {
			return true, err
		}
	}
	return false, err
}

//Check if Provided session details are correct, for Authorizing
func (t *UserSessionCacheDAO) IsAuthorized(userSession model.UserSession) (string, error) {
	key := model.USER_SESSION_CACHE_KEY + model.UNDERSCORE + userSession.DeviceId + model.UNDERSCORE + userSession.UserSecureKey
	var username string
	err := connectors.GetCacheClient().Get(key).Scan(&username)
	if err == nil && username != "" {
		return username, err
	}
	return "", err
}