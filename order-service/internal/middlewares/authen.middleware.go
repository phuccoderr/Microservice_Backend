package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"order-service/global"
	"order-service/pkg/response"
)

type CustomClaims struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func JWTMiddleware() gin.HandlerFunc {
	jwtSecret := []byte(global.Config.Jwt.Secret)
	return func(c *gin.Context) {
		tokenString, err := JWTGetToken(c)
		if err != nil {
			response.ErrorResponse(c, "token not valid", "Unauthorized", http.StatusUnauthorized)
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			response.ErrorResponse(c, "token not valid", "Unauthorized", http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

func JWTGetToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		global.Logger.Error("jwt get token fail")
		return "", errors.New("Jwt not found!")
	}

	return authHeader[7:], nil
}
