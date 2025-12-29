package request

// UpdateMonthlyDataRequest defines structure for updating monthly data requests
type UpdateMonthlyDataRequest struct {
	Realized  *float64    `json:"realized,omitempty"`
	Meta      *float64    `json:"meta,omitempty"`
	Breakdown interface{} `json:"breakdown,omitempty"`
	Year      int         `json:"year,omitempty"`    // Year (e.g., 2024, 2025)
	Month     string      `json:"month,omitempty"`   // Month identifier (JAN, FEV, etc)
	Context   string      `json:"context,omitempty"` // Context of the change
}
