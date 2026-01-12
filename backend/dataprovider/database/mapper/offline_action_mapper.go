package mapper

import (
	"time"

	"gorm.io/gorm"

	"solis/backend/core/domain/entity"
	"solis/backend/core/domain/valueobject"
	"solis/backend/dataprovider/database/model"
)

type OfflineActionMapper struct{}

func NewOfflineActionMapper() *OfflineActionMapper {
	return &OfflineActionMapper{}
}

// ModelToEntity converte Model para Entity
func (m *OfflineActionMapper) ModelToEntity(model *model.OfflineActionModel) (*entity.OfflineAction, error) {
	// Converter action date
	actionDate := valueobject.ReconstructActionDate(model.ActionDate.Format("2006-01-02"))

	// Converter category para value object
	category, err := valueobject.NewOfflineCategory(model.Category)
	if err != nil {
		return nil, err
	}

	// Converter status
	status, err := valueobject.NewOfflineStatus(model.Status)
	if err != nil {
		return nil, err
	}

	// Converter scored
	scored, err := valueobject.NewScoredStatus(model.Scored)
	if err != nil {
		return nil, err
	}

	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return entity.ReconstructOfflineAction(
		model.UUID,
		model.RequestedAmount,
		actionDate,
		&category, // Passar category como pointer para value object
		model.Month,
		model.ApprovedAmount,
		model.OrderNumber,
		model.DepartureDate,
		model.DeliveryForecast,
		model.DeliveryDate,
		model.City,
		model.UF,
		&scored,
		&status,
		model.Observation,
		model.PDV,
		model.RepName,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *OfflineActionMapper) ModelsToEntities(models []*model.OfflineActionModel) ([]*entity.OfflineAction, error) {
	actions := make([]*entity.OfflineAction, len(models))
	for i, model := range models {
		action, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		actions[i] = action
	}
	return actions, nil
}

// EntityToModel converte Entity para Model
func (m *OfflineActionMapper) EntityToModel(action *entity.OfflineAction) *model.OfflineActionModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if action.DeletedAt() != nil {
		deletedAt.Time = *action.DeletedAt()
		deletedAt.Valid = true
	}

	return &model.OfflineActionModel{
		UUID:             action.ID(),
		RequestedAmount:  action.RequestedAmount().Value(),
		ActionDate:       action.ActionDate().Value(),
		Category:         action.Category().String(),
		Month:            action.Month(),
		ApprovedAmount:   action.ApprovedAmount(),
		OrderNumber:      action.OrderNumber(),
		DepartureDate:    action.DepartureDate(),
		DeliveryForecast: action.DeliveryForecast(),
		DeliveryDate:     action.DeliveryDate(),
		City:             action.City(),
		UF:               action.UF(),
		Scored:           action.Scored().String(),
		Status:           action.Status().String(),
		Observation:      action.Observation(),
		PDV:              action.PDV(),
		RepName:          action.RepName(),
		CreatedAt:        action.CreatedAt(),
		UpdatedAt:        action.UpdatedAt(),
		DeletedAt:        deletedAt,
	}
}
