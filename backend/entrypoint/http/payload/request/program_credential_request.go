package request

// CreateCredentialRequest represents the request body for creating a credential
type CreateCredentialRequest struct {
	Name     string `json:"name" validate:"required"`
	User     string `json:"user"`
	Password string `json:"password"`
	Access   string `json:"access"`
	Notes    string `json:"notes"`
}

// UpdateCredentialRequest represents the request body for updating a credential
type UpdateCredentialRequest struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Access   string `json:"access"`
	Notes    string `json:"notes"`
}
