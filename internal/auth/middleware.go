package auth

import (
	"api-auth/internal/shared/errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtSvc JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {

			status, body := errors.ToHTTP(ErrNoToken)

			c.AbortWithStatusJSON(status, body)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {

			status, body := errors.ToHTTP(ErrInvalidToken)

			c.AbortWithStatusJSON(status, body)
			return
		}

		claims, err := jwtSvc.ValidateAndExtract(parts[1])
		if err != nil {

			status, body := errors.ToHTTP(ErrInvalidToken)
			c.AbortWithStatusJSON(status, body)
			return
		}

		//Carrega a informação no contexto do Gin
		c.Set("roles", claims.Roles)
		c.Set("userID", claims.Subject)

		c.Next()
	}
}

func RequireRole(required string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesAny, exists := c.Get("roles")
		status, body := errors.ToHTTP(ErrNoPermission)

		if !exists {
			c.AbortWithStatusJSON(status, body)
			return
		}

		roles := rolesAny.([]string)

		for _, r := range roles {
			if r == required {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(status, body)
	}
}
