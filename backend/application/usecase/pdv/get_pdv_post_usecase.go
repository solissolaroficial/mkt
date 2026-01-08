package pdv

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetPdvPostUseCase busca um post de PDV por ID
type GetPdvPostUseCase struct {
	gateway gateway.PdvPostGateway
}

// NewGetPdvPost cria uma nova instância do GetPdvPostUseCase
func NewGetPdvPost(gateway gateway.PdvPostGateway) *GetPdvPostUseCase {
	return &GetPdvPostUseCase{gateway: gateway}
}

// Execute busca um post de PDV por ID
func (uc *GetPdvPostUseCase) Execute(ctx context.Context, id string) (*entity.PdvPost, error) {
	post, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		return nil, domainErrors.ErrPdvPostNotFound
	}
	return post, nil
}
