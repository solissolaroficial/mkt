package gateway

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type representativeStatsGatewayImpl struct {
	db *gorm.DB
}

// NewRepresentativeStatsGateway creates a new RepresentativeStatsGateway implementation
func NewRepresentativeStatsGateway(db *gorm.DB) gateway.RepresentativeStatsGateway {
	return &representativeStatsGatewayImpl{
		db: db,
	}
}

// GetOnlineActionCount returns the count of online marketing actions for a representative
func (g *representativeStatsGatewayImpl) GetOnlineActionCount(ctx interface{}, representativeUUID string) (int64, error) {
	var count int64
	err := g.db.Model(&model.RepMarketingActionModel{}).
		Where("representative_uuid = ?", representativeUUID).
		Count(&count).Error
	return count, err
}

// GetOfflineActionCount returns the count of offline actions for a representative
func (g *representativeStatsGatewayImpl) GetOfflineActionCount(ctx interface{}, representativeUUID string) (int64, error) {
	var count int64
	err := g.db.Model(&model.OfflineActionModel{}).
		Where("representative_uuid = ?", representativeUUID).
		Count(&count).Error
	return count, err
}

// GetOfflineActionValue returns the total value of offline actions for a representative
func (g *representativeStatsGatewayImpl) GetOfflineActionValue(ctx interface{}, representativeUUID string) (int64, error) {
	var total int64
	err := g.db.Model(&model.OfflineActionModel{}).
		Where("representative_uuid = ?", representativeUUID).
		Select("COALESCE(SUM(value), 0)").
		Scan(&total).Error
	return total, err
}

// GetShowroomItemCount returns the count of showroom items for a representative
func (g *representativeStatsGatewayImpl) GetShowroomItemCount(ctx interface{}, representativeUUID string) (int64, error) {
	var count int64
	err := g.db.Model(&model.ShowroomItemModel{}).
		Where("representative_uuid = ?", representativeUUID).
		Count(&count).Error
	return count, err
}

// GetRepMarketingCount returns the count of representative marketing actions for a representative
func (g *representativeStatsGatewayImpl) GetRepMarketingCount(ctx interface{}, representativeUUID string) (int64, error) {
	var count int64
	err := g.db.Model(&model.RepMarketingActionModel{}).
		Where("representative_uuid = ?", representativeUUID).
		Count(&count).Error
	return count, err
}

// GetRepresentativeStats returns all statistics for a representative
func (g *representativeStatsGatewayImpl) GetRepresentativeStats(ctx interface{}, representativeUUID string) (*entity.RepresentativeStats, error) {
	type StatsResult struct {
		UUID               string
		OnlineActionCount  int64
		OfflineActionCount int64
		OfflineActionValue int64
		ShowroomItemCount  int64
		RepMarketingCount  int64
		TrainingCount      int64
	}

	var result StatsResult

	// Use a single aggregated query for better performance
	err := g.db.Raw(`
		SELECT
			? as uuid,
			COALESCE(online_actions.count, 0) as online_action_count,
			COALESCE(offline_actions.count, 0) as offline_action_count,
			COALESCE(offline_actions.total_value, 0) as offline_action_value,
			COALESCE(showroom_items.count, 0) as showroom_item_count,
			COALESCE(rep_marketing_actions.count, 0) as rep_marketing_count,
			0 as training_count
		FROM (SELECT ? as uuid) as representatives
		LEFT JOIN (
			SELECT representative_uuid, COUNT(*) as count
			FROM rep_marketing_actions
			WHERE representative_uuid = ?
			GROUP BY representative_uuid
		) online_actions ON representatives.uuid = online_actions.representative_uuid
		LEFT JOIN (
			SELECT representative_uuid, COUNT(*) as count, SUM(COALESCE(value, 0)) as total_value
			FROM offline_actions
			WHERE representative_uuid = ?
			GROUP BY representative_uuid
		) offline_actions ON representatives.uuid = offline_actions.representative_uuid
		LEFT JOIN (
			SELECT representative_uuid, COUNT(*) as count
			FROM showroom_items
			WHERE representative_uuid = ?
			GROUP BY representative_uuid
		) showroom_items ON representatives.uuid = showroom_items.representative_uuid
		LEFT JOIN (
			SELECT representative_uuid, COUNT(*) as count
			FROM rep_marketing_actions
			WHERE representative_uuid = ?
			GROUP BY representative_uuid
		) rep_marketing_actions ON representatives.uuid = rep_marketing_actions.representative_uuid
	`, representativeUUID, representativeUUID, representativeUUID, representativeUUID).Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return entity.NewRepresentativeStats(
		result.UUID,
		result.OnlineActionCount,
		result.OfflineActionCount,
		result.OfflineActionValue,
		result.ShowroomItemCount,
		result.RepMarketingCount,
		result.TrainingCount,
	), nil
}

// GetBatchRepresentativeStats returns statistics for multiple representatives
func (g *representativeStatsGatewayImpl) GetBatchRepresentativeStats(ctx interface{}, representativeUUIDs []string) ([]*entity.RepresentativeStats, error) {
	type StatsResult struct {
		UUID               string
		OnlineActionCount  int64
		OfflineActionCount int64
		OfflineActionValue int64
		ShowroomItemCount  int64
		RepMarketingCount  int64
		TrainingCount      int64
	}

	var results []StatsResult

	// Use a single aggregated query for better performance
	err := g.db.Raw(`
		SELECT
			r.uuid as uuid,
			COALESCE(online_actions.count, 0) as online_action_count,
			COALESCE(offline_actions.count, 0) as offline_action_count,
			COALESCE(offline_actions.total_value, 0) as offline_action_value,
			COALESCE(showroom_items.count, 0) as showroom_item_count,
			COALESCE(rep_marketing_actions.count, 0) as rep_marketing_count,
			0 as training_count
		FROM representatives r
		LEFT JOIN (
			SELECT representative_uuid, COUNT(*) as count
			FROM rep_marketing_actions
			WHERE representative_uuid IN ?
			GROUP BY representative_uuid
		) online_actions ON r.uuid = online_actions.representative_uuid
		LEFT JOIN (
			SELECT representative_uuid, COUNT(*) as count, SUM(COALESCE(value, 0)) as total_value
			FROM offline_actions
			WHERE representative_uuid IN ?
			GROUP BY representative_uuid
		) offline_actions ON r.uuid = offline_actions.representative_uuid
		LEFT JOIN (
			SELECT representative_uuid, COUNT(*) as count
			FROM showroom_items
			WHERE representative_uuid IN ?
			GROUP BY representative_uuid
		) showroom_items ON r.uuid = showroom_items.representative_uuid
		LEFT JOIN (
			SELECT representative_uuid, COUNT(*) as count
			FROM rep_marketing_actions
			WHERE representative_uuid IN ?
			GROUP BY representative_uuid
		) rep_marketing_actions ON r.uuid = rep_marketing_actions.representative_uuid
		WHERE r.uuid IN ?
	`, representativeUUIDs, representativeUUIDs, representativeUUIDs, representativeUUIDs, representativeUUIDs).Scan(&results).Error

	if err != nil {
		return nil, err
	}

	stats := make([]*entity.RepresentativeStats, len(results))
	for i, result := range results {
		stats[i] = entity.NewRepresentativeStats(
			result.UUID,
			result.OnlineActionCount,
			result.OfflineActionCount,
			result.OfflineActionValue,
			result.ShowroomItemCount,
			result.RepMarketingCount,
			result.TrainingCount,
		)
	}

	return stats, nil
}
