package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("JWT_SECRET")

func generate(email string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tkn, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tkn, nil
}

func verify(token string) error {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return secret, nil
	})

	if err != nil {
		return err
	}

	if !tkn.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
