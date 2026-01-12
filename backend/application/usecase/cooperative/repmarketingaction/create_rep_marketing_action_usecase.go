package repmarketingaction

import (
	"context"
	"time"

	"solis/backend/core/domain/entity"
	"solis/backend/core/domain/gateway"
)

type CreateRepMarketingActionUseCase struct {
	gateway gateway.RepMarketingActionGateway
}

type CreateRepMarketingActionInput struct {
	RepName     string
	Date        string
	Description string
}

func NewCreateRepMarketingAction(gateway gateway.RepMarketingActionGateway) *CreateRepMarketingActionUseCase {
	return &CreateRepMarketingActionUseCase{gateway: gateway}
}

func (uc *CreateRepMarketingActionUseCase) Execute(ctx context.Context, input CreateRepMarketingActionInput) (*entity.RepMarketingAction, error) {
	// Parse date
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, err
	}

	// Criar entidade (validação interna)
	action, err := entity.NewRepMarketingAction(
		input.RepName,
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
