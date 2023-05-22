package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/consts"
	"github.com/gin-gonic/gin"
)

func (mdw middlewares) AuthorizationMiddleware(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader(consts.AuthorizationHeaderKey)
	if authorizationHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid token",
		})
		return
	}

	isValidAuthType := strings.Contains(authorizationHeader, consts.BearerSchema)
	if !isValidAuthType {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid authorization type",
		})
		return
	}

	tokenString := authorizationHeader[len(consts.BearerSchema)+1:] // Add +1 to handle white space
	valid, claims, err := mdw.appCtx.ValidateJWTTokenService.Call(tokenString)
	if err != nil {
		log.Printf("Error unable to validate token: %v", err)

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error unable to validate token",
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid token",
		})
		return
	}

	mdw.appCtx.UserData = claims.Payload
}
