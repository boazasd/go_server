package models

import (
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
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, firstName TEXT NOT NULL, lastName TEXT NOT NULL, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL)")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sessions (id INTEGER PRIMARY KEY, sessionId TEXT NOT NULL UNIQUE, userId INTEGER NOT NULL, expirationTime TIMESTAMP NOT NULL)")

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
		println("first user pass", pass)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		firstUser := User{
			FirstName: "super",
			LastName:  "admin",
			Email:     "boazprog@gmail.com",
			Password:  utils.HashAndSalt([]byte(pass)),
		}
		CreateUser(firstUser)
	}

	return nil
}

func CloseDatabase() {
	DB.Close()
}

type query struct {
	sort      string
	direction string
}
