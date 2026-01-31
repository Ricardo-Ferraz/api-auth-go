package user

func ToUserResponse(u User) UserResponse {
	return UserResponse{
		Id:       u.Id,
		Username: u.Username,
	}
}
