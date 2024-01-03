package models

import "database/sql"

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

type DbQuery interface {
	// Prepare(query string) (*sql.Stmt, error)
	Scan(row *sql.Row) error
	QueryRow() *sql.Row
}

type DbInsert interface {
	Exec(q *sql.Stmt) (sql.Result, error)
}

func BaseCreate(table string, fieldNames []string, query DbInsert) (int64, error) {

	fields, vPlacholders := BuildFields(fieldNames)

	q, err := DB.Prepare("INSERT INTO " + table + " (" + fields + ") VALUES (" + vPlacholders + ")")

	if err != nil {
		return -1, err
	}

	defer q.Close()
	result, err := query.Exec(q)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

// func BaseGet(table string, queryString string, dq DbQuery) error {
// 	q, err := DB.Prepare(queryString)

// 	if err != nil {
// 		return err
// 	}

// 	defer q.Close()

// 	err = dq.Scan(q.QueryRow())

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
