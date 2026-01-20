package gateway

import (
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	commenterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type commentGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.CommentMapper
}

// NewCommentGateway cria uma nova implementação de CommentGateway
func NewCommentGateway(db *gorm.DB) gateway.CommentGateway {
	return &commentGatewayImpl{
		db:     db,
		mapper: &mapper.CommentMapper{},
	}
}

// Create cria um novo comentário
func (g *commentGatewayImpl) Create(comment *entity.Comment) error {
	commentModel := g.mapper.ToModel(comment)
	if commentModel == nil {
		return errors.New("failed to convert comment to model")
	}

	if err := g.db.Create(commentModel).Error; err != nil {
		return err
	}

	return nil
}

// Update atualiza um comentário existente
func (g *commentGatewayImpl) Update(comment *entity.Comment) error {
	commentModel := g.mapper.ToModel(comment)
	if commentModel == nil {
		return errors.New("failed to convert comment to model")
	}

	result := g.db.Where("uuid = ?", commentModel.UUID).Save(commentModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &commenterrors.CommentNotFoundError{ID: comment.ID().String()}
	}

	return nil
}

// Delete deleta um comentário
func (g *commentGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.Comment{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &commenterrors.CommentNotFoundError{ID: id.String()}
	}

	return nil
}

// FindByID busca um comentário pelo ID
func (g *commentGatewayImpl) FindByID(id uuid.UUID) (*entity.Comment, error) {
	var commentModel model.Comment

	err := g.db.
		Preload("Task").
		Preload("User").
		Where("uuid = ?", id).
		First(&commentModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &commenterrors.CommentNotFoundError{ID: id.String()}
		}
		return nil, err
	}

	return g.mapper.ToEntity(&commentModel), nil
}

// FindAll busca todos os comentários com paginação
func (g *commentGatewayImpl) FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Comment, int64, error) {
	var commentModels []model.Comment
	var total int64

	// Get total count
	if err := g.db.Model(&model.Comment{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Comment{}).
		Preload("Task").
		Preload("User")

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&commentModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	commentModelPointers := make([]*model.Comment, len(commentModels))
	for i := range commentModels {
		commentModelPointers[i] = &commentModels[i]
	}

	// Convert slice of models to slice of entities
	return g.mapper.ToEntityList(commentModelPointers), total, nil
}

// FindByTaskID busca comentários por task ID
func (g *commentGatewayImpl) FindByTaskID(taskID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Comment, int64, error) {
	var commentModels []model.Comment
	var total int64

	// Get total count
	if err := g.db.Model(&model.Comment{}).Where("task_uuid = ?", taskID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Comment{}).
		Preload("Task").
		Preload("User").
		Where("task_uuid = ?", taskID)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Validar limit antes de aplicar
	limit := pagination.Limit()
	offset := pagination.Offset()

	if limit <= 0 {
		// Se limit for 0 ou negativo, não aplicar limit (retornar todos)
		err := query.Offset(offset).Find(&commentModels).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		// Aplicar paginação normal
		err := query.Offset(offset).Limit(limit).Find(&commentModels).Error
		if err != nil {
			return nil, 0, err
		}
	}

	// Convert slice of models to slice of pointers
	commentModelPointers := make([]*model.Comment, len(commentModels))
	for i := range commentModels {
		commentModelPointers[i] = &commentModels[i]
	}

	return g.mapper.ToEntityList(commentModelPointers), total, nil
}

// FindByUserID busca comentários por user ID
func (g *commentGatewayImpl) FindByUserID(userID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Comment, int64, error) {
	var commentModels []model.Comment
	var total int64

	// Get total count
	if err := g.db.Model(&model.Comment{}).Where("user_uuid = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Comment{}).
		Preload("Task").
		Preload("User").
		Where("user_uuid = ?", userID)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&commentModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	commentModelPointers := make([]*model.Comment, len(commentModels))
	for i := range commentModels {
		commentModelPointers[i] = &commentModels[i]
	}

	return g.mapper.ToEntityList(commentModelPointers), total, nil
}
