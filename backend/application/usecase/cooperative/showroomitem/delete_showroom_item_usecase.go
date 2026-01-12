package showroomitem

import (
	"context"

	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type DeleteShowroomItemUseCase struct {
	gateway gateway.ShowroomItemGateway
}

func NewDeleteShowroomItem(gateway gateway.ShowroomItemGateway) *DeleteShowroomItemUseCase {
	return &DeleteShowroomItemUseCase{gateway: gateway}
}

func (uc *DeleteShowroomItemUseCase) Execute(ctx context.Context, id string) error {
	// Verificar se existe
	exists, err := uc.gateway.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domainErrors.ErrShowroomItemNotFound
	}

	// Deletar (soft delete)
	return uc.gateway.Delete(ctx, id)
}
