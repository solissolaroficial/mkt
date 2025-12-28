package request

// CreateKpiRequest defines the structure for creating KPI requests
type CreateKpiRequest struct {
	Title string `json:"title" validate:"required"`
	Color string `json:"color" validate:"required"`
	Unit  string `json:"unit" validate:"oneof=currency percent number"`
}

// UpdateKpiRequest defines the structure for updating KPI requests
type UpdateKpiRequest struct {
	Title *string `json:"title,omitempty"`
	Color *string `json:"color,omitempty"`
	Unit  *string `json:"unit,omitempty" validate:"omitempty,oneof=currency percent number"`
}
