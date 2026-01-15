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
func (g *BudgetGatewayImpl) List(ctx context.Context, criteria interface{}) ([]*entity.BudgetItem, error) {
	query := g.buildQuery(ctx, criteria)

	var models []*model.BudgetItemModel
	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	return mapper.ToBudgetItemDomainList(models)
}

// Update atualiza um BudgetItem existente
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
			"cod_obj":       budgetModel.CodObj,
			"obj":           budgetModel.Obj,
			"cod_grp":       budgetModel.CodGrp,
			"grp":           budgetModel.Grp,
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
func (g *BudgetGatewayImpl) ExistsByCode(ctx context.Context, codObj, codGrp, cod string, year int) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).Model(&model.BudgetItemModel{}).
		Where("cod_obj = ? AND cod_grp = ? AND cod = ? AND year = ?", codObj, codGrp, cod, year).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// BatchCreate cria múltiplos BudgetItems em lote
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

// GetDistinctYears retorna os anos disponíveis
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
func (g *BudgetGatewayImpl) GetSummary(ctx context.Context, criteria interface{}) ([]*gateway.BudgetSummary, error) {
	// Buscar todos os items filtrados
	items, err := g.List(ctx, criteria)
	if err != nil {
		return nil, err
	}

	// Agrupar em memória
	summaryMap := make(map[string]*gateway.BudgetSummary)

	for _, item := range items {
		key := item.CodObj() + "|" + item.CodGrp()

		if summary, exists := summaryMap[key]; exists {
			summary.TotalBudget += item.GetTotalBudget()
			summary.TotalRealized += item.GetTotalRealized()
		} else {
			summaryMap[key] = &gateway.BudgetSummary{
				CodObj:        item.CodObj(),
				Obj:           item.Obj(),
				CodGrp:        item.CodGrp(),
				Grp:           item.Grp(),
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
		if summaries[i].CodObj == summaries[j].CodObj {
			return summaries[i].CodGrp < summaries[j].CodGrp
		}
		return summaries[i].CodObj < summaries[j].CodObj
	})

	return summaries, nil
}

// buildQuery constrói a query baseada nos critérios
func (g *BudgetGatewayImpl) buildQuery(ctx context.Context, criteria interface{}) *gorm.DB {
	query := g.db.WithContext(ctx).Model(&model.BudgetItemModel{})

	if c, ok := criteria.(interface {
		GetCodObj() *string
		GetObj() *string
		GetCodGrp() *string
		GetGrp() *string
		GetCod() *string
		GetDesc() *string
		GetYear() *int
		GetPage() *int
		GetLimit() *int
		GetSortBy() *string
		GetSortOrder() *string
	}); ok {
		if c.GetCodObj() != nil {
			query = query.Where("cod_obj = ?", *c.GetCodObj())
		}
		if c.GetObj() != nil {
			query = query.Where("obj ILIKE ?", "%"+*c.GetObj()+"%")
		}
		if c.GetCodGrp() != nil {
			query = query.Where("cod_grp = ?", *c.GetCodGrp())
		}
		if c.GetGrp() != nil {
			query = query.Where("grp ILIKE ?", "%"+*c.GetGrp()+"%")
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
		*sortBy = "codObj"
	}

	if sortOrder == nil || *sortOrder == "" {
		sortOrder = new(string)
		*sortOrder = "asc"
	}

	column := ""
	switch *sortBy {
	case "codObj":
		column = "cod_obj"
	case "obj":
		column = "obj"
	case "codGrp":
		column = "cod_grp"
	case "grp":
		column = "grp"
	case "cod":
		column = "cod"
	case "desc":
		column = "desc"
	case "createdAt":
		column = "created_at"
	default:
		column = "cod_obj"
	}

	order := *sortOrder
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	return query.Order(column + " " + order)
}
