package middleware

import (
	"cart-service/internal/constants"
	"cart-service/internal/dto"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var jwtSecret = []byte("vVarX3ETLuR35pAe8LLVSEieaIxvBrz6X2B0eiN1HY4cdf3jYwBUKISJhDDXD60gsZiL9HLTYPoVwrSGa628XGmjJkGF04J3f4On")

type CustomClaims struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := JWTGetToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				dto.BuildResponseError(constants.JWT_NOT_FOUND, "Unauthorized", http.StatusUnauthorized))
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				dto.BuildResponseError(constants.JWT_TOKEN_INVALID, "Unauthorized", http.StatusUnauthorized))
			return
		}

		c.Next()
	}
}

func JWTGetToken(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New(constants.JWT_NOT_FOUND)
	}

	return authHeader[7:], nil
}

func JWTDecodeToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(constants.JWT_SIGNING_KEY)
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New(constants.JWT_TOKEN_INVALID)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if ok {
		return claims, nil
	}
	return nil, errors.New(constants.JWT_PARSE_CLAIMS)
}

func JWTGetCustomerID(c *gin.Context) (string, error) {
	token, _ := JWTGetToken(c)
	decodeToken, err := JWTDecodeToken(token)
	if err != nil {
		return "", err
	}

	return decodeToken.ID, nil
}
