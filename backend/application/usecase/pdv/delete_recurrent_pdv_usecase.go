package pdv

import (
	"context"

	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteRecurrentPdvUseCase exclui um PDV recorrente (soft delete)
type DeleteRecurrentPdvUseCase struct {
	gateway gateway.RecurrentPdvGateway
}

// NewDeleteRecurrentPdv cria uma nova instância do DeleteRecurrentPdvUseCase
func NewDeleteRecurrentPdv(gateway gateway.RecurrentPdvGateway) *DeleteRecurrentPdvUseCase {
	return &DeleteRecurrentPdvUseCase{gateway: gateway}
}

// Execute exclui um PDV recorrente (soft delete)
func (uc *DeleteRecurrentPdvUseCase) Execute(ctx context.Context, id string) error {
	// Verificar se existe
	exists, err := uc.gateway.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domainErrors.ErrRecurrentPdvNotFound
	}

	// Deletar (soft delete)
	return uc.gateway.Delete(ctx, id)
}
