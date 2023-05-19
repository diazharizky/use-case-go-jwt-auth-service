package services

import (
	"time"

	"github.com/diazharizky/use-case-go-jwt-auth-service/config"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type generateJWTTokenService struct {
}

func NewGenerateJWTTokenService() generateJWTTokenService {
	return generateJWTTokenService{}
}

func (generateJWTTokenService) Call(payload models.UserData) (*string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := models.JWTClaims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	secretKey := config.Global.GetString("app.secret")
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}
