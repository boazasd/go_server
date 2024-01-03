package utils

import (
	"errors"
	"fmt"
)

type QueryBuilder struct {
	Table       string
	queryString string
	err         error
}

func (qb QueryBuilder) Select(fields []string) QueryBuilder {
	fieldsStr := ""
	for _, field := range fields {
		if !SanitizeForDb(field, true) {
			qb.err = errors.New("field " + field + " not valid")
			return qb
		}
	}

	if len(fields) == 0 {
		fieldsStr += "*"
	} else {
		for i, field := range fields {
			if i == len(fields)-1 {
				fieldsStr += field
			} else {
				fieldsStr += field + ", "
			}
		}
	}
	qb.queryString += "SELECT " + fieldsStr + " FROM " + qb.Table

	return qb
}

func (qb QueryBuilder) Paginate(limit uint, offset uint) QueryBuilder {

	if limit == 0 {
		limit = 10
	}

	qb.queryString += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, limit*offset)
	return qb
}

func (qb QueryBuilder) Sort(sort string, dir string) QueryBuilder {
	if !(SanitizeForDb(sort, true) && SanitizeForDb(dir, true)) {
		qb.err = errors.New("params are not valid")
		return qb
	}

	if dir == "" {
		dir = "ASC"
	}

	if sort != "" {
		qb.queryString += fmt.Sprintf(" ORDER BY %s %s", sort, dir)
	}

	return qb
}

func (ab QueryBuilder) raw(query string) QueryBuilder {
	ab.queryString += " " + query
	return ab
}

func (qb QueryBuilder) Build() (string, error) {
	return qb.queryString, qb.err
}
