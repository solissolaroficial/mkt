package credentials

import (
	"context"
	"errors"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type CreateCredentialUseCase struct {
	gateway gateway.ProgramCredentialGateway
}

type CreateCredentialInput struct {
	Name     string
	User     string
	Password string
	Access   string
	Notes    string
}

func NewCreateCredentialUseCase(gateway gateway.ProgramCredentialGateway) *CreateCredentialUseCase {
	return &CreateCredentialUseCase{
		gateway: gateway,
	}
}

func (uc *CreateCredentialUseCase) Execute(ctx context.Context, input CreateCredentialInput) (*entity.ProgramCredential, error) {
	// Validate required fields
	if input.Name == "" {
		return nil, errors.New("name is required")
	}

	// Check if credential with same name already exists
	exists, err := uc.gateway.ExistsByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domainErrors.ErrProgramCredentialNameExists
	}

	// Create entity
	credential, err := entity.NewProgramCredential(
		input.Name,
		input.User,
		input.Password,
		input.Access,
		input.Notes,
	)
	if err != nil {
		return nil, err
	}

	// Save
	if err := uc.gateway.Save(ctx, credential); err != nil {
		return nil, err
	}

	return credential, nil
}
