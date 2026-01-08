package pdv

import (
	"context"

	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeletePdvPostUseCase exclui um post de PDV (soft delete)
type DeletePdvPostUseCase struct {
	gateway gateway.PdvPostGateway
}

// NewDeletePdvPost cria uma nova instância do DeletePdvPostUseCase
func NewDeletePdvPost(gateway gateway.PdvPostGateway) *DeletePdvPostUseCase {
	return &DeletePdvPostUseCase{gateway: gateway}
}

// Execute exclui um post de PDV (soft delete)
func (uc *DeletePdvPostUseCase) Execute(ctx context.Context, id string) error {
	// Verificar se existe
	exists, err := uc.gateway.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domainErrors.ErrPdvPostNotFound
	}

	// Deletar (soft delete)
	return uc.gateway.Delete(ctx, id)
}
