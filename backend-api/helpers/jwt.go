package helpers

import (
	"kanban-board/backend-api/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

type MyCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, user_id int) string {
	exparationTime := time.Now().Add(60 * 24 * time.Minute)

	claims := MyCustomClaims{
		user_id,
		jwt.RegisteredClaims{
			Subject:   username,
			ExpiresAt: jwt.NewNumericDate(exparationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString(jwtKey)

	return signedToken
}
