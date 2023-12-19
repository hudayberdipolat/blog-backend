package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword)
}

func CheckPassword(userPassword, requestPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(requestPassword))
	if err != nil {
		return err
	}
	return nil
}
