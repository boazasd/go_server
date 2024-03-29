package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"database/sql"
	"log"
)

type IAgoraModel struct {
}

func (*IAgoraModel) DefaultSelectFields() []string {
	return []string{
		"id",
		"link",
		"name",
		"details",
		"category",
		"middleCategory",
		"subCategory",
		"condition",
		"image",
		"processed",
		"city",
		"area",
		"date",
		"createdAt",
		"updatedAt",
	}
}

func (*IAgoraModel) CreateAgoraData(agoraData types.AgoraData) (int64, error) {
	fields, vPlacholders := BuildFields([]string{
		"link",
		"name",
		"details",
		"category",
		"middleCategory",
		"subCategory",
		"condition",
		"image",
		"processed",
		"city",
		"area",
		"date",
	})
	q, err := DB.Exec("INSERT INTO agoraData ("+fields+") VALUES ("+vPlacholders+")",
		agoraData.Link,
		agoraData.Name,
		agoraData.Details,
		agoraData.Category,
		agoraData.MiddleCategory,
		agoraData.SubCategory,
		agoraData.Condition,
		agoraData.Image,
		agoraData.Processed,
		agoraData.City,
		agoraData.Area,
		agoraData.Date,
	)

	if err != nil {
		println(err.Error())
		return -1, err
	}

	id, err := q.LastInsertId()

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
	query, err := qb.Select(am.DefaultSelectFields()).Sort(sort, dir).Paginate(limit, offset).Build()

	if err != nil {
		return agoraData, err
	}

	err = DB.Select(&agoraData, query)

	if err != nil {
		return agoraData, err
	}

	return agoraData, nil
}

func (am *IAgoraModel) UpdateProcessed(args ...interface{}) (sql.Result, error) {
	res, err := DB.Exec("UPDATE agoraData SET processed = true where processed = false")
	return res, err
}

func (am *IAgoraModel) GetForAgentMessage() ([]types.AgoraAgentResults, error) {
	qStr := `
	SELECT 

	agoraAgents.id as agentId,
	agoraAgents.userId,
	agoraAgents.userEmail,

	agoraData.link,
	agoraData.name,
	agoraData.details,
	agoraData.category,
	agoraData.middleCategory,
	agoraData.subCategory,
	agoraData.condition,
	agoraData.image,
	agoraData.area,
	agoraData.date

	FROM agoraData
	inner join agoraAgents
	on (
		agoraData.name like '%' || agoraAgents.searchTxt || '%'
		or
		agoraData.details like '%' || agoraAgents.searchTxt || '%'
	)
	and agoraAgents.category in (agoraData.category,"")
	and (
		agoraAgents.subCategory in (agoraData.middleCategory,"") 
		or 
		agoraAgents.subCategory in (agoraData.subCategory,"") 
	)
	and agoraAgents.condition in (agoraData.condition,"")
	and agoraAgents.area in (agoraData.area,"")
	and (
		agoraData.image != "" 
		or 
		agoraAgents.withImage = false
	)
	where agoraData.processed = false
	`

	uwas := []types.AgoraAgentResults{}
	err := DB.Select(&uwas, qStr)

	if err != nil {
		log.Println(err)
		return []types.AgoraAgentResults{}, err
	}

	return uwas, err
}
