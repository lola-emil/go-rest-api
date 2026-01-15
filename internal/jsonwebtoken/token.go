package jsonwebtoken

import (
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(claims jwt.MapClaims) (*string, error) {
	signIngKey := []byte("my-secret-key")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signIngKey)

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
