package gateway

import (
	"context"
	"errors"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	calendarerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type calendarPostGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.CalendarPostMapper
}

// NewCalendarPostGateway cria uma nova implementação de CalendarPostGateway
func NewCalendarPostGateway(db *gorm.DB) gateway.CalendarPostGateway {
	return &calendarPostGatewayImpl{
		db:     db,
		mapper: mapper.NewCalendarPostMapper(),
	}
}

// Save salva um novo post no calendário
func (g *calendarPostGatewayImpl) Save(ctx context.Context, post *entity.CalendarPost) error {
	postModel, err := g.mapper.EntityToModel(post)
	if err != nil {
		return err
	}

	if err := g.db.WithContext(ctx).Create(postModel).Error; err != nil {
		return err
	}

	return nil
}

// FindByID busca um post pelo ID
func (g *calendarPostGatewayImpl) FindByID(ctx context.Context, id string) (*entity.CalendarPost, error) {
	var postModel model.CalendarPostModel

	err := g.db.WithContext(ctx).
		Preload("Assignee").
		Where("uuid = ?", id).
		First(&postModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, calendarerrors.ErrCalendarPostNotFound
		}
		return nil, err
	}

	return g.mapper.ModelToEntity(&postModel)
}

// Update atualiza um post existente
func (g *calendarPostGatewayImpl) Update(ctx context.Context, post *entity.CalendarPost) error {
	postModel, err := g.mapper.EntityToModel(post)
	if err != nil {
		return err
	}

	result := g.db.WithContext(ctx).Where("uuid = ?", postModel.UUID).Save(postModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return calendarerrors.ErrCalendarPostNotFound
	}

	return nil
}

// Delete deleta um post (soft delete)
func (g *calendarPostGatewayImpl) Delete(ctx context.Context, id string) error {
	result := g.db.WithContext(ctx).Delete(&model.CalendarPostModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return calendarerrors.ErrCalendarPostNotFound
	}

	return nil
}

// FindByCriteria busca posts usando o padrão Criteria.
// Aplica filtros dinâmicos baseados nos critérios fornecidos,
// suporta paginação e ordenação, e faz eager loading do Assignee.
func (g *calendarPostGatewayImpl) FindByCriteria(ctx context.Context, crit *domain.CalendarPostCriteria, pagination *valueobject.Pagination, sortOrder *valueobject.SortOrder) ([]*entity.CalendarPost, error) {
	var postModels []model.CalendarPostModel

	// Build query
	query := g.db.WithContext(ctx).Model(&model.CalendarPostModel{}).
		Preload("Assignee")

	// Apply filters
	if crit.Category() != nil {
		query = query.Where("category = ?", crit.Category().String())
	}

	if crit.Type() != nil {
		query = query.Where("type = ?", crit.Type().String())
	}

	if crit.Status() != nil {
		query = query.Where("status = ?", crit.Status().String())
	}

	if crit.AssigneeID() != nil {
		query = query.Where("assignee_id = ?", crit.AssigneeID())
	}

	if crit.StartDate() != nil {
		query = query.Where("date >= ?", *crit.StartDate())
	}

	if crit.EndDate() != nil {
		query = query.Where("date <= ?", *crit.EndDate())
	}

	if crit.Platform() != nil {
		// Usar operador JSONB com parâmetro para evitar injeção de SQL
		query = query.Where("? = ANY(platforms)", *crit.Platform())
	}

	// Apply sorting
	if sortOrder != nil {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&postModels).Error

	if err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	postModelPointers := make([]*model.CalendarPostModel, len(postModels))
	for i := range postModels {
		postModelPointers[i] = &postModels[i]
	}

	return g.mapper.ModelsToEntities(postModelPointers)
}

// CountByCriteria conta posts usando o padrão Criteria.
// Aplica os mesmos filtros de FindByCriteria mas retorna apenas o total de registros.
func (g *calendarPostGatewayImpl) CountByCriteria(ctx context.Context, crit *domain.CalendarPostCriteria) (int64, error) {
	var total int64

	// Build query
	query := g.db.WithContext(ctx).Model(&model.CalendarPostModel{})

	// Apply filters
	if crit.Category() != nil {
		query = query.Where("category = ?", crit.Category().String())
	}

	if crit.Type() != nil {
		query = query.Where("type = ?", crit.Type().String())
	}

	if crit.Status() != nil {
		query = query.Where("status = ?", crit.Status().String())
	}

	if crit.AssigneeID() != nil {
		query = query.Where("assignee_id = ?", crit.AssigneeID())
	}

	if crit.StartDate() != nil {
		query = query.Where("date >= ?", *crit.StartDate())
	}

	if crit.EndDate() != nil {
		query = query.Where("date <= ?", *crit.EndDate())
	}

	if crit.Platform() != nil {
		// Usar operador JSONB com parâmetro para evitar injeção de SQL
		query = query.Where("? = ANY(platforms)", *crit.Platform())
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// ExistsByID verifica se um post existe pelo ID
func (g *calendarPostGatewayImpl) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).Model(&model.CalendarPostModel{}).Where("uuid = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
