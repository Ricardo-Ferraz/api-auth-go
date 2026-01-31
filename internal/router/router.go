package router

import (
	"api-auth/internal/auth"
	"api-auth/internal/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	fmt.Println("Iniciando servidor...")

	r := gin.Default()

	auths := r.Group("/auth")
	{
		auths.POST("/register", user.Register)
		auths.POST("/login", auth.Login)
		auths.GET("/verify", auth.Verify)
	}

	return r
}
