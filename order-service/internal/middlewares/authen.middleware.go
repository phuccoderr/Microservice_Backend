package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net/http"
	"order-service/global"
	"order-service/internal/constants"
	"order-service/pkg/response"
)

type CustomClaims struct {
	ID    string   `json:"_id"`
	Email string   `json:"email"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

func JWTMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, customer, err := JWTGetTokenAndCustomer(c)
		if err != nil {
			response.ErrorResponse(c, err.Error(), constants.STATUS_UNAUTHORIZED, http.StatusUnauthorized)
			return
		}

		// Role-based authorization
		if !hasRequiredRoles(customer.Roles, requiredRoles...) {
			response.ErrorResponse(c, "forbidden", constants.STATUS_FORBIDDEN, http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func JWTGetTokenAndCustomer(c *gin.Context) (string, *CustomClaims, error) {
	token, err := JWTGetToken(c)
	if err != nil {
		return "", nil, err
	}

	customer, err := JWTDecodeToken(token)
	if err != nil {
		return "", nil, err
	}
	return token, customer, nil

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
			global.Logger.Error("decode token!")
			return nil, errors.New(constants.JWT_SIGNING_KEY)
		}
		return []byte(global.Config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		global.Logger.Error("Jwt token invalid", zap.Error(err))
		return nil, errors.New(constants.JWT_TOKEN_INVALID)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if ok {
		return claims, nil
	}

	global.Logger.Error("Jwt parse customer error", zap.Error(errors.New(constants.JWT_PARSE_CLAIMS)))
	return nil, errors.New(constants.JWT_PARSE_CLAIMS)
}

func hasRequiredRoles(userRoles []string, requiredRoles ...string) bool {
	roleMap := make(map[string]bool)
	for _, role := range requiredRoles {
		roleMap[role] = true
	}
	for _, role := range userRoles {
		if !roleMap[role] {
			return false
		}
	}
	return true
}
