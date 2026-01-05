package gateway

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	taskerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type taskGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.TaskMapper
}

// NewTaskGateway cria uma nova implementação de TaskGateway
func NewTaskGateway(db *gorm.DB) gateway.TaskGateway {
	return &taskGatewayImpl{
		db:     db,
		mapper: &mapper.TaskMapper{},
	}
}

// Create cria uma nova tarefa
func (g *taskGatewayImpl) Create(task *entity.Task) error {
	taskModel, err := g.mapper.ToModel(task)
	if err != nil {
		return err
	}

	if err := g.db.Create(taskModel).Error; err != nil {
		return err
	}

	return nil
}

// Update atualiza uma tarefa existente
func (g *taskGatewayImpl) Update(task *entity.Task) error {
	taskModel, err := g.mapper.ToModel(task)
	if err != nil {
		return err
	}

	result := g.db.Where("uuid = ?", taskModel.UUID).Save(taskModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &taskerrors.TaskNotFoundError{ID: task.ID().String()}
	}

	return nil
}

// Delete deleta uma tarefa (soft delete)
func (g *taskGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.Task{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &taskerrors.TaskNotFoundError{ID: id.String()}
	}

	return nil
}

// FindByID busca uma tarefa pelo ID
func (g *taskGatewayImpl) FindByID(id uuid.UUID) (*entity.Task, error) {
	var taskModel model.Task

	err := g.db.
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee").
		Where("uuid = ?", id).
		First(&taskModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &taskerrors.TaskNotFoundError{ID: id.String()}
		}
		return nil, err
	}

	return g.mapper.ToEntity(&taskModel)
}

