package user

import "api-auth/internal/role"

func ToUserResponse(u User) UserResponse {
	return UserResponse{
		Id:       u.Id,
		Username: u.Username,
	}
}

func ToUserSearchResponse(u User) UserSearchResponse {
	return UserSearchResponse{
		Id:       u.Id,
		Username: u.Username,
		Roles:    role.ToRoleResponseList(u.Roles),
	}
}
