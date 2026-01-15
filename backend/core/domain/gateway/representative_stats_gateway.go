package gateway

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// RepresentativeStatsGateway defines the interface for representative statistics
type RepresentativeStatsGateway interface {
	// GetOnlineActionCount returns the count of online marketing actions for a representative
	GetOnlineActionCount(ctx interface{}, representativeUUID string) (int64, error)

	// GetOfflineActionCount returns the count of offline actions for a representative
	GetOfflineActionCount(ctx interface{}, representativeUUID string) (int64, error)

	// GetOfflineActionValue returns the total value of offline actions for a representative
	GetOfflineActionValue(ctx interface{}, representativeUUID string) (int64, error)

	// GetShowroomItemCount returns the count of showroom items for a representative
	GetShowroomItemCount(ctx interface{}, representativeUUID string) (int64, error)

	// GetRepMarketingCount returns the count of representative marketing actions for a representative
	GetRepMarketingCount(ctx interface{}, representativeUUID string) (int64, error)

	// GetRepresentativeStats returns all statistics for a representative
	GetRepresentativeStats(ctx interface{}, representativeUUID string) (*entity.RepresentativeStats, error)

	// GetBatchRepresentativeStats returns statistics for multiple representatives
	GetBatchRepresentativeStats(ctx interface{}, representativeUUIDs []string) ([]*entity.RepresentativeStats, error)
}
