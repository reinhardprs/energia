package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtLink struct {
}

type jwtCustomClaims struct {
	Email  string `json:"email"`
	UserID int    `json:"userID"`
	jwt.RegisteredClaims
}

func (JwtLink JwtLink) GenerateJWT(userID int, email string) (string, error) {
	claims := &jwtCustomClaims{
		email,
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return t, nil
}
