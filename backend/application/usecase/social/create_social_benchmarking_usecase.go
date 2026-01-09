package social

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type CreateSocialBenchmarkingUseCase struct {
	gateway gateway.SocialBenchmarkingGateway
}

type CreateSocialBenchmarkingInput struct {
	BrandName   string
	AvgLikes    float64
	AvgComments float64
	Followers   *int
}

func NewCreateSocialBenchmarking(gateway gateway.SocialBenchmarkingGateway) *CreateSocialBenchmarkingUseCase {
	return &CreateSocialBenchmarkingUseCase{gateway: gateway}
}

func (uc *CreateSocialBenchmarkingUseCase) Execute(ctx context.Context, input CreateSocialBenchmarkingInput) (*entity.SocialBenchmarking, error) {
	// Criar entidade (validação interna)
	benchmarking, err := entity.NewSocialBenchmarking(
		input.BrandName,
		input.AvgLikes,
		input.AvgComments,
		input.Followers,
	)
	if err != nil {
		return nil, err
	}

	// Salvar via gateway
	if err := uc.gateway.Save(ctx, benchmarking); err != nil {
		return nil, err
	}

	return benchmarking, nil
}
