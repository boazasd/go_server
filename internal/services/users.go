package services

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"errors"
)

func GetUser(id int) (types.User, error) {
	if id == 0 {
		return types.User{}, errors.New("invalid id")
	}

	user, err := models.GetUserById(id)
	return user, err
}

func CreateUser(user types.User) (int64, error) {
	userId, error := models.CreateUser(user)
	if error != nil {
		return -1, error
	} else {
		return userId, nil
	}
}

func GetUsers(sort string, dir string) ([]types.User, error) {
	users, err := models.GetMany(sort, dir, 10, 0)
	return users, err
}

func Login(email string, password string) (int64, error) {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return -1, err
	}

	if !utils.ComparePasswords(user.Password, []byte(password)) {
		return -1, errors.New("invalid password")
	}

	return user.Id, nil
}
