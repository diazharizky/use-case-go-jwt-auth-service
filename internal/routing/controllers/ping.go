package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	respData := map[string]interface{}{
		"data": "pong",
	}

	ctx.JSON(http.StatusOK, respData)
}
