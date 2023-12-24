package utils

import (
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func IsValidField(field string) bool {
	valid := regexp.MustCompile("^[A-Za-z0-9_]+$")
	return valid.MatchString(field)
}

func HashAndSalt(str []byte) string {
	hash, err := bcrypt.GenerateFromPassword(str, 8)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// func Hash(str string) string {
// 	hash, err := hash.Hash(str)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return string(hash)
// }

func ComparePasswords(hashedPwd string, plainPwd []byte) bool { // Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
