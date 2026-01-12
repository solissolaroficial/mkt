package offlineaction

import (
	"context"

	domainErrors "solis/backend/core/domain/errors"
	"solis/backend/core/domain/gateway"
)

type DeleteOfflineActionUseCase struct {
	gateway gateway.OfflineActionGateway
}

func NewDeleteOfflineAction(gateway gateway.OfflineActionGateway) *DeleteOfflineActionUseCase {
	return &DeleteOfflineActionUseCase{gateway: gateway}
}

func (uc *DeleteOfflineActionUseCase) Execute(ctx context.Context, id string) error {
	// Verificar se existe
	exists, err := uc.gateway.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domainErrors.ErrOfflineActionNotFound
	}

	// Deletar (soft delete)
	return uc.gateway.Delete(ctx, id)
}
