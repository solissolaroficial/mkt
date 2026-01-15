package request

import (
	"github.com/seu-usuario/solis-backend/core/domain"
)

// CreateBudgetItemRequest representa o payload para criar um BudgetItem
type CreateBudgetItemRequest struct {
	CodObj       string    `json:"codObj" validate:"required,max=20"`
	Obj          string    `json:"obj" validate:"required,max=200"`
	CodGrp       string    `json:"codGrp" validate:"required,max=20"`
	Grp          string    `json:"grp" validate:"required,max=200"`
	Cod          string    `json:"cod" validate:"required,max=20"`
	Desc         string    `json:"desc" validate:"required,max=500"`
	Vals         []float64 `json:"vals" validate:"required,len=12"`
	RealizedVals []float64 `json:"realizedVals" validate:"required,len=12"`
	Year         int       `json:"year" validate:"required,min=2000,max=2100"`
}

// UpdateBudgetItemRequest representa o payload para atualizar um BudgetItem
type UpdateBudgetItemRequest struct {
	CodObj       *string    `json:"codObj,omitempty" validate:"omitempty,max=20"`
	Obj          *string    `json:"obj,omitempty" validate:"omitempty,max=200"`
	CodGrp       *string    `json:"codGrp,omitempty" validate:"omitempty,max=20"`
	Grp          *string    `json:"grp,omitempty" validate:"omitempty,max=200"`
	Cod          *string    `json:"cod,omitempty" validate:"omitempty,max=20"`
	Desc         *string    `json:"desc,omitempty" validate:"omitempty,max=500"`
	Vals         *[]float64 `json:"vals,omitempty" validate:"omitempty,len=12"`
	RealizedVals *[]float64 `json:"realizedVals,omitempty" validate:"omitempty,len=12"`
	Year         *int       `json:"year,omitempty" validate:"omitempty,min=2000,max=2100"`
}

// BatchCreateBudgetItemsRequest representa o payload para criar múltiplos BudgetItems
type BatchCreateBudgetItemsRequest struct {
	Items []CreateBudgetItemRequest `json:"items" validate:"required,min=1"`
}

// ListBudgetItemsQuery representa os parâmetros de query para listar BudgetItems
type ListBudgetItemsQuery struct {
	CodObj    *string `query:"codObj"`
	Obj       *string `query:"obj"`
	CodGrp    *string `query:"codGrp"`
	Grp       *string `query:"grp"`
	Cod       *string `query:"cod"`
	Desc      *string `query:"desc"`
	Year      *int    `query:"year"`
	Page      *int    `query:"page" validate:"omitempty,min=1"`
	Limit     *int    `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sortBy" validate:"omitempty,oneof=codObj obj codGrp grp cod desc createdAt"`
	SortOrder *string `query:"sortOrder" validate:"omitempty,oneof=asc desc"`
}

// ToCriteria converte os parâmetros de query para BudgetCriteria
func (q *ListBudgetItemsQuery) ToCriteria() *domain.BudgetCriteria {
	criteria := domain.NewBudgetCriteria()

	if q.CodObj != nil {
		criteria.WithCodObj(*q.CodObj)
	}
	if q.Obj != nil {
		criteria.WithObj(*q.Obj)
	}
	if q.CodGrp != nil {
		criteria.WithCodGrp(*q.CodGrp)
	}
	if q.Grp != nil {
		criteria.WithGrp(*q.Grp)
	}
	if q.Cod != nil {
		criteria.WithCod(*q.Cod)
	}
	if q.Desc != nil {
		criteria.WithDesc(*q.Desc)
	}
	if q.Year != nil {
		criteria.WithYear(*q.Year)
	}
	if q.Page != nil {
		criteria.WithPage(*q.Page)
	}
	if q.Limit != nil {
		criteria.WithLimit(*q.Limit)
	}
	if q.SortBy != nil {
		criteria.WithSortBy(*q.SortBy)
	}
	if q.SortOrder != nil {
		criteria.WithSortOrder(*q.SortOrder)
	}

	return criteria
}
