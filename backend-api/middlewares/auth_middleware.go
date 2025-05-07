package middlewares

import (
	"kanban-board/backend-api/config"
	"kanban-board/backend-api/helpers"
	"kanban-board/backend-api/structs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.GetEnv("JWT_SECRET", "secret_key"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
				Success: false,
				Message: "Token is required",
				Errors:  make(map[string]string),
			})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
				Success: false,
				Message: "Invalid Token",
				Errors:  helpers.TranslateErrorMessage(err),
			})
			c.Abort()
			return
		}

		c.Set("username", claims.Subject)

		c.Next()
	}
}
