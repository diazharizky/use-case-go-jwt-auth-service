package middlewares

import (
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/app"
)

type middlewares struct {
	appCtx *app.Context
}

func New(appCtx *app.Context) middlewares {
	return middlewares{
		appCtx: appCtx,
	}
}
