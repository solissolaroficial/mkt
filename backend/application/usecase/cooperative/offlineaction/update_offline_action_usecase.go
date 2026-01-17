package offlineaction

import (
	"context"

	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type UpdateOfflineActionUseCase struct {
	gateway gateway.OfflineActionGateway
}

type UpdateOfflineActionInput struct {
	ID                 string
	ApprovedAmount     *string
	OrderNumber        *string
	DepartureDate      *string
	DeliveryForecast   *string
	DeliveryDate       *string
	City               *string
	UF                 *string
	Scored             *string
	Status             *string
	Observation        *string
	PDV                *string
	RepresentativeUUID *string
}

func NewUpdateOfflineAction(gateway gateway.OfflineActionGateway) *UpdateOfflineActionUseCase {
	return &UpdateOfflineActionUseCase{
		gateway: gateway,
	}
}

func (uc *UpdateOfflineActionUseCase) Execute(ctx context.Context, input UpdateOfflineActionInput) (*entity.OfflineAction, error) {
	// Buscar ação existente
	action, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domainErrors.ErrOfflineActionNotFound
	}

	// Atualizar campos se fornecidos
	if input.ApprovedAmount != nil {
		if err := action.UpdateApprovedAmount(input.ApprovedAmount); err != nil {
			return nil, err
		}
	}

	if input.OrderNumber != nil {
		if err := action.UpdateOrderNumber(input.OrderNumber); err != nil {
			return nil, err
		}
	}

	if input.DepartureDate != nil {
		if err := action.UpdateDepartureDate(input.DepartureDate); err != nil {
			return nil, err
		}
	}

	if input.DeliveryForecast != nil {
		if err := action.UpdateDeliveryForecast(input.DeliveryForecast); err != nil {
			return nil, err
		}
	}

	if input.DeliveryDate != nil {
		if err := action.UpdateDeliveryDate(input.DeliveryDate); err != nil {
			return nil, err
		}
	}

	if input.City != nil || input.UF != nil {
		if err := action.UpdateLocation(input.City, input.UF); err != nil {
			return nil, err
		}
	}

	if input.Scored != nil {
		scored, err := valueobject.NewScoredStatus(*input.Scored)
		if err == nil {
			if err := action.UpdateScored(&scored); err != nil {
				return nil, err
			}
		}
	}

	if input.Status != nil {
		status, err := valueobject.NewOfflineStatus(*input.Status)
		if err == nil {
			if err := action.UpdateStatus(&status); err != nil {
				return nil, err
			}
		}
	}

	if input.Observation != nil {
		if err := action.UpdateObservation(input.Observation); err != nil {
			return nil, err
		}
	}

	if input.PDV != nil {
		if err := action.UpdatePDV(input.PDV); err != nil {
			return nil, err
		}
	}

	if input.RepresentativeUUID != nil {
		representativeUUID, err := uuid.Parse(*input.RepresentativeUUID)
		if err == nil {
			if err := action.UpdateRepresentativeUUID(representativeUUID); err != nil {
				return nil, err
			}
		}
	}

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, action); err != nil {
		return nil, err
	}

	return action, nil
}
