package authentications

import (
	"net/http"
	"time"

	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (controller) Auth(ctx *gin.Context) {
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			map[string]interface{}{
				"data": "Username/password combination is incorrect",
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

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := models.JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte("mysecretkey"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"message": "Error unable to authenticate client",
			},
		)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": signedToken,
	})
}
