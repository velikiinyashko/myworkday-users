package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/velikiinyashko/myworkday/internal/service/user"
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

	userSrv := user.UserService(context.Background())

	var login user.UserAuthDTO
	if err := ctx.BindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := userSrv.Auth(context.Background(), &login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"data":  time.Now(),
	})
}
