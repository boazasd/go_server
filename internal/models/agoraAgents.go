package models

import (
	"bez/bez_server/internal/types"
	"log"

	"github.com/mattn/go-sqlite3"
)

type IAgoraAgents struct {
}

func (*IAgoraAgents) DefaultSelectFields() string {
	return `
	id,
	userId, 
	searchTxt,
	category,
	subCategory,
	area,
	condition,
	onlyWithImage, 
	createdAt, 
	updatedAt
	`
}

func (*IAgoraAgents) Create(agent types.AgoraAgent) (int64, error) {
	fields, vPlacholders := BuildFields([]string{"userId", "searchTxt", "category", "subCategory", "area", "condition", "onlyWithImage"})
	res, err := DB.Exec("INSERT INTO agoraAgents ("+fields+") VALUES ("+vPlacholders+")",
		agent.UserId,
		agent.SearchTxt,
		agent.Category,
		agent.SubCategory,
		agent.Area,
		agent.Condition,
		agent.OnlyWithImage)

	if err != nil {
		// stliteErr, _ := err.(sqlite3.Error)
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			// if sqliteErr.Code == sqlite3.ErrBusy {
			// 	fmt.Println("busy")
			// }
			log.Println(sqliteErr.Error(), sqliteErr.Code)
		} else {
			log.Println("not sqlite error")
		}

		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (w *IAgoraAgents) GetById(id int64) (types.AgoraAgent, error) {
	agent := types.AgoraAgent{}
	err := DB.Get(&agent, "SELECT "+w.DefaultSelectFields()+" FROM agoraAgents WHERE id = ?", id)

	if err != nil {
		return types.AgoraAgent{}, err
	}

	return agent, nil
}

func (w *IAgoraAgents) GetByUserId(id int64) ([]types.AgoraAgent, error) {
	agents := []types.AgoraAgent{}
	err := DB.Select(&agents, "SELECT "+w.DefaultSelectFields()+" FROM agoraAgents WHERE userId = ?", id)

	if err != nil {
		return []types.AgoraAgent{}, err
	}

	return agents, nil
}

// func (w *IAgoraAgents) Update(id int64, agent types.AgoraAgents) (types.AgoraAgents, error) {
// 	agent.Id = id
// 	_, err := DB.NamedExec("UPDATE agoraAgents SET agoraAgents = :wishes WHERE id = :id", agent)

// 	if err != nil {
// 		return types.AgoraAgents{}, err
// 	}

// 	return agent, nil
// }
