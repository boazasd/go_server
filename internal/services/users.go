package services

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"errors"
)

func GetUser(id int64) (types.User, error) {
	if id == 0 {
		return types.User{}, errors.New("invalid id")
	}
	um := models.IUser{}
	user, err := um.GetById(id)
	return user, err
}

func CreateUser(user types.User) (int64, error) {
	um := models.IUser{}
	userId, error := um.Create(user)
	if error != nil {
		return -1, error
	} else {
		return userId, nil
	}
}

func GetUsers(sort string, dir string) ([]types.User, error) {
	um := models.IUser{}
	users, err := um.GetMany(sort, dir, 10, 0)
	return users, err
}

func Login(email string, password string) (int64, error) {
	um := models.IUser{}
	user, err := um.GetByEmail(email)
	if err != nil {
		return -1, err
	}

	if !utils.ComparePasswords(user.Password, []byte(password)) {
		return -1, errors.New("invalid password")
	}

	return user.Id, nil
}

func AddAgoraAgent(agent types.AgoraAgent) (types.AgoraAgent, error) {
	am := models.IAgoraAgents{}
	id, err := am.Create(agent)
	if err != nil {
		return types.AgoraAgent{}, err
	}

	res, err := am.GetById(id)
	if err != nil {
		return types.AgoraAgent{}, err
	}

	return res, nil
}

func GetAgoraAgents(id int64) ([]types.AgoraAgent, error) {
	wm := models.IAgoraAgents{}
	agents, err := wm.GetByUserId(id)

	if err != nil {
		return []types.AgoraAgent{}, err
	}

	return agents, nil
}
