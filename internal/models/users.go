package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"strings"
)

type IUserModel struct {
}

func (IUserModel) DefaultSelectFields() string {
	return "id, firstName, lastName, email, password, roles, createdAt, updatedAt"
}

func (IUserModel) Create(user types.User) (int64, error) {
	fields, vPlacholders := BuildFields([]string{"firstName", "lastName", "email", "password", "roles"})
	res, err := DB.Exec("INSERT INTO users ("+fields+") VALUES ("+vPlacholders+")", user.FirstName, user.LastName, user.Email, user.Password, user.Roles)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (um IUserModel) GetById(id int) (types.User, error) {
	user := types.User{}
	err := DB.Get(&user, "SELECT "+um.DefaultSelectFields()+" FROM users WHERE id = ?", id)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (um IUserModel) GetByEmail(email string) (types.User, error) {
	user := types.User{}
	err := DB.Get(&user, "SELECT id, "+um.DefaultSelectFields()+" FROM users WHERE email = ?", email)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (um IUserModel) GetMany(sort string, dir string, limit uint, offset uint) ([]types.User, error) {
	users := []types.User{}
	qb := utils.QueryBuilder{Table: "users"}
	query, err := qb.Select(strings.Split(um.DefaultSelectFields(), ", ")).Paginate(limit, offset).Sort(sort, dir).Build()

	if err != nil {
		return []types.User{}, err
	}

	err = DB.Select(&users, query)

	if err != nil {
		return []types.User{}, err
	}

	return users, nil
}
