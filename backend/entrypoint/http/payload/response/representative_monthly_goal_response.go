package response

import (
	"github.com/google/uuid"
)

// CreateRepresentativeMonthlyGoalResponse represents a response after creating a monthly goal
type CreateRepresentativeMonthlyGoalResponse struct {
	ID               uuid.UUID `json:"id"`
	RepresentativeID uuid.UUID `json:"representativeId"`
	Month            int       `json:"month"`
	Year             int       `json:"year"`
	Target           float64   `json:"target"`
	Realized         float64   `json:"realized"`
	Percentage       float64   `json:"percentage"`
	CreatedAt        string    `json:"createdAt"`
	UpdatedAt        string    `json:"updatedAt"`
}

// GetRepresentativeMonthlyGoalResponse represents a response for getting a monthly goal
type GetRepresentativeMonthlyGoalResponse struct {
	ID               uuid.UUID `json:"id"`
	RepresentativeID uuid.UUID `json:"representativeId"`
	Month            int       `json:"month"`
	Year             int       `json:"year"`
	Target           float64   `json:"target"`
	Realized         float64   `json:"realized"`
	Percentage       float64   `json:"percentage"`
	Remaining        float64   `json:"remaining"`
	IsTargetMet      bool      `json:"isTargetMet"`
	CreatedAt        string    `json:"createdAt"`
	UpdatedAt        string    `json:"updatedAt"`
}

// UpdateRepresentativeMonthlyGoalResponse represents a response after updating a monthly goal
type UpdateRepresentativeMonthlyGoalResponse struct {
	ID               uuid.UUID `json:"id"`
	RepresentativeID uuid.UUID `json:"representativeId"`
	Month            int       `json:"month"`
	Year             int       `json:"year"`
	Target           float64   `json:"target"`
	Realized         float64   `json:"realized"`
	Percentage       float64   `json:"percentage"`
	UpdatedAt        string    `json:"updatedAt"`
}

// ListRepresentativeMonthlyGoalsResponse represents a response for listing monthly goals
type ListRepresentativeMonthlyGoalsResponse struct {
	Data       []RepresentativeMonthlyGoalData `json:"data"`
	Total      int64                           `json:"total"`
	Page       int                             `json:"page"`
	PageSize   int                             `json:"pageSize"`
	TotalPages int                             `json:"totalPages"`
}

// RepresentativeMonthlyGoalData represents a monthly goal in list
type RepresentativeMonthlyGoalData struct {
	ID               uuid.UUID `json:"id"`
	RepresentativeID uuid.UUID `json:"representativeId"`
	Month            int       `json:"month"`
	Year             int       `json:"year"`
	Target           float64   `json:"target"`
	Realized         float64   `json:"realized"`
	Percentage       float64   `json:"percentage"`
	Remaining        float64   `json:"remaining"`
	IsTargetMet      bool      `json:"isTargetMet"`
	CreatedAt        string    `json:"createdAt"`
	UpdatedAt        string    `json:"updatedAt"`
}

// GetRepresentativeGoalsTableDataResponse represents a response for getting table data (transposed view)
type GetRepresentativeGoalsTableDataResponse struct {
	Year    int                     `json:"year"`
	Months  []MonthData             `json:"months"`
	Rows    []RepresentativeRowData `json:"rows"`
	Summary TableSummaryData        `json:"summary"`
}

// MonthData represents a month column in table
type MonthData struct {
	Month       int     `json:"month"`
	MonthName   string  `json:"monthName"`
	TargetSum   float64 `json:"targetSum"`
	RealizedSum float64 `json:"realizedSum"`
}

// RepresentativeRowData represents a row in transposed table (one representative)
type RepresentativeRowData struct {
	RepresentativeID   string       `json:"representativeId"`
	RepresentativeName string       `json:"representativeName"`
	Company            string       `json:"company"`
	Region             string       `json:"region"`
	City               string       `json:"city"`
	MonthValues        []MonthValue `json:"monthValues"`
	TotalTarget        float64      `json:"totalTarget"`
	TotalRealized      float64      `json:"totalRealized"`
	TotalPercentage    float64      `json:"totalPercentage"`
}

// MonthValue represents a value for a specific month
type MonthValue struct {
	Month      int     `json:"month"`
	Target     float64 `json:"target"`
	Realized   float64 `json:"realized"`
	Percentage float64 `json:"percentage"`
	IsMet      bool    `json:"isMet"`
}

// TableSummaryData represents summary statistics for table
type TableSummaryData struct {
	TotalRepresentatives int     `json:"totalRepresentatives"`
	TotalTarget          float64 `json:"totalTarget"`
	TotalRealized        float64 `json:"totalRealized"`
	OverallPercentage    float64 `json:"overallPercentage"`
	GoalsMetCount        int     `json:"goalsMetCount"`
	GoalsNotMetCount     int     `json:"goalsNotMetCount"`
}
