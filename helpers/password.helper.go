package helpers

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil
}

// convert string to int
func ConvertStringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		// handle error
		return 0
	}
	return i
}