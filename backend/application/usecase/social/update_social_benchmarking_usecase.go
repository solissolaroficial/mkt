package social

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type UpdateSocialBenchmarkingUseCase struct {
	gateway gateway.SocialBenchmarkingGateway
}

type UpdateSocialBenchmarkingInput struct {
	ID          string
	BrandID     *uuid.UUID
	AvgLikes    *float64
	AvgComments *float64
	Followers   *int
}

func NewUpdateSocialBenchmarking(
	gateway gateway.SocialBenchmarkingGateway,
) *UpdateSocialBenchmarkingUseCase {
	return &UpdateSocialBenchmarkingUseCase{
		gateway: gateway,
	}
}

func (uc *UpdateSocialBenchmarkingUseCase) Execute(ctx context.Context, input UpdateSocialBenchmarkingInput) (*entity.SocialBenchmarking, error) {
	// Buscar benchmarking existente
	benchmarking, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	// Atualizar campos se fornecidos
	if input.BrandID != nil {
		if err := benchmarking.UpdateBrandID(*input.BrandID); err != nil {
			return nil, err
		}
	}

	if input.AvgLikes != nil || input.AvgComments != nil || input.Followers != nil {
		avgLikes := benchmarking.AvgLikes()
		if input.AvgLikes != nil {
			avgLikes = *input.AvgLikes
		}

		avgComments := benchmarking.AvgComments()
		if input.AvgComments != nil {
			avgComments = *input.AvgComments
		}

		followers := benchmarking.Followers()
		if input.Followers != nil {
			followers = input.Followers
		}

		if err := benchmarking.UpdateMetrics(avgLikes, avgComments, followers); err != nil {
			return nil, err
		}
	}

	// Salvar no banco de dados
	if err := uc.gateway.Update(ctx, benchmarking); err != nil {
		return nil, err
	}

	return benchmarking, nil
}
