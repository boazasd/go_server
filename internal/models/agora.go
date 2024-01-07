package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"log"
	"strings"
	"time"
)

type IAgoraModel struct {
}

func (*IAgoraModel) DefaultSelectFields() string {
	return "id, link, name, details, city, area, date, createdAt, updatedAt"
}

func (*IAgoraModel) CreateAgoraData(agoraData types.AgoraData) (int64, error) {
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

func (*IAgoraModel) GetAgoraDataByLink(link string) (types.AgoraData, error) {
	agoraData := types.AgoraData{}
	err := DB.Get(&agoraData, "SELECT link FROM agoraData WHERE link = ?", link)

	return agoraData, err
}

func (am *IAgoraModel) GetMany(sort string, dir string, limit uint, offset uint) ([]types.AgoraData, error) {
	agoraData := []types.AgoraData{}
	qb := utils.QueryBuilder{Table: "agoraData"}
	query, err := qb.Select(strings.Split(am.DefaultSelectFields(), ", ")).Sort(sort, dir).Paginate(limit, offset).Build()
	log.Println(query)

	if err != nil {
		return agoraData, err
	}

	err = DB.Select(&agoraData, query)

	if err != nil {
		return agoraData, err
	}

	return agoraData, nil
}
