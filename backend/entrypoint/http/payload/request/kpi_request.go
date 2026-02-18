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

// GetKpisBySlugsRequest defines the structure for getting KPIs by slugs
type GetKpisBySlugsRequest struct {
	Slugs []string `json:"slugs" validate:"required,min=1"`
}

// UpdateMonthlyDataRequest defines the structure for updating monthly data
type UpdateMonthlyDataRequest struct {
	Year      int         `json:"year" validate:"required"`
	Month     string      `json:"month" validate:"required"`
	Realized  *float64    `json:"realized,omitempty"`
	Meta      *float64    `json:"meta,omitempty"`
	Breakdown interface{} `json:"breakdown,omitempty"`
	Context   string      `json:"context,omitempty"`
}

// AddDailyEntryRequest defines the structure for adding a daily entry
type AddDailyEntryRequest struct {
	Year    int     `json:"year" validate:"required"`
	Month   string  `json:"month" validate:"required"`
	Date    string  `json:"date" validate:"required"`
	Value   float64 `json:"value" validate:"required"`
	Context string  `json:"context,omitempty"`
}

// UpdateDailyEntryRequest defines the structure for updating a daily entry
type UpdateDailyEntryRequest struct {
	Year    int     `json:"year" validate:"required"`
	Month   string  `json:"month" validate:"required"`
	Date    string  `json:"date" validate:"required"`
	Value   float64 `json:"value" validate:"required"`
	Context string  `json:"context,omitempty"`
}

// DeleteDailyEntryRequest defines the structure for deleting a daily entry
type DeleteDailyEntryRequest struct {
	Year    int    `json:"year" validate:"required"`
	Month   string `json:"month" validate:"required"`
	Date    string `json:"date" validate:"required"`
	Context string `json:"context,omitempty"`
}
