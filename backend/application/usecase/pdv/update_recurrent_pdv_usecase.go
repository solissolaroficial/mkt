package pdv

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateRecurrentPdvUseCase atualiza um PDV recorrente existente
type UpdateRecurrentPdvUseCase struct {
	gateway gateway.RecurrentPdvGateway
}

// UpdateRecurrentPdvInput representa os dados de entrada para atualizar um PDV recorrente
type UpdateRecurrentPdvInput struct {
	ID               string
	Name             *string
	RepName          *string
	City             *string
	Followers        *int
	InstagramProfile *string
}

// NewUpdateRecurrentPdv cria uma nova instância do UpdateRecurrentPdvUseCase
func NewUpdateRecurrentPdv(gateway gateway.RecurrentPdvGateway) *UpdateRecurrentPdvUseCase {
	return &UpdateRecurrentPdvUseCase{gateway: gateway}
}

// Execute atualiza um PDV recorrente existente
func (uc *UpdateRecurrentPdvUseCase) Execute(ctx context.Context, input UpdateRecurrentPdvInput) (*entity.RecurrentPdv, error) {
	// Buscar PDV existente
	pdv, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domainErrors.ErrRecurrentPdvNotFound
	}

	// Atualizar campos se fornecidos
	if input.Name != nil {
		if err := pdv.UpdateName(*input.Name); err != nil {
			return nil, err
		}
	}

	if input.RepName != nil {
		if err := pdv.UpdateRepName(*input.RepName); err != nil {
			return nil, err
		}
	}

	if input.City != nil {
		if err := pdv.UpdateCity(input.City); err != nil {
			return nil, err
		}
	}

	if input.Followers != nil {
		if err := pdv.UpdateFollowers(input.Followers); err != nil {
			return nil, err
		}
	}

	if input.InstagramProfile != nil {
		if err := pdv.UpdateInstagramProfile(input.InstagramProfile); err != nil {
			return nil, err
		}
	}

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, pdv); err != nil {
		return nil, err
	}

	return pdv, nil
}
