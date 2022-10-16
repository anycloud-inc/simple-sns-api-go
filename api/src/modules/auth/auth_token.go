package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthToken string

func MakeAuthToken(userId int) (AuthToken, error) {
	claims := jwt.MapClaims{
		"user_id":  userId,
		"resource": "User",
		"exp":      time.Now().Add(time.Hour * 720).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// TODO: use env
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	return AuthToken(tokenString), err
}
