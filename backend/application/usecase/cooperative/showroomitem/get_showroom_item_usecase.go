package showroomitem

import (
	"context"

	"solis/backend/core/domain/entity"
	domainErrors "solis/backend/core/domain/errors"
	"solis/backend/core/domain/gateway"
)

type GetShowroomItemUseCase struct {
	gateway gateway.ShowroomItemGateway
}

func NewGetShowroomItem(gateway gateway.ShowroomItemGateway) *GetShowroomItemUseCase {
	return &GetShowroomItemUseCase{gateway: gateway}
}

func (uc *GetShowroomItemUseCase) Execute(ctx context.Context, id string) (*entity.ShowroomItem, error) {
	item, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, domainErrors.ErrShowroomItemNotFound
	}
	return item, nil
}
