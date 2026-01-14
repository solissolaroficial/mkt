package gateway

import (
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// AccountPayableGateway define a interface para operações de banco de dados de AccountPayable
type AccountPayableGateway interface {
	// Create cria um novo AccountPayable
	Create(account *entity.AccountPayable) error

	// FindByID busca um AccountPayable pelo ID
	FindByID(id uuid.UUID) (*entity.AccountPayable, error)

	// List busca AccountPayables baseados em critérios
	List(criteria interface{}) ([]*entity.AccountPayable, error)

	// Update atualiza um AccountPayable existente
	Update(account *entity.AccountPayable) error

	// Delete remove um AccountPayable (soft delete)
	Delete(id uuid.UUID) error

	// Count conta o número de AccountPayables baseados em critérios
	Count(criteria interface{}) (int64, error)

	// ExistsBySupplierAndDueDate verifica se existe uma conta com o mesmo fornecedor e data de vencimento
	ExistsBySupplierAndDueDate(supplier string, dueDate time.Time) (bool, error)
}
