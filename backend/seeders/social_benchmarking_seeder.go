package seeders

import (
	"context"
	"log"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type SocialBenchmarkingSeeder struct {
	benchmarkingGateway gateway.SocialBenchmarkingGateway
}

func NewSocialBenchmarkingSeeder(benchmarkingGateway gateway.SocialBenchmarkingGateway) *SocialBenchmarkingSeeder {
	return &SocialBenchmarkingSeeder{
		benchmarkingGateway: benchmarkingGateway,
	}
}

func (s *SocialBenchmarkingSeeder) Seed(ctx context.Context) error {
	// Verificar se já foi seeded
	count, _ := s.benchmarkingGateway.CountByCriteria(ctx, domain.NewSocialBenchmarkingCriteria())
	if count > 0 {
		log.Println("Social benchmarkings already seeded, skipping...")
		return nil
	}

	// Criar benchmarkings iniciais (mesmos dados do frontend mock)
	benchmarkings := []struct {
		BrandName   string
		AvgLikes    float64
		AvgComments float64
		Followers   int
	}{
		{"Solis Solar", 145.5, 12.3, 15400},
		{"Competitor A", 120.2, 8.5, 12000},
		{"Competitor B", 98.4, 5.2, 8500},
		{"Competitor C", 210.8, 18.9, 25000},
	}

	for _, b := range benchmarkings {
		benchmarking, err := entity.NewSocialBenchmarking(
			b.BrandName,
			b.AvgLikes,
			b.AvgComments,
			&b.Followers,
		)
		if err != nil {
			log.Printf("Error creating benchmarking %s: %v", b.BrandName, err)
			continue
		}

		if err := s.benchmarkingGateway.Save(ctx, benchmarking); err != nil {
			log.Printf("Error saving benchmarking %s: %v", b.BrandName, err)
			continue
		}

		log.Printf("Created benchmarking: %s", b.BrandName)
	}

	log.Println("Social benchmarkings seeded successfully")
	return nil
}
