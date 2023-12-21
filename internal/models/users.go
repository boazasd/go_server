package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

// func createUserTable() {
// 	res,err := DB.query("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, email TEXT, password TEXT, roles TEXT)")
// 	if err != nil {
// 		return err
// 	}
// }

func CreateUser(user User) (int64, error) {
	fmt.Println("Creating user...", user.FirstName, user.LastName, user.Email, user.Password)
	tx, err := DB.Begin()
	if err != nil {
		return -1, err
	}

	q, err := tx.Prepare("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)")

	if err != nil {
		return -1, err
	}

	defer q.Close()

	result, err1 := q.Exec(user.FirstName, user.LastName, user.Email, user.Password)

	if err1 != nil {
		return -1, err1
	}
	id, err2 := result.LastInsertId()

	if err2 != nil {
		return -1, err2
	}

	tx.Commit()

	return id, nil
}

func GetUser(id int) (User, error) {
	q, err := DB.Prepare("SELECT * FROM users WHERE id = ?")

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

func GetUsers() ([]User, error) {
	q, err := DB.Prepare("SELECT * FROM users")
	users := []User{}

	if err != nil {
		return []User{}, errors.New("Error")
	}

	rows, err := q.Query()

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
