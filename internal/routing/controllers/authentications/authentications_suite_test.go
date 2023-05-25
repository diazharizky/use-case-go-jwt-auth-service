package authentications_test

import (
	"testing"

	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/app"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/routing/controllers/authentications"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	testRouter *gin.Engine
	appCtx     *app.Context
)

var _ = BeforeSuite(func() {
	testRouter = gin.Default()

	controller := authentications.NewController(appCtx)
	testRouter.POST("", controller.Auth)
})

func TestAuthentications(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Authentications Controller Suite")
}
