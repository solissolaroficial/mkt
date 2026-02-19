package credentials

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type DeleteCredentialUseCase struct {
	gateway gateway.ProgramCredentialGateway
}

type DeleteCredentialInput struct {
	ID uuid.UUID
}

func NewDeleteCredentialUseCase(gateway gateway.ProgramCredentialGateway) *DeleteCredentialUseCase {
	return &DeleteCredentialUseCase{
		gateway: gateway,
	}
}

func (uc *DeleteCredentialUseCase) Execute(ctx context.Context, input DeleteCredentialInput) error {
	// Verify credential exists
	_, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	// Delete credential
	return uc.gateway.Delete(ctx, input.ID)
}
