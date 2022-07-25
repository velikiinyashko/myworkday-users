package main

import (
	"github.com/gin-gonic/gin"
	"github.com/velikiinyashko/myworkday/internal/controller"
)

func main() {

	r := gin.Default()

	auth := controller.AuthService(&gin.Context{})

	v1 := r.Group("/v1")
	{
		v1.GET("/users", nil)
		v1.POST("/singin", auth.SingIn)
		v1.POST("/singup", auth.SingUp)
	}

	r.Run(":8080")
}
