package brand

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type ListBrandsUseCase struct {
	gateway gateway.BrandGateway
}

func NewListBrandsUseCase(gateway gateway.BrandGateway) *ListBrandsUseCase {
	return &ListBrandsUseCase{
		gateway: gateway,
	}
}

func (uc *ListBrandsUseCase) Execute(ctx context.Context) ([]*entity.Brand, error) {
	brands, err := uc.gateway.List(ctx)
	if err != nil {
		return nil, err
	}

	return brands, nil
}
