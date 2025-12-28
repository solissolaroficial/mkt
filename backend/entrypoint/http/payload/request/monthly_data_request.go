package request

// UpdateMonthlyDataRequest defines the structure for updating monthly data requests
type UpdateMonthlyDataRequest struct {
	Realized  *float64    `json:"realized,omitempty"`
	Meta      *float64    `json:"meta,omitempty"`
	Breakdown interface{} `json:"breakdown,omitempty"`
}
