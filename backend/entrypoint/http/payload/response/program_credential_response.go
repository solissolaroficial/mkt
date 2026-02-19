package response

import "time"

// ProgramCredentialResponse represents a credential in API responses
type ProgramCredentialResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	User      string    `json:"user"`
	Password  string    `json:"password"`
	Access    string    `json:"access"`
	Notes     string    `json:"notes"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CredentialsListResponse returns a list of credentials
func CredentialsListResponse(credentials []ProgramCredentialResponse) map[string]interface{} {
	return map[string]interface{}{
		"credentials": credentials,
	}
}

// CredentialResponse returns a single credential
func CredentialResponse(credential ProgramCredentialResponse) map[string]interface{} {
	return map[string]interface{}{
		"credential": credential,
	}
}
