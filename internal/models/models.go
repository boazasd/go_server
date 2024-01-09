package models

import (
	"bez/bez_server/internal/types"
	"bez/bez_server/internal/utils"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

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
		roles TEXT NOT NULL,
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
		category TEXT NOT NULL,
		middleCategory TEXT NOT NULL,
		subCategory TEXT NOT NULL,
		area TEXT NOT NULL,
		condition TEXT NOT NULL,
		image TEXT NOT NULL,
		processed   boolean,
		date      TIMESTAMP NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS agoraAgents (
		id INTEGER PRIMARY KEY, 
		userId INTEGER NOT NULL,
		userEmail STRING NOT NULL,
		searchTxt TEXT NOT NULL,
		category TEXT NOT NULL,
		subCategory TEXT NOT NULL,
		area TEXT NOT NULL,
		condition TEXT NOT NULL,
		withImage BOOLEAN NOT NULL,
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
	db, err := sqlx.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func CreateFirstUser() {
	um := IUser{}
	users, err := um.GetMany("", "", 10, 0)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if len(users) == 0 {
		// pass, err := utils.RandomString(10, "")
		pass := "12345678"

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		hashedPass, error := utils.HashAndSalt([]byte(pass))

		if error != nil {
			log.Println(err)
			os.Exit(1)
		}

		println("super admin password is:", pass)
		firstUser := types.User{
			FirstName: "super",
			LastName:  "admin",
			Email:     "boazprog@gmail.com",
			Password:  hashedPass,
			Roles:     "super",
		}
		um := IUser{}
		_, err = um.Create(firstUser)

		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

func CloseDatabase() {
	DB.Close()
}
