package models

import (
	"bez/bez_server/internal/types"
	"time"
)

func CreateAgoraData(agoraData types.AgoraData) (int64, error) {
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
	agoraData := types.AgoraData{}
	err := DB.Select(&agoraData, "SELECT link FROM agoraData WHERE link = ?", link)

	return agoraData, err
}
