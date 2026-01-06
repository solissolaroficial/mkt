package calendar

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type DeleteCalendarPostUseCase struct {
	gateway gateway.CalendarPostGateway
}

func NewDeleteCalendarPost(gateway gateway.CalendarPostGateway) *DeleteCalendarPostUseCase {
	return &DeleteCalendarPostUseCase{gateway: gateway}
}

func (uc *DeleteCalendarPostUseCase) Execute(ctx context.Context, id string) error {
	// Verificar se existe
	exists, err := uc.gateway.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.ErrCalendarPostNotFound
	}

	// Deletar (soft delete)
	return uc.gateway.Delete(ctx, id)
}
