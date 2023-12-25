package models

import (
	"bez/bez_server/internal/utils"
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

func createSession(userId int64) (int, error) {

	tx, err := DB.Begin()
	if err != nil {
		return -1, err
	}

	uuid := uuid.New()
	session := Session{
		SessionId:      utils.Hash([]byte(uuid.String())),
		UserId:         userId,
		expirationTime: time.Now().Add(expirationTime),
	}

	q, err := tx.Prepare("INSERT INTO sessions (first_name, last_name, email, password) VALUES (?, ?, ?, ?)")
	print(session)
}

func getSession(sessionId string) Session {
	session := Session{
		SessionId: sessionId,
	}

	return session
}

func updateSession(sessionId string) {

}

func checkSession(sessionId string) bool {

	return true
}
