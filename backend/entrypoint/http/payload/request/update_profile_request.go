package request

// UpdateProfileRequest represents the request body for updating a user's profile
type UpdateProfileRequest struct {
	Name  string `json:"name" validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
	Role  string `json:"role" validate:"required,oneof=admin marketing commercial"`
}
