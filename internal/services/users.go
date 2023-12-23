package services

import (
	"bez/bez_server/internal/models"
	"bez/bez_server/internal/utils"
	"errors"
)

func GetUser(id int) (models.User, error) {
	if id == 0 {
		return models.User{}, errors.New("Invalid id")
	}

	user, err := models.GetUserById(id)
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

func Login(email string, password string) error {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return err
	}

	if !utils.ComparePasswords(user.Password, []byte(password)) {
		return errors.New("Invalid password")
	}

	return nil
}
