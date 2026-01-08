package response

// MetaResponse define a estrutura de metadados para respostas paginadas
type MetaResponse struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
}
