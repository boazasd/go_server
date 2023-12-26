package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func CreateUser(user User) (int64, error) {
	println(user.FirstName, user.LastName, user.Password, user.Email)

	q, err := DB.Prepare("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)")
	defer q.Close()

	if err != nil {
		return -1, err
	}

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

func GetUserById(id int) (User, error) {
	q, err := DB.Prepare("SELECT * FROM users WHERE id = ?")
	defer q.Close()

	if err != nil {
		return User{}, err
	}

	user := User{}
	err = q.QueryRow(id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	q, err := DB.Prepare("SELECT * FROM users WHERE email = ?")
	defer q.Close()

	if err != nil {
		return User{}, err
	}

	user := User{}
	err = q.QueryRow(email).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUsers(sort string, dir string) ([]User, error) {
	if sort == "" {
		sort = "id"
	}
	if dir == "" {
		dir = "ASC"
	}
	qString := fmt.Sprintf("SELECT * FROM users ORDER BY %s %s", sort, dir)

	q, err := DB.Prepare(qString)
	defer q.Close()

	users := []User{}

	if err != nil {
		return []User{}, errors.New("Error")
	}

	rows, err := q.Query()
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return []User{}, errors.New("Error")
		}
		users = append(users, user)
	}

	return users, nil
}
