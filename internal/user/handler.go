package user

import (
	"api-auth/internal/shared/errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req CreateUserRequest

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

func SearchById(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || id <= 0 {

		status, body := errors.ToHTTP(
			&errors.AppError{
				Code:    errors.CodeValidation,
				Message: "Id Inválido",
				Err:     err,
			},
		)
		c.JSON(status, body)
		return
	}

	resp, err := FindById(id)

	if err != nil {
		status, body := errors.ToHTTP(err)
		c.JSON(status, body)
		return
	}

	c.JSON(http.StatusOK, resp)
}
