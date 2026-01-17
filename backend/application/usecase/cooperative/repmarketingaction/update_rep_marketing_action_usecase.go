package repmarketingaction

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type UpdateRepMarketingActionUseCase struct {
	gateway gateway.RepMarketingActionGateway
}

type UpdateRepMarketingActionInput struct {
	ID                 string
	RepresentativeUUID *string
	Date               *string
	Description        *string
}

func NewUpdateRepMarketingAction(gateway gateway.RepMarketingActionGateway) *UpdateRepMarketingActionUseCase {
	return &UpdateRepMarketingActionUseCase{
		gateway: gateway,
	}
}

func (uc *UpdateRepMarketingActionUseCase) Execute(ctx context.Context, input UpdateRepMarketingActionInput) (*entity.RepMarketingAction, error) {
	// Buscar ação existente
	action, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domainErrors.ErrRepMarketingActionNotFound
	}

	// Atualizar campos se fornecidos
	if input.RepresentativeUUID != nil {
		representativeUUID, err := uuid.Parse(*input.RepresentativeUUID)
		if err == nil {
			if err := action.UpdateRepresentativeUUID(representativeUUID); err != nil {
				return nil, err
			}
		}
	}

	if input.Date != nil {
		date, err := time.Parse("2006-01-02", *input.Date)
		if err != nil {
			return nil, err
		}
		if err := action.UpdateDate(date); err != nil {
			return nil, err
		}
	}

	if input.Description != nil {
		if err := action.UpdateDescription(*input.Description); err != nil {
			return nil, err
		}
	}

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, action); err != nil {
		return nil, err
	}

	return action, nil
}
