package credentials

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type GetCredentialUseCase struct {
	gateway gateway.ProgramCredentialGateway
}

type GetCredentialInput struct {
	ID uuid.UUID
}

func NewGetCredentialUseCase(gateway gateway.ProgramCredentialGateway) *GetCredentialUseCase {
	return &GetCredentialUseCase{
		gateway: gateway,
	}
}

func (uc *GetCredentialUseCase) Execute(ctx context.Context, input GetCredentialInput) (*entity.ProgramCredential, error) {
	credential, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return credential, nil
}
