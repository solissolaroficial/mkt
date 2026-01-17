package offlineaction

import (
	"context"

	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type CreateOfflineActionUseCase struct {
	gateway gateway.OfflineActionGateway
}

type CreateOfflineActionInput struct {
	RequestedAmount    float64
	ActionDate         string
	Category           string
	PDV                string
	RepresentativeUUID string
	Observation        string
}

func NewCreateOfflineAction(gateway gateway.OfflineActionGateway) *CreateOfflineActionUseCase {
	return &CreateOfflineActionUseCase{
		gateway: gateway,
	}
}

func (uc *CreateOfflineActionUseCase) Execute(ctx context.Context, input CreateOfflineActionInput) (*entity.OfflineAction, error) {
	// Validar e criar value objects
	actionDate, err := valueobject.NewActionDate(input.ActionDate)
	if err != nil {
		return nil, err
	}

	category, err := valueobject.NewOfflineCategory(input.Category)
	if err != nil {
		return nil, err
	}

	requestedAmount, err := valueobject.NewMonetaryValue(input.RequestedAmount)
	if err != nil {
		return nil, err
	}

	representativeUUID, err := uuid.Parse(input.RepresentativeUUID)
	if err != nil {
		return nil, err
	}

	pdv := ""
	if input.PDV != "" {
		pdv = input.PDV
	}

	observation := input.Observation

	// Criar entidade
	action, err := entity.NewOfflineAction(
		requestedAmount.Value(),
		actionDate,
		category,
		pdv,
		representativeUUID,
		observation,
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
