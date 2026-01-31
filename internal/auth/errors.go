package auth

import "api-auth/internal/shared/errors"

var (
	ErrInvalidCredentials = &errors.AppError{
		Code:    errors.CodeInvalidCredentials,
		Message: "credenciais inválidas",
	}

	ErrNoPermission = &errors.AppError{
		Code:    errors.CodeErrorNoPermission,
		Message: "Sem permissão",
	}

	ErrNoToken = &errors.AppError{
		Code:    errors.CodeErrNoToken,
		Message: "Token não informado",
	}

	ErrInvalidToken = &errors.AppError{
		Code:    errors.CodeErrInvalidToken,
		Message: "Token inválido",
	}
)
