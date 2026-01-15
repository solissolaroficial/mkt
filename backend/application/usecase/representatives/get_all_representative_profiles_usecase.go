package representatives

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// GetAllRepresentativeProfilesOutput defines output data for all representative profiles
type GetAllRepresentativeProfilesOutput struct {
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

type GetAllRepresentativeProfilesUseCase struct {
	representativeGateway      gateway.RepresentativeGateway
	representativeStatsGateway gateway.RepresentativeStatsGateway
}

func NewGetAllRepresentativeProfilesUseCase(
	representativeGateway gateway.RepresentativeGateway,
	representativeStatsGateway gateway.RepresentativeStatsGateway,
) *GetAllRepresentativeProfilesUseCase {
	return &GetAllRepresentativeProfilesUseCase{
		representativeGateway:      representativeGateway,
		representativeStatsGateway: representativeStatsGateway,
	}
}

func (uc *GetAllRepresentativeProfilesUseCase) Execute(ctx context.Context) ([]*GetAllRepresentativeProfilesOutput, error) {
	// Find all representatives
	pagination := valueobject.NewPagination(1, 1000) // Get all representatives
	representatives, _, err := uc.representativeGateway.FindAll(&pagination, nil)
	if err != nil {
		return nil, err
	}

	// Get stats for all representatives
	repUUIDs := make([]string, len(representatives))
	for i, rep := range representatives {
		repUUIDs[i] = rep.UUID().String()
	}

	statsList, err := uc.representativeStatsGateway.GetBatchRepresentativeStats(ctx, repUUIDs)
	if err != nil {
		return nil, err
	}

	// Create stats map for easy lookup
	statsMap := make(map[string]*entity.RepresentativeStats)
	for _, stats := range statsList {
		statsMap[stats.UUID()] = stats
	}

	// Build output
	output := make([]*GetAllRepresentativeProfilesOutput, 0, len(representatives))
	for _, representative := range representatives {
		stats, ok := statsMap[representative.UUID().String()]
		if !ok {
			// If no stats found, use default values
			stats = entity.NewRepresentativeStats(
				representative.UUID().String(),
				0, 0, 0, 0, 0, 0,
			)
		}

		output = append(output, &GetAllRepresentativeProfilesOutput{
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
		})
	}

	return output, nil
}
