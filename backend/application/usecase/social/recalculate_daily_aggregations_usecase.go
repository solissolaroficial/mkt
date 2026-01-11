package social

import (
	"context"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type RecalculateDailyAggregationsUseCase struct {
	postGateway             gateway.SocialPostGateway
	dailyAggregationGateway gateway.SocialDailyAggregationGateway
	benchmarkingGateway     gateway.SocialBenchmarkingGateway
}

func NewRecalculateDailyAggregationsUseCase(
	postGateway gateway.SocialPostGateway,
	dailyAggregationGateway gateway.SocialDailyAggregationGateway,
	benchmarkingGateway gateway.SocialBenchmarkingGateway,
) *RecalculateDailyAggregationsUseCase {
	return &RecalculateDailyAggregationsUseCase{
		postGateway:             postGateway,
		dailyAggregationGateway: dailyAggregationGateway,
		benchmarkingGateway:     benchmarkingGateway,
	}
}

func (uc *RecalculateDailyAggregationsUseCase) Execute(brandName string, date time.Time) (*entity.SocialDailyAggregation, error) {
	// Buscar todos os posts da marca na data
	posts, err := uc.postGateway.ListByBrandAndDate(brandName, date)
	if err != nil {
		return nil, err
	}

	// Se não houver posts, deletar agregação existente (se houver)
	if len(posts) == 0 {
		// Buscar agregação existente
		existingAggregation, err := uc.dailyAggregationGateway.GetByBrandAndDate(brandName, date)
		if err == nil && existingAggregation != nil {
			// Deletar agregação
			if err := uc.dailyAggregationGateway.Delete(existingAggregation.ID()); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}

	// Calcular agregações
	totalPosts := len(posts)
	totalLikes := 0
	totalComments := 0
	totalShares := 0
	followersAtDate := 0

	for _, post := range posts {
		totalLikes += post.Likes()
		totalComments += post.Comments()
		if post.Shares() != nil {
			totalShares += *post.Shares()
		}
		if post.FollowersAtPost() != nil && *post.FollowersAtPost() > followersAtDate {
			followersAtDate = *post.FollowersAtPost()
		}
	}

	avgLikes := float64(totalLikes) / float64(totalPosts)
	avgComments := float64(totalComments) / float64(totalPosts)
	avgShares := float64(totalShares) / float64(totalPosts)

	// Buscar agregação existente
	existingAggregation, err := uc.dailyAggregationGateway.GetByBrandAndDate(brandName, date)

	var aggregation *entity.SocialDailyAggregation

	if err == nil && existingAggregation != nil {
		// Atualizar agregação existente
		err = existingAggregation.UpdateAggregations(
			totalPosts,
			totalLikes,
			totalComments,
			&totalShares,
			avgLikes,
			avgComments,
			&avgShares,
			&followersAtDate,
		)
		if err != nil {
			return nil, err
		}

		if err := uc.dailyAggregationGateway.Update(existingAggregation); err != nil {
			return nil, err
		}

		aggregation = existingAggregation
	} else {
		// Criar nova agregação
		aggregation, err = entity.NewSocialDailyAggregation(
			brandName,
			date,
			totalPosts,
			totalLikes,
			totalComments,
			&totalShares,
			avgLikes,
			avgComments,
			&avgShares,
			&followersAtDate,
		)
		if err != nil {
			return nil, err
		}

		if err := uc.dailyAggregationGateway.Create(aggregation); err != nil {
			return nil, err
		}
	}

	// Recalcular benchmarking global (média de todos os dias)
	uc.recalculateGlobalBenchmarking(brandName)

	return aggregation, nil
}

func (uc *RecalculateDailyAggregationsUseCase) recalculateGlobalBenchmarking(brandName string) error {
	ctx := context.Background()

	// Buscar todas as agregações da marca
	criteria := domain.NewSocialDailyAggregationCriteria().WithBrandName(brandName)

	// Buscar todas as agregações em lotes (sempre usar paginação)
	allAggregations := make([]*entity.SocialDailyAggregation, 0)
	page := 1
	pageSize := 10000

	for {
		pagination := valueobject.NewPagination(page, pageSize)
		batch, _, err := uc.dailyAggregationGateway.List(criteria, &pagination)
		if err != nil {
			return err
		}
		allAggregations = append(allAggregations, batch...)

		if int64(len(batch)) < int64(pageSize) {
			break
		}
		page++
	}

	// Se não houver agregações, não fazer nada
	if len(allAggregations) == 0 {
		return nil
	}

	// Calcular médias globais com todas as agregações
	totalPosts := 0
	totalLikes := 0
	totalComments := 0
	followers := 0

	for _, agg := range allAggregations {
		totalPosts += agg.TotalPosts()
		totalLikes += agg.TotalLikes()
		totalComments += agg.TotalComments()
		if agg.FollowersAtDate() != nil && *agg.FollowersAtDate() > followers {
			followers = *agg.FollowersAtDate()
		}
	}

	avgLikes := float64(totalLikes) / float64(len(allAggregations))
	avgComments := float64(totalComments) / float64(len(allAggregations))

	// Buscar benchmarking existente
	existingBenchmarking, err := uc.benchmarkingGateway.GetByBrand(brandName)

	var benchmarking *entity.SocialBenchmarking

	if err == nil && existingBenchmarking != nil {
		// Atualizar benchmarking existente
		err = existingBenchmarking.UpdateMetrics(avgLikes, avgComments, &followers)
		if err != nil {
			return err
		}

		if err := uc.benchmarkingGateway.Update(ctx, existingBenchmarking); err != nil {
			return err
		}

		benchmarking = existingBenchmarking
	} else {
		// Criar novo benchmarking
		benchmarking, err = entity.NewSocialBenchmarking(
			brandName,
			avgLikes,
			avgComments,
			&followers,
		)
		if err != nil {
			return err
		}

		if err := uc.benchmarkingGateway.Save(ctx, benchmarking); err != nil {
			return err
		}
	}

	return nil
}
