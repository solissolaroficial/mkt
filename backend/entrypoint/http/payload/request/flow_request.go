package request

// CreateFlowRequest representa a requisição para criar um fluxo
type CreateFlowRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
	SortOrder   int     `json:"sort_order"`
}

// UpdateFlowRequest representa a requisição para atualizar um fluxo
type UpdateFlowRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Color       *string `json:"color"`
	SortOrder   *int    `json:"sort_order"`
}

// ReorderFlowsRequest representa a requisição para reordenar fluxos
type ReorderFlowsRequest struct {
	FlowIDs []string `json:"flow_ids" validate:"required"`
}
