package routing

import (
	"fmt"
	"net/http"

	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/app"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/routing/controllers"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/routing/controllers/authentications"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/routing/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(appCtx *app.Context) (router *gin.Engine) {
	router = gin.Default()

	mwares := middlewares.New(appCtx)

	v1 := router.Group("v1")
	{
		v1.GET("/ping", controllers.Ping)

		auth := v1.Group("/authentications")
		{
			authController := authentications.NewController(appCtx)

			auth.POST("", authController.Auth)
		}

		v1.GET("/users", mwares.AuthorizationMiddleware, func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"message": fmt.Sprintf("Hello %s!", appCtx.UserData.Username),
			})
		})
	}

	return
}
