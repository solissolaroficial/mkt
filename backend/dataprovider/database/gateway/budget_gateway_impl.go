package gateway

import (
	"context"
	"errors"
	"sort"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	budgeterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type BudgetGatewayImpl struct {
	db *gorm.DB
}

// NewBudgetGateway cria uma nova instância do BudgetGateway
func NewBudgetGateway(db *gorm.DB) gateway.BudgetGateway {
	return &BudgetGatewayImpl{db: db}
}

// Create cria um novo BudgetItem no banco de dados
func (g *BudgetGatewayImpl) Create(ctx context.Context, budget *entity.BudgetItem) error {
	if budget == nil {
		return errors.New("budget cannot be nil")
	}

	budgetModel, err := mapper.ToBudgetItemModel(budget)
	if err != nil {
		return err
	}

	if err := g.db.WithContext(ctx).Create(budgetModel).Error; err != nil {
		return err
	}

	return nil
}

// FindByID busca um BudgetItem pelo ID
// Retorna ErrBudgetNotFound se não encontrado
func (g *BudgetGatewayImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.BudgetItem, error) {
	var model model.BudgetItemModel

	err := g.db.WithContext(ctx).Where("uuid = ?", id).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, budgeterrors.ErrBudgetNotFound
		}
		return nil, err
	}

	return mapper.ToBudgetItemDomain(&model)
}

// List busca BudgetItems baseados em critérios
// Aceita *BudgetCriteria como parâmetro
func (g *BudgetGatewayImpl) List(ctx context.Context, criteria interface{}) ([]*entity.BudgetItem, error) {
	query := g.buildQuery(ctx, criteria)

	var models []*model.BudgetItemModel
	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	return mapper.ToBudgetItemDomainList(models)
}

