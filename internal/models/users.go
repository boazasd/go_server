package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type UserDbInterface struct {
	types.User
}

func (UserDbInterface) defaultFields() []string {
	return []string{"firstName", "lastName", "email", "password", "roles"}
}

func (userInsert UserDbInterface) Exec(q *sql.Stmt) (sql.Result, error) {
	result, err := q.Exec(
		userInsert.FirstName,
		userInsert.LastName,
		userInsert.Email,
		userInsert.Password,
		strings.Join(userInsert.Roles, ";"),
	)

	return result, err
}

func (userQuery UserDbInterface) Scan(row *sql.Row) error {

	roles := ""
	err := row.Scan(
		&userQuery.Id,
		&userQuery.FirstName,
		&userQuery.LastName,
		&userQuery.Email,
		&userQuery.Password,
		&roles,
		&userQuery.CreatedAt,
		&userQuery.UpdatedAt,
	)
	userQuery.Roles = strings.Split(roles, ";")
	return err
}

func (UserDbInterface) Create(user types.User) (int64, error) {
	udi := UserDbInterface{user}
	userId, error := BaseCreate("users", udi.defaultFields(), &udi)
	if error != nil {
		return -1, error
	} else {
		return userId, nil
	}
}

func (UserDbInterface) GetById(id int) (types.User, error) {
	udi := UserDbInterface{}
	fieldNames, _ := BuildFields(udi.defaultFields())
	q, err := DB.Prepare("SELECT " + fieldNames + " FROM users WHERE id = ?")

	if err != nil {
		return types.User{}, err
	}
	defer q.Close()
	udi.Scan(q.QueryRow(id))

	return udi.User, err
}

func (UserDbInterface) GetByEmail(email string) (types.User, error) {
	udi := UserDbInterface{}
	fieldNames, _ := BuildFields(udi.defaultFields())
	q, err := DB.Prepare("SELECT " + fieldNames + " FROM users WHERE email = ?")

	if err != nil {
		return types.User{}, err
	}
	defer q.Close()

	udi.Scan(q.QueryRow(email))

	return udi.User, err
}

func (UserDbInterface) GetMany(sort string, dir string, limit uint, pageNumber uint) ([]types.User, error) {
	udi := UserDbInterface{}
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

	fieldNames, _ := BuildFields(udi.defaultFields())
	qString := "SELECT " + fieldNames + " FROM users"

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

		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&roles,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		user.Roles = strings.Split(roles, ";")
		if err != nil {
			return []types.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}
