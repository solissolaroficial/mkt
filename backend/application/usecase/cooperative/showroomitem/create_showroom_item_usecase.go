package showroomitem

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type CreateShowroomItemUseCase struct {
	gateway gateway.ShowroomItemGateway
}

type CreateShowroomItemInput struct {
	PDV              string
	RepName          string
	City             *string
	Contact          *string
	DeliveryForecast *string
	WorkshopDate     *string
}

func NewCreateShowroomItem(gateway gateway.ShowroomItemGateway) *CreateShowroomItemUseCase {
	return &CreateShowroomItemUseCase{gateway: gateway}
}

func (uc *CreateShowroomItemUseCase) Execute(ctx context.Context, input CreateShowroomItemInput) (*entity.ShowroomItem, error) {
	// Criar entidade (validação interna)
	item, err := entity.NewShowroomItem(
		input.PDV,
		input.RepName,
	)
	if err != nil {
		return nil, err
	}

	// Atualizar campos opcionais se fornecidos
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

	// Salvar via gateway
	if err := uc.gateway.Save(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}
