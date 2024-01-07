package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"log"
	"strings"
)

type IUser struct {
}

func (*IUser) DefaultSelectFields() string {
	return "id, firstName, lastName, email, password, roles, createdAt, updatedAt"
}

func (*IUser) Create(user types.User) (int64, error) {
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

func (um *IUser) GetById(id int64) (types.User, error) {
	user := types.User{}
	err := DB.Get(&user, "SELECT "+um.DefaultSelectFields()+" FROM users WHERE id = ?", id)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (um *IUser) GetByEmail(email string) (types.User, error) {
	user := types.User{}
	err := DB.Get(&user, "SELECT id, "+um.DefaultSelectFields()+" FROM users WHERE email = ?", email)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func (um *IUser) GetMany(sort string, dir string, limit uint, offset uint) ([]types.User, error) {
	users := []types.User{}
	qb := utils.QueryBuilder{Table: "users"}
	query, err := qb.Select(strings.Split(um.DefaultSelectFields(), ", ")).Sort(sort, dir).Paginate(limit, offset).Build()
	log.Println(query)
	if err != nil {
		return []types.User{}, err
	}

	err = DB.Select(&users, query)

	if err != nil {
		return []types.User{}, err
	}

	return users, nil
}

func (um *IUser) GetForAgentMessage(agoraData []types.AgoraData) ([]types.UserWithAgant, error) {
	stmnt := `SELECT users.id, users.email, users.firstName, users.lastName, agoraAgents.searchTxt
	FROM users 
	inner join agoraAgents
	on users.id = agoraAgents.userId where`

	for _, ag := range agoraData {
		ok := utils.SanitizeForDb(ag.Name, true)
		if ok {
			stmnt += " agoraAgents.searchTxt LIKE '%" + ag.Name + "%' OR"
		}
	}
	stmnt = stmnt[:len(stmnt)-2]

	println(stmnt)

	uwas := []types.UserWithAgant{}
	err := DB.Select(&uwas, stmnt)

	return uwas, err
}
