package pdv

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetRecurrentPdvUseCase busca um PDV recorrente por ID
type GetRecurrentPdvUseCase struct {
	gateway gateway.RecurrentPdvGateway
}

// NewGetRecurrentPdv cria uma nova instância do GetRecurrentPdvUseCase
func NewGetRecurrentPdv(gateway gateway.RecurrentPdvGateway) *GetRecurrentPdvUseCase {
	return &GetRecurrentPdvUseCase{gateway: gateway}
}

// Execute busca um PDV recorrente por ID
func (uc *GetRecurrentPdvUseCase) Execute(ctx context.Context, id string) (*entity.RecurrentPdv, error) {
	pdv, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, domainErrors.ErrRecurrentPdvNotFound
	}
	return pdv, nil
}
