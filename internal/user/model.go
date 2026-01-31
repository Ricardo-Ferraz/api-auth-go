package user

import "api-auth/internal/role"

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex"`
	PassHash string
	Roles    []role.Role `gorm:"foreignKey:UserId"`
}

func (u *User) ReturnNameRoles() []string {
	roles := make([]string, 0, len(u.Roles))

	for _, role := range u.Roles {
		roles = append(roles, role.Name)
	}

	return roles
}
