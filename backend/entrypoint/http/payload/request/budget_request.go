package request

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
)

// CreateBudgetItemRequest representa o payload para criar um BudgetItem
type CreateBudgetItemRequest struct {
	ObjectUUID   *string   `json:"objectUUID" validate:"omitempty,uuid"`
	ObjectName   string    `json:"objectName" validate:"required,max=200"`
	GroupUUID    *string   `json:"groupUUID" validate:"omitempty,uuid"`
	GroupName    string    `json:"groupName" validate:"required,max=200"`
	Cod          string    `json:"cod" validate:"required,max=20"`
	Desc         string    `json:"desc" validate:"required,max=500"`
	Vals         []float64 `json:"vals" validate:"required,len=12"`
	RealizedVals []float64 `json:"realizedVals" validate:"required,len=12"`
	Year         int       `json:"year" validate:"required,min=2000,max=2100"`
}

// UpdateBudgetItemRequest representa o payload para atualizar um BudgetItem
type UpdateBudgetItemRequest struct {
	ObjectUUID   *string    `json:"objectUUID,omitempty" validate:"omitempty,uuid"`
	ObjectName   string     `json:"objectName,omitempty" validate:"omitempty,max=200"`
	GroupUUID    *string    `json:"groupUUID,omitempty" validate:"omitempty,uuid"`
	GroupName    string     `json:"groupName,omitempty" validate:"omitempty,max=200"`
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
	ObjectUUID *string `query:"objectUUID"`
	ObjectName *string `query:"objectName"`
	GroupUUID  *string `query:"groupUUID"`
	GroupName  *string `query:"groupName"`
	Cod        *string `query:"cod"`
	Desc       *string `query:"desc"`
	Year       *int    `query:"year"`
	Page       *int    `query:"page" validate:"omitempty,min=1"`
	Limit      *int    `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy     *string `query:"sortBy" validate:"omitempty,oneof=objectUUID objectName groupUUID groupName cod desc createdAt"`
	SortOrder  *string `query:"sortOrder" validate:"omitempty,oneof=asc desc"`
}

// ToCriteria converte os parâmetros de query para BudgetCriteria
func (q *ListBudgetItemsQuery) ToCriteria() *domain.BudgetCriteria {
	criteria := domain.NewBudgetCriteria()

	if q.ObjectUUID != nil {
		parsedUUID, err := uuid.Parse(*q.ObjectUUID)
		if err != nil {
			// Em caso de erro, continua sem esse filtro
			// O erro será tratado no usecase
		} else {
			criteria.WithObjectUUID(&parsedUUID)
		}
	}

	if q.ObjectName != nil {
		criteria.WithObjectName(q.ObjectName)
	}

	if q.GroupUUID != nil {
		parsedUUID, err := uuid.Parse(*q.GroupUUID)
		if err != nil {
			// Em caso de erro, continua sem esse filtro
			// O erro será tratado no usecase
		} else {
			criteria.WithGroupUUID(&parsedUUID)
		}
	}

	if q.GroupName != nil {
		criteria.WithGroupName(q.GroupName)
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
