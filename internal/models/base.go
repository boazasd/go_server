package models

import (
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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, email TEXT, password TEXT)")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db.Close()
	return nil
}

func ConnectDatabse() error {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}
