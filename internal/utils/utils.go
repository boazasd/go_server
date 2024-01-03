package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"log"
	"math/big"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func SanitizeForDb(field string, allowEmpty bool) bool {
	if allowEmpty && field == "" {
		return true
	}

	valid := regexp.MustCompile("^[A-Za-z0-9_]+$")
	return valid.MatchString(field)
}

func HashAndSalt(str []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(str, 8)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func RandomString(n int, letters string) (string, error) {
	if letters == "" {
		letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	}
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func Hash(str []byte) string {
	h := sha256.New()
	h.Write(str)
	hash := h.Sum(nil)
	return string(hash)
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	println("hashed", hashedPwd, plainPwd)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
