package routing

import (
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/routing/controllers"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/routing/controllers/authentications"
	"github.com/gin-gonic/gin"
)

func NewRouter() (router *gin.Engine) {
	router = gin.Default()

	v1 := router.Group("v1")
	{
		v1.GET("/ping", controllers.Ping)

		auth := v1.Group("/authentications")
		{
			authController := authentications.NewController()
			auth.POST("", authController.Auth)
		}
	}

	return
}
