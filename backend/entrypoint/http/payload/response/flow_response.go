package response

// FlowResponse representa a resposta de um fluxo
type FlowResponse struct {
	UUID        string  `json:"uuid"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
	SortOrder   int     `json:"sort_order"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// FlowListResponse representa a resposta de lista de fluxos
type FlowListResponse struct {
	Data       []FlowResponse `json:"data"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}

// ToFlowResponse converte uma entity.Flow para FlowResponse
func ToFlowResponse(flow interface{}) FlowResponse {
	// This will be implemented in the controller
	return FlowResponse{}
}
