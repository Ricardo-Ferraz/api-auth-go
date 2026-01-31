package role

//DTO de entrada para Criar uma role
type CreateRoleRequest struct {
	IdUser   int64  `json:"idUser" binding:"required"`
	NameRole string `json:"nameRole" binding:"required"`
}

//DTO de saida para criação de uma Role
type RoleResponse struct {
	Id       int64  `json:"id"`
	NameRole string `json:"nameRole"`
}
