package social

import (
	"context"
	"errors"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type GetSocialBenchmarkingUseCase struct {
	gateway gateway.SocialBenchmarkingGateway
}

func NewGetSocialBenchmarking(gateway gateway.SocialBenchmarkingGateway) *GetSocialBenchmarkingUseCase {
	return &GetSocialBenchmarkingUseCase{gateway: gateway}
}

func (uc *GetSocialBenchmarkingUseCase) Execute(ctx context.Context, id string) (*entity.SocialBenchmarking, error) {
	benchmarking, err := uc.gateway.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrSocialBenchmarkingNotFound) {
			return nil, domainErrors.ErrSocialBenchmarkingNotFound
		}
		return nil, err
	}
	return benchmarking, nil
}
