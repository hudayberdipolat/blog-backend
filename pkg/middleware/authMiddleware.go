package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	tokenKey "github.com/hudayberdipolat/blog-backend/pkg/generateToken"
	"net/http"
)

// Claims represents the JWT claims.
type Claims struct {
	UserID      int    `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	jwt.StandardClaims
}

// AuthMiddleware is the JWT authentication middleware.

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  "Unauthorized",
			"status": http.StatusUnauthorized,
		})
	}

	claims, err := verifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":  "Invalid token",
			"status": http.StatusUnauthorized,
		})
	}
	c.Locals("phone_number", claims.PhoneNumber)
	c.Locals("user_id", claims.UserID)
	return c.Next()
}

// verifyToken verifies the JWT token.
func verifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return tokenKey.JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
