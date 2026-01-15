package representatives

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetRepresentativeStatsInput defines input data for getting representative statistics
type GetRepresentativeStatsInput struct {
	ID uuid.UUID
}

// GetRepresentativeStatsOutput defines output data for representative statistics
type GetRepresentativeStatsOutput struct {
	UUID               string
	OnlineCount        int64 // Alias para compatibilidade com implementação antiga
	OnlineActionCount  int64
	OfflineCount       int64 // Alias para compatibilidade com implementação antiga
	OfflineActionCount int64
	OfflineValue       int64 // Alias para compatibilidade com implementação antiga
	OfflineActionValue int64
	ShowroomItemCount  int64
	RepMarketingCount  int64
	TotalActions       int64
}

type GetRepresentativeStatsUseCase struct {
	representativeGateway      gateway.RepresentativeGateway
	representativeStatsGateway gateway.RepresentativeStatsGateway
}

func NewGetRepresentativeStatsUseCase(
	representativeGateway gateway.RepresentativeGateway,
	representativeStatsGateway gateway.RepresentativeStatsGateway,
) *GetRepresentativeStatsUseCase {
	return &GetRepresentativeStatsUseCase{
		representativeGateway:      representativeGateway,
		representativeStatsGateway: representativeStatsGateway,
	}
}

func (uc *GetRepresentativeStatsUseCase) Execute(ctx context.Context, input GetRepresentativeStatsInput) (*GetRepresentativeStatsOutput, error) {
	// Check if representative exists
	exists, err := uc.representativeGateway.ExistsByID(input.ID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.ErrRepresentativeNotFound
	}

	// Get statistics
	stats, err := uc.representativeStatsGateway.GetRepresentativeStats(ctx, input.ID.String())
	if err != nil {
		return nil, err
	}

	return &GetRepresentativeStatsOutput{
		UUID:               stats.UUID(),
		OnlineCount:        stats.OnlineActionCount(),
		OnlineActionCount:  stats.OnlineActionCount(),
		OfflineCount:       stats.OfflineActionCount(),
		OfflineActionCount: stats.OfflineActionCount(),
		OfflineValue:       stats.OfflineActionValue(),
		OfflineActionValue: stats.OfflineActionValue(),
		ShowroomItemCount:  stats.ShowroomItemCount(),
		RepMarketingCount:  stats.RepMarketingCount(),
		TotalActions:       stats.TotalActions(),
	}, nil
}
