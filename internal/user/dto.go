package user

import "api-auth/internal/role"

//DTO de entrada para Criar um usuario
type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//DTO de saida na criação de um usuario
type UserResponse struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

//DTO de saída para busca de um usuario
type UserSearchResponse struct {
	Id       int64               `json:"id"`
	Username string              `json:"username"`
	Roles    []role.RoleResponse `json:"roles"`
}
