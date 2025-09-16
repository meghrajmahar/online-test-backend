package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId string `json:"uuid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func Sign(userId, role string) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	claims := Claims{
		UserId: userId, Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
}
