package models

import (
	"time"
)

type Session struct {
	Id             string
	SessionId      string
	UserId         int64
	ExpirationTime time.Time
}

func CreateSession(session Session) error {

	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	q, err := tx.Prepare("INSERT INTO sessions (sessionId, userId, expirationTime) VALUES (?, ?, ?)")

	if err != nil {
		return err
	}

	defer q.Close()

	_, err = q.Exec(session.SessionId, session.UserId, session.ExpirationTime)

	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func GetSession(sessionId string) (Session, error) {
	q, err := DB.Prepare("SELECT * FROM sessions WHERE sessionId = ?")

	if err != nil {
		return Session{}, err
	}

	session := Session{}
	err = q.QueryRow(sessionId).Scan(&session.Id, &session.SessionId, &session.UserId, &session.ExpirationTime)

	if err != nil {
		return Session{}, err
	}

	return session, nil
}

func UpdateSession(session Session) error {
	q, err := DB.Prepare(`
	UPDATE sessions 
	SET expirationTime = $2
	WHERE id = $1
	`)

	if err != nil {
		return err
	}

	_, err = q.Exec(session.Id, session.ExpirationTime)

	if err != nil {
		return err
	}

	return nil
}

func DeleteSession(sessionId string) error {
	q, err := DB.Prepare(`
	DELETE sessions 
	WHERE id = $1
	`)

	if err != nil {
		return err
	}

	_, err = q.Exec(sessionId)

	if err != nil {
		return err
	}

	return nil
}
