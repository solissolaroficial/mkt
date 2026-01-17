package repmarketingaction

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type CreateRepMarketingActionUseCase struct {
	gateway gateway.RepMarketingActionGateway
}

type CreateRepMarketingActionInput struct {
	RepresentativeUUID string
	Date               string
	Description        string
}

func NewCreateRepMarketingAction(gateway gateway.RepMarketingActionGateway) *CreateRepMarketingActionUseCase {
	return &CreateRepMarketingActionUseCase{
		gateway: gateway,
	}
}

func (uc *CreateRepMarketingActionUseCase) Execute(ctx context.Context, input CreateRepMarketingActionInput) (*entity.RepMarketingAction, error) {
	// Parse date
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, err
	}

	// Parse representative UUID
	representativeUUID, err := uuid.Parse(input.RepresentativeUUID)
	if err != nil {
		return nil, err
	}

	// Criar entidade
	action, err := entity.NewRepMarketingAction(
		representativeUUID,
		date,
		input.Description,
	)
	if err != nil {
		return nil, err
	}

	// Salvar via gateway
	if err := uc.gateway.Save(ctx, action); err != nil {
		return nil, err
	}

	return action, nil
}
