package role

import (
	"api-auth/internal/shared/database"
	"errors"

	"github.com/go-sql-driver/mysql"
)

func Create(req CreateRoleRequest) (*RoleResponse, error) {

	roleModel := Role{
		Name:   req.NameRole,
		UserId: req.IdUser,
	}

	err := database.DB.Create(&roleModel).Error

	if err != nil {
		if isMissingConstraintError(err) {
			return nil, ErrIdUserNotFound
		}

		return nil, ErrRoleCreationFailed
	}

	resp := toRoleResponse(roleModel)

	return &resp, nil
}

func isMissingConstraintError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1452 // Erro por não encontrar a chave estrangeria
	}
	return false
}
