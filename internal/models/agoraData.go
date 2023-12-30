package models

import (
	"bez/bez_server/internal/types"
	"time"
)

func CreateAgoraData(agoraData types.AgoraData) (int64, error) {
	println("CreateAgoraData insert")
	q, err := DB.Prepare("INSERT INTO agoraData (link, name, details, city, area, date, updatedAt, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		println(err.Error())
		return -1, err
	}

	defer q.Close()

	now := time.Now()

	result, err := q.Exec(
		agoraData.Link,
		agoraData.Name,
		agoraData.Details,
		agoraData.City,
		agoraData.Area,
		agoraData.Date,
		now,
		now,
	)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetAgoraDataByLink(link string) (types.AgoraData, error) {
	q, err := DB.Prepare("SELECT link FROM agoraData WHERE link = ?")

	if err != nil {
		return types.AgoraData{}, err
	}

	defer q.Close()

	agoraData := types.AgoraData{}
	err = q.QueryRow(link).Scan(&agoraData.Link)

	if err != nil {
		return types.AgoraData{}, err
	}

	return agoraData, nil
}
