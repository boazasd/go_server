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

func SetOrUpdateWishes(id int64, wish string) (types.Wishes, error) {
	wishes := models.IWishes{}
	wishRes, err := wishes.Upsert(types.Wishes{UserId: id, Wishes: wish})

	if err != nil {
		return types.Wishes{}, err
	}

	return wishRes, nil
}

func GetWishes(id int64) (types.Wishes, error) {
	wishes := models.IWishes{}
	wish, err := wishes.GetByUserId(id)

	if err != nil {
		return types.Wishes{}, err
	}

	return wish, nil
}
