package services

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/types"
)

func GetAgoraData(sort string, dir string, pageSize uint, pageNumber uint) ([]types.AgoraData, error) {
	am := models.IAgoraModel{}
	data, err := am.GetMany(sort, dir, pageSize, pageNumber)
	return data, err
}

func DeleteAgoraAgent(id int64) error {
	am := models.IAgoraAgents{}
	err := am.Delete(id)
	return err
}
