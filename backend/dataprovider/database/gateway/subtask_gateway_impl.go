package gateway

import (
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	subtaskerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type subtaskGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.SubtaskMapper
}

// NewSubtaskGateway cria uma nova implementação de SubtaskGateway
func NewSubtaskGateway(db *gorm.DB) gateway.SubtaskGateway {
	return &subtaskGatewayImpl{
		db:     db,
		mapper: &mapper.SubtaskMapper{},
	}
}

// Create cria uma nova subtarefa
func (g *subtaskGatewayImpl) Create(subtask *entity.Subtask) error {
	subtaskModel := g.mapper.ToModel(subtask)

	if err := g.db.Create(subtaskModel).Error; err != nil {
		return err
	}

	return nil
}

// Update atualiza uma subtarefa existente
func (g *subtaskGatewayImpl) Update(subtask *entity.Subtask) error {
	subtaskModel := g.mapper.ToModel(subtask)

	result := g.db.Where("uuid = ?", subtaskModel.UUID).Save(subtaskModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &subtaskerrors.SubtaskNotFoundError{ID: subtask.ID().String()}
	}

	return nil
}

// Delete deleta uma subtarefa
func (g *subtaskGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.Subtask{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &subtaskerrors.SubtaskNotFoundError{ID: id.String()}
	}

	return nil
}

// FindByID busca uma subtarefa pelo ID
func (g *subtaskGatewayImpl) FindByID(id uuid.UUID) (*entity.Subtask, error) {
	var subtaskModel model.Subtask

	err := g.db.
		Preload("Task").
		Preload("Assignee").
		Where("uuid = ?", id).
		First(&subtaskModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &subtaskerrors.SubtaskNotFoundError{ID: id.String()}
		}
		return nil, err
	}

	return g.mapper.ToEntity(&subtaskModel), nil
}

// FindAll busca todas as subtarefas com paginação
func (g *subtaskGatewayImpl) FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error) {
	var subtaskModels []model.Subtask
	var total int64

	// Get total count
	if err := g.db.Model(&model.Subtask{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Subtask{}).
		Preload("Task").
		Preload("Assignee")

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&subtaskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	subtaskModelPointers := make([]*model.Subtask, len(subtaskModels))
	for i := range subtaskModels {
		subtaskModelPointers[i] = &subtaskModels[i]
	}

	entities := g.mapper.ToEntityList(subtaskModelPointers)
	return entities, total, nil
}

// FindByTaskID busca subtarefas por task ID
func (g *subtaskGatewayImpl) FindByTaskID(taskID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error) {
	var subtaskModels []model.Subtask
	var total int64

	// Get total count
	if err := g.db.Model(&model.Subtask{}).Where("task_uuid = ?", taskID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Subtask{}).
		Preload("Task").
		Preload("Assignee").
		Where("task_uuid = ?", taskID)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&subtaskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	subtaskModelPointers := make([]*model.Subtask, len(subtaskModels))
	for i := range subtaskModels {
		subtaskModelPointers[i] = &subtaskModels[i]
	}

	entities := g.mapper.ToEntityList(subtaskModelPointers)
	return entities, total, nil
}

// FindByAssigneeID busca subtarefas por assignee ID
func (g *subtaskGatewayImpl) FindByAssigneeID(assigneeID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error) {
	var subtaskModels []model.Subtask
	var total int64

	// Get total count
	if err := g.db.Model(&model.Subtask{}).Where("assignee_uuid = ?", assigneeID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Subtask{}).
		Preload("Task").
		Preload("Assignee").
		Where("assignee_uuid = ?", assigneeID)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&subtaskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	subtaskModelPointers := make([]*model.Subtask, len(subtaskModels))
	for i := range subtaskModels {
		subtaskModelPointers[i] = &subtaskModels[i]
	}

	entities := g.mapper.ToEntityList(subtaskModelPointers)
	return entities, total, nil
}

// FindCompleted busca subtarefas concluídas
func (g *subtaskGatewayImpl) FindCompleted(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error) {
	var subtaskModels []model.Subtask
	var total int64

	// Get total count
	if err := g.db.Model(&model.Subtask{}).Where("completed = ?", true).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Subtask{}).
		Preload("Task").
		Preload("Assignee").
		Where("completed = ?", true)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&subtaskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	subtaskModelPointers := make([]*model.Subtask, len(subtaskModels))
	for i := range subtaskModels {
		subtaskModelPointers[i] = &subtaskModels[i]
	}

	entities := g.mapper.ToEntityList(subtaskModelPointers)
	return entities, total, nil
}

// FindPending busca subtarefas pendentes
func (g *subtaskGatewayImpl) FindPending(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error) {
	var subtaskModels []model.Subtask
	var total int64

	// Get total count
	if err := g.db.Model(&model.Subtask{}).Where("completed = ?", false).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Subtask{}).
		Preload("Task").
		Preload("Assignee").
		Where("completed = ?", false)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&subtaskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	subtaskModelPointers := make([]*model.Subtask, len(subtaskModels))
	for i := range subtaskModels {
		subtaskModelPointers[i] = &subtaskModels[i]
	}

	entities := g.mapper.ToEntityList(subtaskModelPointers)
	return entities, total, nil
}
