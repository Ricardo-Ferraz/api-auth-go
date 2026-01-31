package user

import (
	"api-auth/internal/security"
	"api-auth/internal/shared/database"
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
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

	resp := ToUserResponse(userModel)

	return &resp, nil
}

func FindById(id int64) (*UserSearchResponse, error) {
	userModel := User{}

	err := database.DB.
		Debug().
		Where("id = ?", id).
		Preload("Roles").
		First(&userModel).Error

	if err != nil {
		if isNotFoundError(err) {
			return nil, ErrUserNotFound
		}

		return nil, ErrUserSearchFailed
	}
	resp := ToUserSearchResponse(userModel)

	return &resp, nil
}

func isNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func isUniqueConstraintError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062 // Erro de campo/chave duplicado no mysql
	}
	return false
}
