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
	Prepare(query string) (DbQuery, error)
	Scan(dest ...interface{}) error
}

type InsertQuery interface {
	Insert(q *sql.Stmt) (sql.Result, error)
}

func Create(table string, fieldNames []string, query InsertQuery) (int64, error) {

	fields, vPlacholders := BuildFields(fieldNames)

	q, err := DB.Prepare("INSERT INTO " + table + " (" + fields + ") VALUES (" + vPlacholders + ")")

	if err != nil {
		return -1, err
	}

	defer q.Close()
	result, err := query.Insert(q)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}
