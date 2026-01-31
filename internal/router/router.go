package router

import (
	"api-auth/internal/auth"
	"api-auth/internal/role"
	"api-auth/internal/user"

	"api-auth/internal/version"
	"fmt"

	"github.com/gin-gonic/gin"
)

var rolePrin = "ROLE_0"

func SetupRouter(jwt auth.JWTService) *gin.Engine {

	fmt.Println("Iniciando servidor...")

	r := gin.Default()

	auths := r.Group("/auth")
	{
		auths.POST("/register", user.Register)
		auths.POST("/login", auth.Login)
		auths.GET("/verify", auth.Verify)
	}

	roles := r.Group("/role", auth.AuthMiddleware(jwt), auth.RequireRole(rolePrin))
	roles.POST("/register", role.Register)

	users := r.Group("/users", auth.AuthMiddleware(jwt), auth.RequireRole(rolePrin))
	users.GET("/:id", user.SearchById)

	r.GET("/version", version.VersionHandler)

	return r
}
