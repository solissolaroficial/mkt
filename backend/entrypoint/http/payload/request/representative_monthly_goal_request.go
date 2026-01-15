package request

// CreateRepresentativeMonthlyGoalRequest represents a request to create a monthly goal
type CreateRepresentativeMonthlyGoalRequest struct {
	RepresentativeID string  `json:"representativeId" validate:"required,uuid"`
	Month            int     `json:"month" validate:"required,min=1,max=12"`
	Year             int     `json:"year" validate:"required,min=2000,max=2100"`
	Target           float64 `json:"target" validate:"required,min=0"`
}

// UpdateRepresentativeMonthlyGoalRequest represents a request to update a monthly goal
type UpdateRepresentativeMonthlyGoalRequest struct {
	Target   *float64 `json:"target" validate:"omitempty,min=0"`
	Realized *float64 `json:"realized" validate:"omitempty,min=0"`
}

// GetRepresentativeMonthlyGoalRequest represents a request to get a monthly goal
type GetRepresentativeMonthlyGoalRequest struct {
	ID string `validate:"required,uuid"`
}

// ListRepresentativeMonthlyGoalsRequest represents a request to list monthly goals
type ListRepresentativeMonthlyGoalsRequest struct {
	Page             int     `query:"page" validate:"omitempty,min=1"`
	PageSize         int     `query:"pageSize" validate:"omitempty,min=1,max=100"`
	RepresentativeID *string `query:"representativeId" validate:"omitempty,uuid"`
	Month            *int    `query:"month" validate:"omitempty,min=1,max=12"`
	Year             *int    `query:"year" validate:"omitempty,min=2000,max=2100"`
	SortBy           string  `query:"sortBy" validate:"omitempty,oneof=month year target realized createdAt updatedAt"`
	SortOrder        string  `query:"sortOrder" validate:"omitempty,oneof=asc desc ASC DESC"`
}

// GetRepresentativeGoalsTableDataRequest represents a request to get table data
type GetRepresentativeGoalsTableDataRequest struct {
	Year  int  `query:"year" validate:"required,min=2000,max=2100"`
	Month *int `query:"month" validate:"omitempty,min=1,max=12"`
}

// DeleteRepresentativeMonthlyGoalRequest represents a request to delete a monthly goal
type DeleteRepresentativeMonthlyGoalRequest struct {
	ID string `validate:"required,uuid"`
}
