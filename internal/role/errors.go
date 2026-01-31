package role

import "api-auth/internal/shared/errors"

var (
	ErrIdUserNotFound = &errors.AppError{
		Code:    errors.CodeNotFound,
		Message: "Usuário não encontrado",
	}

	ErrRoleCreationFailed = &errors.AppError{
		Code:    errors.CodeNotFound,
		Message: "Não foi possível adicionar a Role",
	}
)
