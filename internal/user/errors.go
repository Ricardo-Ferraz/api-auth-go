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
		Message: "Não foi possível criar usuário",
	}

	ErrUserSearchFailed = &errors.AppError{
		Code:    errors.CodeInternal,
		Message: "Não foi possível buscar usuário",
	}
)
