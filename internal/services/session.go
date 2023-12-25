package services

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/utils"
	"time"

	"github.com/google/uuid"
)

var expirationTime = time.Hour
var refreshWindow = time.Minute * 20

func CreateSession(userId int64) (string, error) {
	uuid := uuid.New()
	session := models.Session{
		SessionId:      utils.Hash([]byte(uuid.String())),
		UserId:         userId,
		ExpirationTime: time.Now().Add(expirationTime),
	}

	if err := models.CreateSession(session); err != nil {
		return "", err
	}

	return uuid.String(), nil
}

func CheckSession(sessionId string) bool {
	hashed := utils.Hash([]byte(sessionId))
	session, err := models.GetSession(hashed)

	if err != nil {
		return false
	}

	if session.Id == "" {
		return false
	}

	if session.ExpirationTime.Before(time.Now()) {
		return false
	}

	RefreshSession(session)

	return true
}

func RefreshSession(session models.Session) error {
	session.ExpirationTime = time.Now().Add(expirationTime)
	if err := models.UpdateSession(session); err != nil {
		return err
	}

	return nil
}