// FindAll busca todas as tarefas com paginação
func (g *taskGatewayImpl) FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error) {
	var taskModels []model.Task
	var total int64

	// Get total count
	if err := g.db.Model(&model.Task{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Task{}).
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee")

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&taskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	taskModelPointers := make([]*model.Task, len(taskModels))
	for i := range taskModels {
		taskModelPointers[i] = &taskModels[i]
	}

	entities, err := g.mapper.ToEntityList(taskModelPointers)
	return entities, total, err
}

// FindByCriteria busca tarefas usando o padrão Criteria
func (g *taskGatewayImpl) FindByCriteria(criteria *gateway.TaskCriteria, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error) {
	var taskModels []model.Task
	var total int64

	// Build query
	query := g.db.Model(&model.Task{}).
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee")

	// Apply filters
	if len(criteria.Statuses) > 0 {
		query = query.Where("status IN ?", criteria.Statuses)
	}

	if len(criteria.Categories) > 0 {
		query = query.Where("category IN ?", criteria.Categories)
	}

	if len(criteria.Priorities) > 0 {
		query = query.Where("priority IN ?", criteria.Priorities)
	}

	if len(criteria.Flows) > 0 {
		query = query.Where("flow IN ?", criteria.Flows)
	}

	if len(criteria.AssigneeIDs) > 0 {
		query = query.Where("assignee_id IN ?", criteria.AssigneeIDs)
	}

	if criteria.Archived != nil {
		query = query.Where("archived = ?", *criteria.Archived)
	}

	if criteria.DueDateFrom != nil {
		query = query.Where("due_date >= ?", *criteria.DueDateFrom)
	}

	if criteria.DueDateTo != nil {
		query = query.Where("due_date <= ?", *criteria.DueDateTo)
	}

	if criteria.Title != nil {
		query = query.Where("title ILIKE ?", "%"+*criteria.Title+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&taskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	taskModelPointers := make([]*model.Task, len(taskModels))
	for i := range taskModels {
		taskModelPointers[i] = &taskModels[i]
	}

	entities, err := g.mapper.ToEntityList(taskModelPointers)
	return entities, total, err
}

// CountByCriteria conta tarefas usando o padrão Criteria
func (g *taskGatewayImpl) CountByCriteria(criteria *gateway.TaskCriteria) (int64, error) {
	var total int64

	// Build query
	query := g.db.Model(&model.Task{})

	// Apply filters
	if len(criteria.Statuses) > 0 {
		query = query.Where("status IN ?", criteria.Statuses)
	}

	if len(criteria.Categories) > 0 {
		query = query.Where("category IN ?", criteria.Categories)
	}

	if len(criteria.Priorities) > 0 {
		query = query.Where("priority IN ?", criteria.Priorities)
	}

	if len(criteria.Flows) > 0 {
		query = query.Where("flow IN ?", criteria.Flows)
	}

	if len(criteria.AssigneeIDs) > 0 {
		query = query.Where("assignee_id IN ?", criteria.AssigneeIDs)
	}

	if criteria.Archived != nil {
		query = query.Where("archived = ?", *criteria.Archived)
	}

	if criteria.DueDateFrom != nil {
		query = query.Where("due_date >= ?", *criteria.DueDateFrom)
	}

	if criteria.DueDateTo != nil {
		query = query.Where("due_date <= ?", *criteria.DueDateTo)
	}

	if criteria.Title != nil {
		query = query.Where("title ILIKE ?", "%"+*criteria.Title+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// FindByAssigneeID busca tarefas por assignee
func (g *taskGatewayImpl) FindByAssigneeID(assigneeID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error) {
	var taskModels []model.Task
	var total int64

	// Get total count
	if err := g.db.Model(&model.Task{}).Where("assignee_id = ?", assigneeID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Task{}).
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee").
		Where("assignee_id = ?", assigneeID)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&taskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	taskModelPointers := make([]*model.Task, len(taskModels))
	for i := range taskModels {
		taskModelPointers[i] = &taskModels[i]
	}

	entities, err := g.mapper.ToEntityList(taskModelPointers)
	return entities, total, err
}

// FindByStatus busca tarefas por status
func (g *taskGatewayImpl) FindByStatus(status string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error) {
	var taskModels []model.Task
	var total int64

	// Get total count
	if err := g.db.Model(&model.Task{}).Where("status = ?", status).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Task{}).
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee").
		Where("status = ?", status)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&taskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	taskModelPointers := make([]*model.Task, len(taskModels))
	for i := range taskModels {
		taskModelPointers[i] = &taskModels[i]
	}

	entities, err := g.mapper.ToEntityList(taskModelPointers)
	return entities, total, err
}

// FindByCategory busca tarefas por categoria
func (g *taskGatewayImpl) FindByCategory(category string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error) {
	var taskModels []model.Task
	var total int64

	// Get total count
	if err := g.db.Model(&model.Task{}).Where("category = ?", category).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Task{}).
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee").
		Where("category = ?", category)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&taskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	taskModelPointers := make([]*model.Task, len(taskModels))
	for i := range taskModels {
		taskModelPointers[i] = &taskModels[i]
	}

	entities, err := g.mapper.ToEntityList(taskModelPointers)
	return entities, total, err
}

// FindByFlow busca tarefas por fluxo
func (g *taskGatewayImpl) FindByFlow(flow string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error) {
	var taskModels []model.Task
	var total int64

	// Get total count
	if err := g.db.Model(&model.Task{}).Where("flow = ?", flow).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Task{}).
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee").
		Where("flow = ?", flow)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&taskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	taskModelPointers := make([]*model.Task, len(taskModels))
	for i := range taskModels {
		taskModelPointers[i] = &taskModels[i]
	}

	entities, err := g.mapper.ToEntityList(taskModelPointers)
	return entities, total, err
}

// FindArchived busca tarefas arquivadas
func (g *taskGatewayImpl) FindArchived(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error) {
	var taskModels []model.Task
	var total int64

	// Get total count
	if err := g.db.Model(&model.Task{}).Where("archived = ?", true).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Task{}).
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee").
		Where("archived = ?", true)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&taskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	taskModelPointers := make([]*model.Task, len(taskModels))
	for i := range taskModels {
		taskModelPointers[i] = &taskModels[i]
	}

	entities, err := g.mapper.ToEntityList(taskModelPointers)
	return entities, total, err
}

// FindOverdue busca tarefas atrasadas
func (g *taskGatewayImpl) FindOverdue(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error) {
	var taskModels []model.Task
	var total int64
	now := time.Now()

	// Get total count
	if err := g.db.Model(&model.Task{}).
		Where("due_date < ? AND status != ?", now, constants.TaskStatusCompleted).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Task{}).
		Preload("Subtasks").
		Preload("Comments").
		Preload("Assignee").
		Where("due_date < ? AND status != ?", now, constants.TaskStatusCompleted)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&taskModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	taskModelPointers := make([]*model.Task, len(taskModels))
	for i := range taskModels {
		taskModelPointers[i] = &taskModels[i]
	}

	entities, err := g.mapper.ToEntityList(taskModelPointers)
	return entities, total, err
}

// Reorder atualiza a ordem de múltiplas tarefas
func (g *taskGatewayImpl) Reorder(taskIDs []uuid.UUID) error {
	if len(taskIDs) == 0 {
		return nil
	}

	// Atualiza o sort_order de cada tarefa
	for i, taskID := range taskIDs {
		result := g.db.Model(&model.Task{}).
			Where("uuid = ?", taskID).
			Update("sort_order", i)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
