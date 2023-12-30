package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"errors"
	"fmt"
)

func CreateUser(user types.User) (int64, error) {

	q, err := DB.Prepare("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)")

	if err != nil {
		return -1, err
	}

	defer q.Close()
	result, err := q.Exec(user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func GetUserById(id int) (types.User, error) {
	q, err := DB.Prepare("SELECT * FROM users WHERE id = ?")

	if err != nil {
		return types.User{}, err
	}

	defer q.Close()

	user := types.User{}
	err = q.QueryRow(id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)

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

	user := types.User{}
	err = q.QueryRow(email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)

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
		return users, errors.New("Error")
	}
	defer q.Close()

	rows, err := q.Query()

	if err != nil {
		return users, errors.New("Error")
	}

	defer rows.Close()

	for rows.Next() {
		user := types.User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return []types.User{}, errors.New("Error")
		}
		users = append(users, user)
	}

	return users, nil
}
