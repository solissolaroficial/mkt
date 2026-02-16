package gateway

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// FlowGateway define a interface para operações de persistência de Flow
type FlowGateway interface {
	// Create cria um novo fluxo
	Create(flow *entity.Flow) (*entity.Flow, error)

	// GetByUUID busca um fluxo por UUID
	GetByUUID(uuid string) (*entity.Flow, error)

	// GetAll lista todos os fluxos ativos
	GetAll() ([]*entity.Flow, error)

	// Update atualiza um fluxo
	Update(flow *entity.Flow) (*entity.Flow, error)

	// Delete remove um fluxo (soft delete)
	Delete(id string) error

	// Reorder reordena os fluxos
	Reorder(flowIDs []string) error
}
