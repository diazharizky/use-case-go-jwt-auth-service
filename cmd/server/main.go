package main

import (
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/app"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/server"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/services"
)

func main() {
	appCtx := &app.Context{}

	appCtx.GenerateJWTTokenService = services.NewGenerateJWTTokenService()
	appCtx.ValidateJWTTokenService = services.NewValidateJWTTokenService()

	server.Start(appCtx)
}
