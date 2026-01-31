package role

func toRoleResponse(r Role) RoleResponse {
	return RoleResponse{
		Id:       r.Id,
		NameRole: r.Name,
	}
}

func ToRoleResponseList(roles []Role) []RoleResponse {
	responses := make([]RoleResponse, 0, len(roles))

	for _, r := range roles {
		responses = append(responses, RoleResponse{
			Id:       r.Id,
			NameRole: r.Name,
		})
	}

	return responses
}
