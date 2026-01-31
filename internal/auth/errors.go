package auth

import "api-auth/internal/shared/errors"

var (
	ErrInvalidCredentials = &errors.AppError{
		Code:    errors.CodeInvalidCredentials,
		Message: "credenciais inválidas",
	}
)
