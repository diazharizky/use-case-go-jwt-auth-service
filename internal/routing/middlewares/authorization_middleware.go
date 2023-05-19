package middlewares

import (
	"log"
	"net/http"

	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/consts"
	"github.com/gin-gonic/gin"
)

func (mdw middlewares) AuthorizationMiddleware(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader(consts.AUTHORIZATION_HEADER_KEY)
	if authorizationHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid token",
		})
		return
	}

	tokenString := authorizationHeader[len(consts.BEARER_SCHEMA)+1:] // Add +1 to handle white space
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
