package helpers

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(password string) string {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashPassword)
}
