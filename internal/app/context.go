package app

import "github.com/diazharizky/use-case-go-jwt-auth-service/internal/models"

type Context struct {
	UserData                models.UserData
	GenerateJWTTokenService IGenerateJWTTokenService
	ValidateJWTTokenService IValidateJWTTokenService
}