// Update atualiza um BudgetItem existente
// Retorna ErrBudgetNotFound se não encontrado
func (g *BudgetGatewayImpl) Update(ctx context.Context, budget *entity.BudgetItem) error {
	if budget == nil {
		return errors.New("budget cannot be nil")
	}

	budgetModel, err := mapper.ToBudgetItemModel(budget)
	if err != nil {
		return err
	}

	result := g.db.WithContext(ctx).Model(&model.BudgetItemModel{}).
		Where("uuid = ?", budget.ID()).
		Updates(map[string]interface{}{
			"object_uuid":   budgetModel.ObjectUUID,
			"group_uuid":    budgetModel.GroupUUID,
			"cod":           budgetModel.Cod,
			"desc":          budgetModel.Desc,
			"vals":          budgetModel.Vals,
			"realized_vals": budgetModel.RealizedVals,
			"year":          budgetModel.Year,
			"updated_at":    budgetModel.UpdatedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return budgeterrors.ErrBudgetNotFound
	}

	return nil
}

// Delete remove um BudgetItem (soft delete)
// Retorna ErrBudgetNotFound se não encontrado
func (g *BudgetGatewayImpl) Delete(ctx context.Context, id uuid.UUID) error {
	result := g.db.WithContext(ctx).Delete(&model.BudgetItemModel{}, "uuid = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return budgeterrors.ErrBudgetNotFound
	}

	return nil
}

// Count conta o número de BudgetItems baseados em critérios
func (g *BudgetGatewayImpl) Count(ctx context.Context, criteria interface{}) (int64, error) {
	query := g.buildQuery(ctx, criteria)

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistsByCode verifica se existe um item com o mesmo código
// Considera objectUUID, groupUUID, cod e year para unicidade
func (g *BudgetGatewayImpl) ExistsByCode(ctx context.Context, objectUUID *uuid.UUID, groupUUID *uuid.UUID, cod string, year int) (bool, error) {
	var count int64
	query := g.db.WithContext(ctx).Model(&model.BudgetItemModel{})

	// Construir query baseada nos parâmetros fornecidos
	if objectUUID != nil {
		query = query.Where("object_uuid = ?", objectUUID)
	}
	if groupUUID != nil {
		query = query.Where("group_uuid = ?", groupUUID)
	}
	if cod != "" {
		query = query.Where("cod = ?", cod)
	}
	if year > 0 {
		query = query.Where("year = ?", year)
	}

	err := query.Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// BatchCreate cria múltiplos BudgetItems em lote
// Usa transação para garantir atomicidade
func (g *BudgetGatewayImpl) BatchCreate(ctx context.Context, budgets []*entity.BudgetItem) error {
	if len(budgets) == 0 {
		return nil
	}

	models, err := mapper.ToBudgetItemModelList(budgets)
	if err != nil {
		return err
	}

	return g.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&models).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetDistinctYears retorna os anos disponíveis no banco
func (g *BudgetGatewayImpl) GetDistinctYears(ctx context.Context) ([]int, error) {
	var years []int

	err := g.db.WithContext(ctx).Model(&model.BudgetItemModel{}).
		Distinct("year").
		Order("year DESC").
		Pluck("year", &years).Error

	if err != nil {
		return nil, err
	}

	return years, nil
}

// GetSummary retorna resumo agregado por objeto/grupo
// Aceita *BudgetCriteria como parâmetro
func (g *BudgetGatewayImpl) GetSummary(ctx context.Context, criteria interface{}) ([]*gateway.BudgetSummary, error) {
	// Buscar todos os items filtrados
	items, err := g.List(ctx, criteria)
	if err != nil {
		return nil, err
	}

	// Agrupar em memória
	summaryMap := make(map[string]*gateway.BudgetSummary)

	for _, item := range items {
		key := item.ObjectName() + "|" + item.GroupName()

		if summary, exists := summaryMap[key]; exists {
			summary.TotalBudget += item.GetTotalBudget()
			summary.TotalRealized += item.GetTotalRealized()
		} else {
			summaryMap[key] = &gateway.BudgetSummary{
				ObjectUUID:    item.ObjectUUID(),
				ObjectName:    item.ObjectName(),
				GroupUUID:     item.GroupUUID(),
				GroupName:     item.GroupName(),
				TotalBudget:   item.GetTotalBudget(),
				TotalRealized: item.GetTotalRealized(),
			}
		}
	}

	// Converter map para slice
	summaries := make([]*gateway.BudgetSummary, 0, len(summaryMap))
	for _, summary := range summaryMap {
		summary.Variance = summary.TotalBudget - summary.TotalRealized
		summaries = append(summaries, summary)
	}

	// Ordenar
	sort.Slice(summaries, func(i, j int) bool {
		if summaries[i].ObjectName == summaries[j].ObjectName {
			return summaries[i].GroupName < summaries[j].GroupName
		}
		return summaries[i].ObjectName < summaries[j].ObjectName
	})

	return summaries, nil
}

// buildQuery constrói a query baseada nos critérios
func (g *BudgetGatewayImpl) buildQuery(ctx context.Context, criteria interface{}) *gorm.DB {
	query := g.db.WithContext(ctx).Model(&model.BudgetItemModel{})

	if c, ok := criteria.(interface {
		GetObjectUUID() *uuid.UUID
		GetObjectName() *string
		GetGroupUUID() *uuid.UUID
		GetGroupName() *string
		GetCod() *string
		GetDesc() *string
		GetYear() *int
		GetPage() *int
		GetLimit() *int
		GetSortBy() *string
		GetSortOrder() *string
	}); ok {
		if c.GetObjectUUID() != nil {
			query = query.Where("object_uuid = ?", *c.GetObjectUUID())
		}
		if c.GetObjectName() != nil {
			query = query.Where("object_name ILIKE ?", "%"+*c.GetObjectName()+"%")
		}
		if c.GetGroupUUID() != nil {
			query = query.Where("group_uuid = ?", *c.GetGroupUUID())
		}
		if c.GetGroupName() != nil {
			query = query.Where("group_name ILIKE ?", "%"+*c.GetGroupName()+"%")
		}
		if c.GetCod() != nil {
			query = query.Where("cod = ?", *c.GetCod())
		}
		if c.GetDesc() != nil {
			query = query.Where("desc ILIKE ?", "%"+*c.GetDesc()+"%")
		}
		if c.GetYear() != nil {
			query = query.Where("year = ?", *c.GetYear())
		}

		// Aplicar ordenação
		query = g.applySorting(query, c.GetSortBy(), c.GetSortOrder())

		// Aplicar paginação
		if c.GetPage() != nil && c.GetLimit() != nil {
			offset := (*c.GetPage() - 1) * *c.GetLimit()
			query = query.Offset(offset).Limit(*c.GetLimit())
		}
	}

	return query
}

// applySorting aplica ordenação à query
func (g *BudgetGatewayImpl) applySorting(query *gorm.DB, sortBy, sortOrder *string) *gorm.DB {
	if sortBy == nil || *sortBy == "" {
		sortBy = new(string)
		*sortBy = "objectUUID"
	}

	if sortOrder == nil || *sortOrder == "" {
		sortOrder = new(string)
		*sortOrder = "asc"
	}

	column := ""
	switch *sortBy {
	case "objectUUID":
		column = "object_uuid"
	case "objectName":
		column = "object_name"
	case "groupUUID":
		column = "group_uuid"
	case "groupName":
		column = "group_name"
	case "cod":
		column = "cod"
	case "desc":
		column = "desc"
	case "createdAt":
		column = "created_at"
	default:
		column = "object_uuid"
	}

	order := *sortOrder
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	return query.Order(column + " " + order)
}
