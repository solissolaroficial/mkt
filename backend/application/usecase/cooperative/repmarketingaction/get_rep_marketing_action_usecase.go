package repmarketingaction

import (
	"context"

	"solis/backend/core/domain/entity"
	domainErrors "solis/backend/core/domain/errors"
	"solis/backend/core/domain/gateway"
)

type GetRepMarketingActionUseCase struct {
	gateway gateway.RepMarketingActionGateway
}

func NewGetRepMarketingAction(gateway gateway.RepMarketingActionGateway) *GetRepMarketingActionUseCase {
	return &GetRepMarketingActionUseCase{gateway: gateway}
}

func (uc *GetRepMarketingActionUseCase) Execute(ctx context.Context, id string) (*entity.RepMarketingAction, error) {
	action, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, domainErrors.ErrRepMarketingActionNotFound
	}
	return action, nil
}
