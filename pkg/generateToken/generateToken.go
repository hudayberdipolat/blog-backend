package generateToken

import (
	"github.com/golang-jwt/jwt"
	"time"
)

var JwtKey = []byte("AIHnaSDFhs!@#mdfshdfoin_I(JNsd")

func GenerateToken(phoneNumber string, userId int) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["phone_number"] = phoneNumber
	claims["uid"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix() // Token expires in 12 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
