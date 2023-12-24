package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id             string
	SessionId      string
	UserId         int64
	expirationTime time.Time
}

var expirationTime = time.Hour
var refreshWindow = time.Minute * 20

func createSession(userId int64) {
	uuid := uuid.New()
	session := Session{
		SessionId:      uuid.String(),
		UserId:         userId,
		expirationTime: time.Now().Add(expirationTime),
	}

	print(session)
}
