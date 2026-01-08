package pdv

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateRecurrentPdvUseCase cria um novo PDV recorrente
type CreateRecurrentPdvUseCase struct {
	gateway gateway.RecurrentPdvGateway
}

// CreateRecurrentPdvInput representa os dados de entrada para criar um PDV recorrente
type CreateRecurrentPdvInput struct {
	Name    string
	RepName string
	City    *string
}

// NewCreateRecurrentPdv cria uma nova instância do CreateRecurrentPdvUseCase
func NewCreateRecurrentPdv(gateway gateway.RecurrentPdvGateway) *CreateRecurrentPdvUseCase {
	return &CreateRecurrentPdvUseCase{gateway: gateway}
}

// Execute cria um novo PDV recorrente
func (uc *CreateRecurrentPdvUseCase) Execute(ctx context.Context, input CreateRecurrentPdvInput) (*entity.RecurrentPdv, error) {
	// Criar entidade
	pdv, err := entity.NewRecurrentPdv(input.Name, input.RepName)
	if err != nil {
		return nil, err
	}

	// Atualizar cidade se fornecida
	if input.City != nil {
		if err := pdv.UpdateCity(input.City); err != nil {
			return nil, err
		}
	}

	// Salvar via gateway
	if err := uc.gateway.Save(ctx, pdv); err != nil {
		return nil, err
	}

	return pdv, nil
}
