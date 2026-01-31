package user

import "api-auth/internal/shared/errors"

var (
	ErrUserNotFound = &errors.AppError{
		Code:    errors.CodeNotFound,
		Message: "Não encontrado",
	}

	ErrUsernameAlreadyExists = &errors.AppError{
		Code:    errors.CodeConflict,
		Message: "Username já cadastrado",
	}

	ErrUserCreationFailed = &errors.AppError{
		Code:    errors.CodeInternal,
		Message: "não foi possível criar usuário",
	}
)
