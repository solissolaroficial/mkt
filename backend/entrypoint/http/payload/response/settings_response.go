package response

// ChangePasswordResponse representa a resposta de alteração de senha
type ChangePasswordResponse struct {
	Message string `json:"message"`
}

// SuccessChangePasswordResponse cria uma resposta de sucesso
func SuccessChangePasswordResponse() ChangePasswordResponse {
	return ChangePasswordResponse{
		Message: "Password changed successfully",
	}
}

// UpdateProfileResponse representa a resposta de atualização de perfil
type UpdateProfileResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

// SuccessUpdateProfileResponse cria uma resposta de sucesso
func SuccessUpdateProfileResponse(user UserResponse) UpdateProfileResponse {
	return UpdateProfileResponse{
		Message: "Profile updated successfully",
		Data:    user,
	}
}

// GetProfileResponse representa a resposta de busca de perfil
type GetProfileResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}

// SuccessGetProfileResponse cria uma resposta de sucesso
func SuccessGetProfileResponse(user UserResponse) GetProfileResponse {
	return GetProfileResponse{
		Message: "Profile retrieved successfully",
		Data:    user,
	}
}
