package calendar

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type GetCalendarPostUseCase struct {
	gateway gateway.CalendarPostGateway
}

func NewGetCalendarPost(gateway gateway.CalendarPostGateway) *GetCalendarPostUseCase {
	return &GetCalendarPostUseCase{gateway: gateway}
}

func (uc *GetCalendarPostUseCase) Execute(ctx context.Context, id string) (*entity.CalendarPost, error) {
	post, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, domainErrors.ErrCalendarPostNotFound
	}
	return post, nil
}
