package role

import (
	"api-auth/internal/shared/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req CreateRoleRequest

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

	resp, err := Create(req)

	if err != nil {
		status, body := errors.ToHTTP(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, resp)
}
