package helpers

import (
	// "fmt"
	// "os"
	// "strconv"

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

//string to int
// func StringToInt(s string) int {
// 	i, err := strconv.Atoi(s)
// 	if err != nil {
// 		// handle error
// 		fmt.Println(err)
// 		os.Exit(2)
// 	}
// 	return i
// }