package models

import "bez/bez_server/internal/types"

func CreateSession(session types.Session) error {

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

func GetSession(sessionId string) (types.Session, error) {
	q, err := DB.Prepare("SELECT id, sessionId, userId, expirationTime FROM sessions WHERE sessionId = ?")

	if err != nil {
		return types.Session{}, err
	}

	session := types.Session{}
	err = q.QueryRow(sessionId).Scan(&session.Id, &session.SessionId, &session.UserId, &session.ExpirationTime)

	if err != nil {
		return types.Session{}, err
	}

	return session, nil
}

func UpdateSession(session types.Session) error {
	_, err := DB.Exec(`
	UPDATE sessions 
	SET expirationTime = $2
	WHERE id = $1
	`)

	return err
}

func DeleteSession(sessionId string) error {
	_, err := DB.Exec(`
	DELETE sessions 
	WHERE id = ?
	`, sessionId)

	return err
}
