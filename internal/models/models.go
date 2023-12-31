package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func CreateDatabase() error {

	os.MkdirAll("./data", 0755)
	_, error := os.Stat("./data/data.db")
	if errors.Is(error, os.ErrNotExist) {
		fmt.Println("Creating database...")
		os.Create("./data/data.db")
	}

	db, err := sql.Open("sqlite3", "./data/data.db")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer db.Close()
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY, 
		firstName TEXT NOT NULL, 
		lastName TEXT NOT NULL, 
		email TEXT NOT NULL UNIQUE, 
		password TEXT NOT NULL,
		roles TEXT [],
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY, 
		sessionId TEXT NOT NULL UNIQUE, 
		userId INTEGER NOT NULL, 
		expirationTime TIMESTAMP NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS agoraData (
		id INTEGER PRIMARY KEY, 
		link      TEXT NOT NULL UNIQUE,
		name      TEXT NOT NULL,
		details   TEXT NOT NULL,
		city      TEXT NOT NULL,
		area      TEXT NOT NULL,
		date      TIMESTAMP NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}

func ConnectDatabse() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}

	DB = db

	users, err := GetUsers("", "", 10, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(users) == 0 {
		pass, err := utils.RandomString(10, "")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		hashedPass, error := utils.HashAndSalt([]byte(pass))

		if error != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		println("super admin password is:", pass)
		firstUser := types.User{
			FirstName: "super",
			LastName:  "admin",
			Email:     "boazprog@gmail.com",
			Password:  hashedPass,
			Roles:     []string{"super"},
		}
		u := UserModel{Entity: firstUser}

		_, err = Create("users", []string{"firstName", "lastName", "email", "password", "roles"}, &u)
		// _, err = CreateUser(firstUser)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
	}

	return nil
}

func CloseDatabase() {
	DB.Close()
}

// type query struct {
// 	sort      string
// 	direction string
// }
