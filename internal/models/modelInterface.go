package models

import (
	"database/sql"
)

type IDb interface {
	Exec(q *sql.Stmt) (sql.Result, error)
	// Scan(row *sql.Row) error
	// ScanMany(rows *sql.Rows) error
	DefaultFields() []string
	TableName() string
}

func BaseCreate(dbi IDb) (int64, error) {

	fields, vPlacholders := BuildFields(dbi.DefaultFields())

	q, err := DB.Prepare("INSERT INTO " + dbi.TableName() + " (" + fields + ") VALUES (" + vPlacholders + ")")

	if err != nil {
		return -1, err
	}

	defer q.Close()
	result, err := dbi.Exec(q)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func BaseGet(queryString string, dbi IDb, queryParams ...any) (*sql.Row, error) {
	q, err := DB.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	row := q.QueryRow(queryParams...)

	return row, nil
}

func BaseGetMany(queryString string, dq IDb, queryParams ...any) (*sql.Rows, error) {
	q, err := DB.Prepare(queryString)

	if err != nil {
		return nil, err
	}

	defer q.Close()

	rows, err := q.Query(queryParams...)

	return rows, nil
}

func BuildFields(fields []string) (string, string) {
	fieldNames := ""
	valuePlaceholders := ""
	for i, field := range fields {
		if i == len(fields)-1 {
			fieldNames += field
			valuePlaceholders += "?"
		} else {
			fieldNames += field + ", "
			valuePlaceholders += "?, "
		}
	}
	return fieldNames, valuePlaceholders
}
