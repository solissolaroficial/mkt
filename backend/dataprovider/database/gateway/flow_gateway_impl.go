package gateway

import (
	"errors"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	flowerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type flowGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.FlowMapper
}

// NewFlowGateway cria uma nova implementação de FlowGateway
func NewFlowGateway(db *gorm.DB) gateway.FlowGateway {
	return &flowGatewayImpl{
		db:     db,
		mapper: &mapper.FlowMapper{},
	}
}

// Create cria um novo fluxo
func (g *flowGatewayImpl) Create(flow *entity.Flow) (*entity.Flow, error) {
	flowModel := g.mapper.ToModel(flow)

	if err := g.db.Create(flowModel).Error; err != nil {
		return nil, err
	}

	return g.mapper.ToEntity(flowModel), nil
}

// GetByID busca um fluxo por ID
func (g *flowGatewayImpl) GetByID(id string) (*entity.Flow, error) {
	var flowModel model.Flow

	err := g.db.Where("uuid = ?", id).First(&flowModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &flowerrors.FlowNotFoundError{}
		}
		return nil, err
	}

	return g.mapper.ToEntity(&flowModel), nil
}

// GetByUUID busca um fluxo por UUID
func (g *flowGatewayImpl) GetByUUID(uuid string) (*entity.Flow, error) {
	var flowModel model.Flow

	err := g.db.Where("uuid = ?", uuid).First(&flowModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &flowerrors.FlowNotFoundError{}
		}
		return nil, err
	}

	return g.mapper.ToEntity(&flowModel), nil
}

// GetAll lista todos os fluxos ativos ordenados por sort_order
func (g *flowGatewayImpl) GetAll() ([]*entity.Flow, error) {
	var flowModels []model.Flow

	err := g.db.
		Where("deleted_at IS NULL").
		Order("sort_order ASC").
		Find(&flowModels).Error

	if err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	flowModelPointers := make([]*model.Flow, len(flowModels))
	for i := range flowModels {
		flowModelPointers[i] = &flowModels[i]
	}

	return g.mapper.ToEntityList(flowModelPointers), nil
}

// Update atualiza um fluxo
func (g *flowGatewayImpl) Update(flow *entity.Flow) (*entity.Flow, error) {
	flowModel := g.mapper.ToModel(flow)

	result := g.db.Where("uuid = ?", flowModel.UUID).Save(flowModel)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &flowerrors.FlowNotFoundError{}
	}

	return g.mapper.ToEntity(flowModel), nil
}

// Delete remove um fluxo (soft delete)
func (g *flowGatewayImpl) Delete(id string) error {
	result := g.db.Delete(&model.Flow{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &flowerrors.FlowNotFoundError{}
	}

	return nil
}

// Reorder reordena os fluxos
func (g *flowGatewayImpl) Reorder(flowIDs []string) error {
	if len(flowIDs) == 0 {
		return nil
	}

	// Atualiza o sort_order de cada fluxo
	for i, flowID := range flowIDs {
		result := g.db.Model(&model.Flow{}).
			Where("uuid = ?", flowID).
			Update("sort_order", i)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
