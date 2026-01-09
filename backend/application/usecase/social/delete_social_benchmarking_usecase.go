package social

import (
	"context"

	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type DeleteSocialBenchmarkingUseCase struct {
	gateway gateway.SocialBenchmarkingGateway
}

func NewDeleteSocialBenchmarking(gateway gateway.SocialBenchmarkingGateway) *DeleteSocialBenchmarkingUseCase {
	return &DeleteSocialBenchmarkingUseCase{gateway: gateway}
}

func (uc *DeleteSocialBenchmarkingUseCase) Execute(ctx context.Context, id string) error {
	// Verificar se existe
	exists, err := uc.gateway.ExistsByID(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return domainErrors.ErrSocialBenchmarkingNotFound
	}

	// Deletar (soft delete)
	return uc.gateway.Delete(ctx, id)
}
