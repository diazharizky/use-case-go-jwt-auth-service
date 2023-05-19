package app

import (
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/models"
)

type IGenerateJWTTokenService interface {
	Call(payload models.UserData) (tokenString *string, err error)
}

type IValidateJWTTokenService interface {
	Call(tokenString string) (isValid bool, claims *models.JWTClaims, err error)
}
