package request

// CreatePdvPostRequest representa os dados para criar um post de PDV
type CreatePdvPostRequest struct {
	RepName  string  `json:"rep_name" validate:"required,max=100"`
	PdvName  string  `json:"pdv_name" validate:"required,max=200"`
	PostDate string  `json:"post_date" validate:"required,datetime=2006-01-02"`
	Platform string  `json:"platform" validate:"required,oneof=instagram facebook linkedin tiktok"`
	Link     *string `json:"link" validate:"omitempty,url,max=500"`
	ProofUrl *string `json:"proof_url" validate:"omitempty,url,max=500"`
}

// UpdatePdvPostStatusRequest representa os dados para atualizar o status de um post de PDV
type UpdatePdvPostStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=pending approved rejected published cancelled"`
}

// UpdatePdvPostRequest representa os dados para atualizar um post de PDV
type UpdatePdvPostRequest struct {
	RepName  *string `json:"rep_name" validate:"omitempty,max=100"`
	PdvName  *string `json:"pdv_name" validate:"omitempty,max=200"`
	PostDate *string `json:"post_date" validate:"omitempty,datetime=2006-01-02"`
	Platform *string `json:"platform" validate:"omitempty,oneof=instagram facebook linkedin youtube tiktok"`
	Link     *string `json:"link" validate:"omitempty,url,max=500"`
	ProofUrl *string `json:"proof_url" validate:"omitempty,url,max=500"`
}

// CreateRecurrentPdvRequest representa os dados para criar um PDV recorrente
type CreateRecurrentPdvRequest struct {
	Name    string  `json:"name" validate:"required,max=200"`
	RepName string  `json:"rep_name" validate:"required,max=100"`
	City    *string `json:"city" validate:"omitempty,max=100"`
}

// UpdateRecurrentPdvRequest representa os dados para atualizar um PDV recorrente
type UpdateRecurrentPdvRequest struct {
	Name             *string `json:"name" validate:"omitempty,max=200"`
	RepName          *string `json:"rep_name" validate:"omitempty,max=100"`
	City             *string `json:"city" validate:"omitempty,max=100"`
	Followers        *int    `json:"followers" validate:"omitempty,min=0"`
	InstagramProfile *string `json:"instagram_profile" validate:"omitempty,max=100"`
}

// ListPdvPostsQuery representa os parâmetros de query para listar posts de PDV
type ListPdvPostsQuery struct {
	RepName   *string `query:"rep_name" validate:"omitempty,max=100"`
	Month     *string `query:"month" validate:"omitempty,oneof=JAN FEV MAR ABR MAI JUN JUL AGO SET OUT NOV DEZ"`
	Platform  *string `query:"platform" validate:"omitempty,oneof=instagram facebook linkedin youtube tiktok"`
	Status    *string `query:"status" validate:"omitempty,oneof=pending approved rejected published cancelled"`
	StartDate *string `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate   *string `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
	Page      int     `query:"page" validate:"omitempty,min=1"`
	Limit     int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sort_by" validate:"omitempty,oneof=post_date created_at"`
	SortOrder *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}

// ListRecurrentPdvsQuery representa os parâmetros de query para listar PDVs recorrentes
type ListRecurrentPdvsQuery struct {
	RepName   *string `query:"rep_name" validate:"omitempty,max=100"`
	City      *string `query:"city" validate:"omitempty,max=100"`
	Page      int     `query:"page" validate:"omitempty,min=1"`
	Limit     int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sort_by" validate:"omitempty,oneof=name created_at"`
	SortOrder *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
