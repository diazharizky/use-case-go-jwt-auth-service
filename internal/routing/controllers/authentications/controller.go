package authentications

import (
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/app"
)

type controller struct {
	appCtx *app.Context
}

func NewController(appCtx *app.Context) controller {
	return controller{
		appCtx: appCtx,
	}
}
