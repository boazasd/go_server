package utils

import "regexp"

func IsValidField(field string) bool {
	valid := regexp.MustCompile("^[A-Za-z0-9_]+$")
	return valid.MatchString(field)
}
