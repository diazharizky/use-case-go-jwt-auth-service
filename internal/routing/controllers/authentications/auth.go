package authentications

import (
	"log"
	"net/http"

	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/models"
	"github.com/gin-gonic/gin"
)

func (ctl controller) Auth(ctx *gin.Context) {
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"data": "Unable to parse username/password from the request headers",
			},
		)
		return
	}

	if username == "" || password == "" {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"data": "Username/password combination is incorrect",
			},
		)
		return
	}

	signedToken, err := ctl.appCtx.GenerateJWTTokenService.Call(
		models.UserData{
			Username: username,
		},
	)
	if err != nil {
		log.Printf("Error unable to generate token: %v", err)

		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"error": "Internal server error",
			},
		)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": signedToken,
	})
}
