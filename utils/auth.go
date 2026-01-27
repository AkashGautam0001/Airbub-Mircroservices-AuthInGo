package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)

	if error != nil {
		fmt.Println("Error hashing password:", error)
		return "", error
	}

	return string(hash), nil

}

func CheckPasswordHash(plainPassword string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainPassword))
	fmt.Println("CheckPasswordHash", err)
	return err == nil
}
