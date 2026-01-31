package auth

import (
	"api-auth/internal/shared/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		status, body := errors.ToHTTP(
			&errors.AppError{
				Code:    errors.CodeValidation,
				Message: "Payload Inválido",
				Err:     err,
			},
		)

		c.JSON(status, body)
		return
	}

	resp, err := Token(req)

	if err != nil {
		status, body := errors.ToHTTP(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Verify(c *gin.Context) {
	isValid := isMyToken(c.GetHeader("Authorization"))

	c.JSON(http.StatusOK, gin.H{
		"valid": isValid,
	})
}
