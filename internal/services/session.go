package services

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var expirationTime = time.Hour

// var refreshWindow = time.Minute * 20

func CreateSessionCookie(sessionId string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = sessionId
	cookie.Expires = time.Now().Add(expirationTime)
	return cookie
}

func CreateOrRefreshSession(userId int64) (*http.Cookie, error) {
	uuid := uuid.New().String()
	session := types.Session{
		SessionId:      utils.Hash([]byte(uuid)),
		UserId:         userId,
		ExpirationTime: time.Now().Add(expirationTime),
	}

	if err := models.CreateSession(session); err != nil {
		return nil, err
	}

	cookie := CreateSessionCookie(uuid)

	return cookie, nil
}

func CheckSession(sessionId string) (bool, error) {
	hashed := utils.Hash([]byte(sessionId))
	session, err := models.GetSession(hashed)

	if err != nil {
		return false, err
	}

	if session.Id == "" {
		return false, err
	}

	if session.ExpirationTime.Before(time.Now()) {
		return false, errors.New("session expired")
	}

	return true, nil
}

func DeleteSession(sessionId string) error {
	err := models.DeleteSession(sessionId)
	if err != nil {
		return err
	}

	return nil
}
