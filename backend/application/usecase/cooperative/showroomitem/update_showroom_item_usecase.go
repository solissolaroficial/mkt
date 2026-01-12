package showroomitem

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type UpdateShowroomItemUseCase struct {
	gateway gateway.ShowroomItemGateway
}

type UpdateShowroomItemInput struct {
	ID               string
	PDV              *string
	RepName          *string
	City             *string
	Contact          *string
	DeliveryForecast *string
	WorkshopDate     *string
	Delivered        *bool
}

func NewUpdateShowroomItem(gateway gateway.ShowroomItemGateway) *UpdateShowroomItemUseCase {
	return &UpdateShowroomItemUseCase{gateway: gateway}
}

func (uc *UpdateShowroomItemUseCase) Execute(ctx context.Context, input UpdateShowroomItemInput) (*entity.ShowroomItem, error) {
	// Buscar item existente
	item, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domainErrors.ErrShowroomItemNotFound
	}

	// Atualizar campos se fornecidos
	if input.PDV != nil {
		if err := item.UpdatePDV(*input.PDV); err != nil {
			return nil, err
		}
	}

	if input.RepName != nil {
		if err := item.UpdateRepName(*input.RepName); err != nil {
			return nil, err
		}
	}

	if input.City != nil {
		if err := item.UpdateCity(input.City); err != nil {
			return nil, err
		}
	}

	if input.Contact != nil {
		if err := item.UpdateContact(input.Contact); err != nil {
			return nil, err
		}
	}

	if input.DeliveryForecast != nil {
		if err := item.UpdateDeliveryForecast(input.DeliveryForecast); err != nil {
			return nil, err
		}
	}

	if input.WorkshopDate != nil {
		if err := item.UpdateWorkshopDate(input.WorkshopDate); err != nil {
			return nil, err
		}
	}

	if input.Delivered != nil {
		if *input.Delivered {
			item.MarkAsDelivered()
		} else {
			item.MarkAsNotDelivered()
		}
	}

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}
