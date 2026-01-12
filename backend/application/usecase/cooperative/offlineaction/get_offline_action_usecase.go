package offlineaction

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type GetOfflineActionUseCase struct {
	gateway gateway.OfflineActionGateway
}

func NewGetOfflineAction(gateway gateway.OfflineActionGateway) *GetOfflineActionUseCase {
	return &GetOfflineActionUseCase{gateway: gateway}
}

func (uc *GetOfflineActionUseCase) Execute(ctx context.Context, id string) (*entity.OfflineAction, error) {
	action, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, domainErrors.ErrOfflineActionNotFound
	}
	return action, nil
}
