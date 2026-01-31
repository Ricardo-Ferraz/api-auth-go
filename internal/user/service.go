package user

import (
	"api-auth/internal/security"
	"api-auth/internal/shared/database"
	"errors"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func Create(req CreateUserRequest) (*UserResponse, error) {
	hash, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, ErrUserCreationFailed
	}

	userModel := User{
		Username: req.Username,
		PassHash: hash,
	}

	err = database.DB.Create(&userModel).Error

	if err != nil {
		if isUniqueConstraintError(err) {
			return nil, ErrUsernameAlreadyExists
		}

		return nil, ErrUserCreationFailed
	}

	fmt.Println(userModel)

	resp := ToUserResponse(userModel)

	return &resp, nil
}

func isUniqueConstraintError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062 // Erro de campo/chave duplicado no mysql
	}
	return false
}
