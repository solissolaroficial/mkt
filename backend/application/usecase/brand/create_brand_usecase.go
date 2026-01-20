package brand

import (
	"context"
	"errors"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type CreateBrandUseCase struct {
	gateway gateway.BrandGateway
}

type CreateBrandInput struct {
	Name string
}

func NewCreateBrandUseCase(gateway gateway.BrandGateway) *CreateBrandUseCase {
	return &CreateBrandUseCase{
		gateway: gateway,
	}
}

func (uc *CreateBrandUseCase) Execute(ctx context.Context, input CreateBrandInput) (*entity.Brand, error) {
	// Validar se já existe
	exists, err := uc.gateway.ExistsByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("brand with this name already exists")
	}

	// Criar entidade
	brand, err := entity.NewBrand(input.Name)
	if err != nil {
		return nil, err
	}

	// Salvar
	if err := uc.gateway.Save(ctx, brand); err != nil {
		return nil, err
	}

	return brand, nil
}
