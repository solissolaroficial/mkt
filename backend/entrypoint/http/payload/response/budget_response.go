package response

import (
	"time"

	"github.com/google/uuid"
)

// BudgetItemResponse representa a resposta de um BudgetItem
type BudgetItemResponse struct {
	UUID         uuid.UUID  `json:"uuid"`
	CodObj       string     `json:"codObj"`
	Obj          string     `json:"obj"`
	CodGrp       string     `json:"codGrp"`
	Grp          string     `json:"grp"`
	Cod          string     `json:"cod"`
	Desc         string     `json:"desc"`
	Vals         []float64  `json:"vals"`
	RealizedVals []float64  `json:"realizedVals"`
	Year         int        `json:"year"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
}

// BudgetSummaryResponse representa o resumo agregado de orçamento
type BudgetSummaryResponse struct {
	CodObj        string  `json:"codObj"`
	Obj           string  `json:"obj"`
	CodGrp        string  `json:"codGrp"`
	Grp           string  `json:"grp"`
	TotalBudget   float64 `json:"totalBudget"`
	TotalRealized float64 `json:"totalRealized"`
	Variance      float64 `json:"variance"`
}

// ListBudgetItemsResponse representa a resposta paginada de BudgetItems
type ListBudgetItemsResponse struct {
	Data  []BudgetItemResponse `json:"data"`
	Total int64                `json:"total"`
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
}
