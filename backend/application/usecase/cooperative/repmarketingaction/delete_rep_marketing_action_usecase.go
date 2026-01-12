package repmarketingaction

import (
	"context"

	domainErrors "solis/backend/core/domain/errors"
	"solis/backend/core/domain/gateway"
)

type DeleteRepMarketingActionUseCase struct {
	gateway gateway.RepMarketingActionGateway
}

func NewDeleteRepMarketingAction(gateway gateway.RepMarketingActionGateway) *DeleteRepMarketingActionUseCase {
	return &DeleteRepMarketingActionUseCase{gateway: gateway}
}

func (uc *DeleteRepMarketingActionUseCase) Execute(ctx context.Context, id string) error {
	// Verificar se existe
	exists, err := uc.gateway.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domainErrors.ErrRepMarketingActionNotFound
	}

	// Deletar (soft delete)
	return uc.gateway.Delete(ctx, id)
}
