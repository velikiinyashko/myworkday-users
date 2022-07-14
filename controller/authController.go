package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/velikiinyashko/myworkday/model"
	"github.com/velikiinyashko/myworkday/service"
)

type Auth interface {
	Auth(ctx *gin.Context)
}

type auth struct {
	ctx *gin.Context
}

func AuthService(ctx *gin.Context) Auth {
	return &auth{ctx: ctx}
}

func (a *auth) Auth(ctx *gin.Context) {

	userSrv := service.UserService(context.Background())

	var login model.UserAuthDTO
	if err := ctx.BindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := userSrv.Auth(context.Background(), &login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isauth": "true",
		"data":   time.Now(),
	})
}
