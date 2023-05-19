package services

import (
	"fmt"

	"github.com/diazharizky/use-case-go-jwt-auth-service/config"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type validateJWTTokenService struct{}

func NewValidateJWTTokenService() validateJWTTokenService {
	return validateJWTTokenService{}
}

func (svc validateJWTTokenService) Call(tokenString string) (bool, *models.JWTClaims, error) {
	claims := &models.JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error unexpected signing method: %v", token.Header["alg"])
		}

		secretKey := config.Global.GetString("app.secret")
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, nil, fmt.Errorf("error unable to validate token: %v", err)
	}

	if !token.Valid {
		return false, nil, nil
	}

	return true, claims, nil
}
