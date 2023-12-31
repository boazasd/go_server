package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type UserModel struct {
	Entity types.User
}

func (model *UserModel) Insert(q *sql.Stmt) (sql.Result, error) {
	result, err := q.Exec(
		model.Entity.FirstName,
		model.Entity.LastName,
		model.Entity.Email,
		model.Entity.Password,
		strings.Join(model.Entity.Roles, ";"),
	)

	return result, err
}

func GetUserById(id int) (types.User, error) {
	q, err := DB.Prepare("SELECT * FROM users WHERE id = ?")

	if err != nil {
		return types.User{}, err
	}

	defer q.Close()

	user := types.User{}
	roles := ""
	createdAt := 0
	updatedAt := 0
	err = q.QueryRow(id).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&roles,
		&createdAt,
		&updatedAt,
	)
	user.Roles = strings.Split(roles, ";")

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (types.User, error) {
	q, err := DB.Prepare("SELECT * FROM users WHERE email = ?")

	if err != nil {
		return types.User{}, err
	}

	defer q.Close()
	roles := ""
	createdAt := 0
	updatedAt := 0
	user := types.User{}

	err = q.QueryRow(email).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&roles,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return types.User{}, err
	}

	return user, nil
}

func GetUsers(sort string, dir string, limit uint, pageNumber uint) ([]types.User, error) {
	users := []types.User{}

	if !(utils.SanitizeForDb(sort, true) && utils.SanitizeForDb(dir, true)) {
		return users, errors.New("params are not valid")
	}

	if dir == "" {
		dir = "ASC"
	}
	if limit == 0 {
		limit = 10
	}

	qString := "SELECT * FROM users"

	if sort != "" {
		qString += fmt.Sprintf(" ORDER BY %s %s", sort, dir)
	}

	qString += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, limit*pageNumber)
	q, err := DB.Prepare(qString)

	if err != nil {
		return users, err
	}
	defer q.Close()

	rows, err := q.Query()

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		user := types.User{}
		roles := ""
		createdAt := 0
		updatedAt := 0

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&roles,
			&createdAt,
			&updatedAt,
		)

		user.Roles = strings.Split(roles, ";")
		if err != nil {
			return []types.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
