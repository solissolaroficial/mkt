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
