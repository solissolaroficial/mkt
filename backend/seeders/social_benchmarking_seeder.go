package seeders

import (
	"context"
	"log"

	"github.com/google/uuid"
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

	// Criar benchmarkings iniciais (usando UUIDs aleatórios para brandID)
	// NOTA: Em produção, estes brandIDs devem corresponder a marcas reais no banco
	benchmarkings := []struct {
		BrandID     uuid.UUID
		AvgLikes    float64
		AvgComments float64
		Followers   int
	}{
		{uuid.MustParse("00000000-0000-0000-0000-000000000001"), 145.5, 12.3, 15400},
		{uuid.MustParse("00000000-0000-0000-0000-000000000002"), 120.2, 8.5, 12000},
		{uuid.MustParse("00000000-0000-0000-0000-000000000003"), 98.4, 5.2, 8500},
		{uuid.MustParse("00000000-0000-0000-0000-000000000004"), 210.8, 18.9, 25000},
	}

	for _, b := range benchmarkings {
		benchmarking, err := entity.NewSocialBenchmarking(
			b.BrandID,
			b.AvgLikes,
			b.AvgComments,
			&b.Followers,
		)
		if err != nil {
			log.Printf("Error creating benchmarking %s: %v", b.BrandID, err)
			continue
		}

		if err := s.benchmarkingGateway.Save(ctx, benchmarking); err != nil {
			log.Printf("Error saving benchmarking %s: %v", b.BrandID, err)
			continue
		}

		log.Printf("Created benchmarking: %s", b.BrandID)
	}

	log.Println("Social benchmarkings seeded successfully")
	return nil
}
