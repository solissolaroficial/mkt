package representatives

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetRepresentativeProfileInput defines input data for getting a representative profile by name
type GetRepresentativeProfileInput struct {
	Name string
}

// GetRepresentativeProfileOutput defines output data for representative profile
type GetRepresentativeProfileOutput struct {
	UUID              string
	Code              int
	Name              string
	Email             string
	Phone             string
	Company           string
	Region            string
	City              string
	Attendant         string
	Active            bool
	CreatedAt         string
	UpdatedAt         string
	TrainingCount     int64
	OnlineCount       int64
	OfflineCount      int64
	OfflineValue      int64
	ShowroomItemCount int64
	RepMarketingCount int64
	TotalActions      int64
}

type GetRepresentativeProfileUseCase struct {
	representativeGateway      gateway.RepresentativeGateway
	representativeStatsGateway gateway.RepresentativeStatsGateway
}

func NewGetRepresentativeProfileUseCase(
	representativeGateway gateway.RepresentativeGateway,
	representativeStatsGateway gateway.RepresentativeStatsGateway,
) *GetRepresentativeProfileUseCase {
	return &GetRepresentativeProfileUseCase{
		representativeGateway:      representativeGateway,
		representativeStatsGateway: representativeStatsGateway,
	}
}

func (uc *GetRepresentativeProfileUseCase) Execute(ctx context.Context, input GetRepresentativeProfileInput) (*GetRepresentativeProfileOutput, error) {
	// Find representative by name using criteria
	criteria := domain.NewRepresentativeCriteria()
	criteria = criteria.WithName(input.Name)

	representatives, _, err := uc.representativeGateway.FindByCriteria(criteria, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(representatives) == 0 {
		return nil, errors.ErrRepresentativeNotFound
	}

	representative := representatives[0]

	// Get representative stats
	stats, err := uc.representativeStatsGateway.GetRepresentativeStats(ctx, representative.UUID().String())
	if err != nil {
		return nil, err
	}

	return &GetRepresentativeProfileOutput{
		UUID:              representative.UUID().String(),
		Code:              representative.Code().Value(),
		Name:              representative.Name(),
		Email:             representative.Email(),
		Phone:             representative.Phone(),
		Company:           representative.Company(),
		Region:            representative.Region(),
		City:              representative.City(),
		Attendant:         representative.Attendant(),
		Active:            representative.Active(),
		CreatedAt:         representative.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:         representative.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
		TrainingCount:     stats.TrainingCount(),
		OnlineCount:       stats.OnlineActionCount(),
		OfflineCount:      stats.OfflineActionCount(),
		OfflineValue:      stats.OfflineActionValue(),
		ShowroomItemCount: stats.ShowroomItemCount(),
		RepMarketingCount: stats.RepMarketingCount(),
		TotalActions:      stats.TotalActions(),
	}, nil
}
