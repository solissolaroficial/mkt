package credentials

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type ListCredentialsUseCase struct {
	gateway gateway.ProgramCredentialGateway
}

func NewListCredentialsUseCase(gateway gateway.ProgramCredentialGateway) *ListCredentialsUseCase {
	return &ListCredentialsUseCase{
		gateway: gateway,
	}
}

func (uc *ListCredentialsUseCase) Execute(ctx context.Context) ([]*entity.ProgramCredential, error) {
	credentials, err := uc.gateway.FindActive(ctx)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}
