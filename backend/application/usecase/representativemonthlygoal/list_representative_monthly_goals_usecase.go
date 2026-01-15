package representativemonthlygoal

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// ListRepresentativeMonthlyGoalsInput defines input data for listing monthly goals
type ListRepresentativeMonthlyGoalsInput struct {
	RepresentativeID *uuid.UUID
	Month            *int
	Year             *int
	Page             int
	PageSize         int
	SortBy           string
	SortOrder        string
}

// ListRepresentativeMonthlyGoalsOutput defines output data for listing monthly goals
type ListRepresentativeMonthlyGoalsOutput struct {
	Data       []*RepresentativeMonthlyGoalData
	Total      int64
	Page       int
	PageSize   int
	TotalPages int
}

// RepresentativeMonthlyGoalData represents a monthly goal in the list
type RepresentativeMonthlyGoalData struct {
	ID                 uuid.UUID
	RepresentativeID   uuid.UUID
	RepresentativeName string
	Month              int
	Year               int
	Target             float64
	Realized           float64
	Percentage         float64
	Remaining          float64
	IsTargetMet        bool
	CreatedAt          string
	UpdatedAt          string
}

type ListRepresentativeMonthlyGoalsUseCase struct {
	monthlyGoalGateway    gateway.RepresentativeMonthlyGoalGateway
	representativeGateway gateway.RepresentativeGateway
}

func NewListRepresentativeMonthlyGoalsUseCase(
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway,
	representativeGateway gateway.RepresentativeGateway,
) *ListRepresentativeMonthlyGoalsUseCase {
	return &ListRepresentativeMonthlyGoalsUseCase{
		monthlyGoalGateway:    monthlyGoalGateway,
		representativeGateway: representativeGateway,
	}
}

func (uc *ListRepresentativeMonthlyGoalsUseCase) Execute(ctx context.Context, input ListRepresentativeMonthlyGoalsInput) (*ListRepresentativeMonthlyGoalsOutput, error) {
	criteria := domain.NewRepresentativeMonthlyGoalCriteria()

	if input.RepresentativeID != nil {
		criteria = criteria.WithRepresentativeID(*input.RepresentativeID)
	}
	if input.Month != nil {
		criteria = criteria.WithMonth(*input.Month)
	}
	if input.Year != nil {
		criteria = criteria.WithYear(*input.Year)
	}
	if input.Page > 0 {
		criteria = criteria.WithPagination(input.Page, input.PageSize)
	}

	goals, total, err := uc.monthlyGoalGateway.List(criteria)
	if err != nil {
		return nil, err
	}

	// Convert to output format with representative names
	data := make([]*RepresentativeMonthlyGoalData, 0, len(goals))
	for _, goal := range goals {
		repName := ""
		if rep, err := uc.representativeGateway.FindByID(goal.RepresentativeID()); err == nil {
			repName = rep.Name()
		}

		data = append(data, &RepresentativeMonthlyGoalData{
			ID:                 goal.ID(),
			RepresentativeID:   goal.RepresentativeID(),
			RepresentativeName: repName,
			Month:              goal.Month(),
			Year:               goal.Year(),
			Target:             goal.Target(),
			Realized:           goal.Realized(),
			Percentage:         goal.PercentageAchieved(),
			Remaining:          goal.Remaining(),
			IsTargetMet:        goal.IsTargetMet(),
			CreatedAt:          goal.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:          goal.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
		})
	}

	totalPages := int(total) / input.PageSize
	if int(total)%input.PageSize > 0 {
		totalPages++
	}

	return &ListRepresentativeMonthlyGoalsOutput{
		Data:       data,
		Total:      total,
		Page:       input.Page,
		PageSize:   input.PageSize,
		TotalPages: totalPages,
	}, nil
}
