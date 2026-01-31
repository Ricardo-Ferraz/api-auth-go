package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToHTTP(err error) (int, any) {
	appErr, ok := err.(*AppError)
	if !ok {
		return http.StatusInternalServerError, gin.H{
			"code":    CodeInternal,
			"message": "erro interno",
		}
	}

	switch appErr.Code {
	case CodeNotFound:
		return http.StatusNotFound, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		}

	case CodeInvalidCredentials:
		return http.StatusUnauthorized, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		}

	case CodeConflict:
		return http.StatusConflict, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		}

	case CodeValidation:
		return http.StatusBadRequest, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		}

	case CodeErrorNoPermission:
		return http.StatusUnauthorized, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		}

	case CodeErrNoToken:
		return http.StatusUnauthorized, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		}

	case CodeErrInvalidToken:
		return http.StatusUnauthorized, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
		}

	}

	return http.StatusInternalServerError, gin.H{
		"code":    CodeInternal,
		"message": "erro interno",
	}
}
