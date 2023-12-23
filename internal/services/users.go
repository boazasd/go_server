package services

import (
	"bez/bez_server/internal/models"
	"errors"
)

func GetUser(id int) (models.User, error) {
	if id == 0 {
		return models.User{}, errors.New("Invalid id")
	}

	user, err := models.GetUser(id)
	return user, err
}

func CreateUser(user models.User) (int64, error) {
	userId, error := models.CreateUser(user)
	if error != nil {
		return -1, error
	} else {
		return userId, nil
	}
}

func GetUsers(sort string, dir string) ([]models.User, error) {
	users, err := models.GetUsers(sort, dir)
	return users, err
}
