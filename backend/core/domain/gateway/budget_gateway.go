package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// BudgetGateway define a interface para operações de persistência de BudgetItem
type BudgetGateway interface {
	// Create cria um novo BudgetItem no banco de dados
	Create(ctx context.Context, budget *entity.BudgetItem) error

	// FindByID busca um BudgetItem pelo ID
	// Retorna ErrBudgetNotFound se não encontrado
	FindByID(ctx context.Context, id uuid.UUID) (*entity.BudgetItem, error)

	// List busca BudgetItems baseados em critérios
	// Aceita *BudgetCriteria como parâmetro
	List(ctx context.Context, criteria interface{}) ([]*entity.BudgetItem, error)

	// Update atualiza um BudgetItem existente
	// Retorna ErrBudgetNotFound se não encontrado
	Update(ctx context.Context, budget *entity.BudgetItem) error

	// Delete remove um BudgetItem (soft delete)
	// Retorna ErrBudgetNotFound se não encontrado
	Delete(ctx context.Context, id uuid.UUID) error

	// Count conta o número de BudgetItems baseados em critérios
	Count(ctx context.Context, criteria interface{}) (int64, error)

	// ExistsByCode verifica se existe um item com o mesmo código
	// Considera codObj, codGrp, cod e year para unicidade
	ExistsByCode(ctx context.Context, codObj, codGrp, cod string, year int) (bool, error)

	// BatchCreate cria múltiplos BudgetItems em lote
	// Usa transação para garantir atomicidade
	BatchCreate(ctx context.Context, budgets []*entity.BudgetItem) error

	// GetDistinctYears retorna os anos disponíveis no banco
	GetDistinctYears(ctx context.Context) ([]int, error)

	// GetSummary retorna resumo agregado por objeto/grupo
	// Aceita *BudgetCriteria como parâmetro
	GetSummary(ctx context.Context, criteria interface{}) ([]*BudgetSummary, error)
}

// BudgetSummary representa um resumo agregado de orçamento
type BudgetSummary struct {
	CodObj        string  // Código do objeto
	Obj           string  // Nome do objeto
	CodGrp        string  // Código do grupo
	Grp           string  // Nome do grupo
	TotalBudget   float64 // Soma de todos os valores orçados
	TotalRealized float64 // Soma de todos os valores realizados
	Variance      float64 // Diferença (orçado - realizado)
}
