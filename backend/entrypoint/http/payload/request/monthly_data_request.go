package request

// UpdateMonthlyDataRequest defines structure for updating monthly data requests
type UpdateMonthlyDataRequest struct {
	Realized  *float64    `json:"realized,omitempty"`
	Meta      *float64    `json:"meta,omitempty"`
	Breakdown interface{} `json:"breakdown,omitempty"`
	Context   string      `json:"context,omitempty"` // Context of the change
}
