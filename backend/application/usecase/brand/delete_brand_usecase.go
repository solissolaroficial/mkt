package brand

import (
	"context"

	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type DeleteBrandUseCase struct {
	gateway gateway.BrandGateway
}

func NewDeleteBrandUseCase(gateway gateway.BrandGateway) *DeleteBrandUseCase {
	return &DeleteBrandUseCase{
		gateway: gateway,
	}
}

func (uc *DeleteBrandUseCase) Execute(ctx context.Context, id uuid.UUID) error {
	return uc.gateway.Delete(ctx, id)
}
