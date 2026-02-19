package credentials

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type UpdateCredentialUseCase struct {
	gateway gateway.ProgramCredentialGateway
}

type UpdateCredentialInput struct {
	ID       uuid.UUID
	Name     string
	User     string
	Password string
	Access   string
	Notes    string
}

func NewUpdateCredentialUseCase(gateway gateway.ProgramCredentialGateway) *UpdateCredentialUseCase {
	return &UpdateCredentialUseCase{
		gateway: gateway,
	}
}

func (uc *UpdateCredentialUseCase) Execute(ctx context.Context, input UpdateCredentialInput) (*entity.ProgramCredential, error) {
	// Find existing credential
	credential, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	// Update fields
	if err := credential.Update(
		input.Name,
		input.User,
		input.Password,
		input.Access,
		input.Notes,
	); err != nil {
		return nil, err
	}

	// Save changes
	if err := uc.gateway.Update(ctx, credential); err != nil {
		return nil, err
	}

	return credential, nil
}
