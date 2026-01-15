package request

// CreateRepresentativeRequest represents a request to create a representative
type CreateRepresentativeRequest struct {
	Code      int    `json:"code" validate:"required,min=100,max=999"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
	Company   string `json:"company" validate:"required"`
	Region    string `json:"region" validate:"required"`
	City      string `json:"city"`
	Attendant string `json:"attendant" validate:"required"`
}

// UpdateRepresentativeRequest represents a request to update a representative
type UpdateRepresentativeRequest struct {
	Name      *string `json:"name" validate:"omitempty"`
	Email     *string `json:"email" validate:"omitempty,email"`
	Phone     *string `json:"phone" validate:"omitempty"`
	Company   *string `json:"company" validate:"omitempty"`
	Region    *string `json:"region" validate:"omitempty"`
	City      *string `json:"city" validate:"omitempty"`
	Attendant *string `json:"attendant" validate:"omitempty"`
	Active    *bool   `json:"active"`
}

// ListRepresentativesRequest represents a request to list representatives
type ListRepresentativesRequest struct {
	Page      int     `query:"page" validate:"omitempty,min=1"`
	PageSize  int     `query:"pageSize" validate:"omitempty,min=1,max=100"`
	Name      *string `query:"name"`
	Company   *string `query:"company"`
	Email     *string `query:"email"`
	Region    *string `query:"region"`
	City      *string `query:"city"`
	Active    *bool   `query:"active"`
	Code      *int    `query:"code" validate:"omitempty,min=100,max=999"`
	SortBy    string  `query:"sortBy" validate:"omitempty,oneof=name email company region city code createdAt updatedAt"`
	SortOrder string  `query:"sortOrder" validate:"omitempty,oneof=asc desc ASC DESC"`
}
