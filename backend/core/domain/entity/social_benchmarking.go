package entity

import (
	"errors"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/valueobject"

	"github.com/google/uuid"
)

type SocialBenchmarking struct {
	id             uuid.UUID
	brandName      *valueobject.BrandName
	avgLikes       float64
	avgComments    float64
	followers      *int
	engagementRate *valueobject.EngagementRate
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      *time.Time
}

// NewSocialBenchmarking cria uma nova entidade SocialBenchmarking
func NewSocialBenchmarking(
	brandName string,
	avgLikes float64,
	avgComments float64,
	followers *int,
) (*SocialBenchmarking, error) {
	// Validar e criar value object para brand name
	brand, err := valueobject.NewBrandName(brandName)
	if err != nil {
		return nil, err
	}

	// Validar valores numéricos
	if avgLikes < 0 {
		return nil, errors.New("avgLikes cannot be negative")
	}

	if avgComments < 0 {
		return nil, errors.New("avgComments cannot be negative")
	}

	if followers != nil && *followers < 0 {
		return nil, errors.New("followers cannot be negative")
	}

	// Calcular engagement rate
	var engagementRate *valueobject.EngagementRate
	if followers != nil && *followers > 0 {
		engagementRate, err = valueobject.CalculateEngagementRate(avgLikes, avgComments, *followers)
		if err != nil {
			return nil, err
		}
	} else {
		engagementRate = valueobject.ReconstructEngagementRate(0)
	}

	benchmarking := &SocialBenchmarking{
		id:             uuid.New(),
		brandName:      brand,
		avgLikes:       avgLikes,
		avgComments:    avgComments,
		followers:      followers,
		engagementRate: engagementRate,
		createdAt:      time.Now(),
		updatedAt:      time.Now(),
	}

	if err := benchmarking.Validate(); err != nil {
		return nil, err
	}

	return benchmarking, nil
}

// ReconstructSocialBenchmarking reconstrói a entidade do banco de dados
func ReconstructSocialBenchmarking(
	id uuid.UUID,
	brandName string,
	avgLikes float64,
	avgComments float64,
	followers *int,
	engagementRate float64,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *SocialBenchmarking {
	brand := valueobject.ReconstructBrandName(brandName)
	engRate := valueobject.ReconstructEngagementRate(engagementRate)

	return &SocialBenchmarking{
		id:             id,
		brandName:      brand,
		avgLikes:       avgLikes,
		avgComments:    avgComments,
		followers:      followers,
		engagementRate: engRate,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
		deletedAt:      deletedAt,
	}
}

// Getters
func (s *SocialBenchmarking) ID() uuid.UUID                               { return s.id }
func (s *SocialBenchmarking) BrandName() *valueobject.BrandName           { return s.brandName }
func (s *SocialBenchmarking) AvgLikes() float64                           { return s.avgLikes }
func (s *SocialBenchmarking) AvgComments() float64                        { return s.avgComments }
func (s *SocialBenchmarking) Followers() *int                             { return s.followers }
func (s *SocialBenchmarking) EngagementRate() *valueobject.EngagementRate { return s.engagementRate }
func (s *SocialBenchmarking) CreatedAt() time.Time                        { return s.createdAt }
func (s *SocialBenchmarking) UpdatedAt() time.Time                        { return s.updatedAt }
func (s *SocialBenchmarking) DeletedAt() *time.Time                       { return s.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (s *SocialBenchmarking) Validate() error {
	if s.brandName == nil {
		return errors.New("brandName is required")
	}

	if s.avgLikes < 0 {
		return errors.New("avgLikes cannot be negative")
	}

	if s.avgComments < 0 {
		return errors.New("avgComments cannot be negative")
	}

	if s.followers != nil && *s.followers < 0 {
		return errors.New("followers cannot be negative")
	}

	return nil
}

// UpdateBrandName atualiza o nome da marca
func (s *SocialBenchmarking) UpdateBrandName(brandName string) error {
	brand, err := valueobject.NewBrandName(brandName)
	if err != nil {
		return err
	}

	s.brandName = brand
	s.updatedAt = time.Now()
	return nil
}

// UpdateMetrics atualiza as métricas (likes, comments, followers)
func (s *SocialBenchmarking) UpdateMetrics(avgLikes, avgComments float64, followers *int) error {
	// Validar valores numéricos
	if avgLikes < 0 {
		return errors.New("avgLikes cannot be negative")
	}

	if avgComments < 0 {
		return errors.New("avgComments cannot be negative")
	}

	if followers != nil && *followers < 0 {
		return errors.New("followers cannot be negative")
	}

	s.avgLikes = avgLikes
	s.avgComments = avgComments
	s.followers = followers

	// Recalcular engagement rate
	var engagementRate *valueobject.EngagementRate
	var err error
	if followers != nil && *followers > 0 {
		engagementRate, err = valueobject.CalculateEngagementRate(avgLikes, avgComments, *followers)
		if err != nil {
			return err
		}
	} else {
		engagementRate = valueobject.ReconstructEngagementRate(0)
	}

	s.engagementRate = engagementRate
	s.updatedAt = time.Now()
	return nil
}

// SoftDelete marca o benchmarking como deletado
func (s *SocialBenchmarking) SoftDelete() {
	now := time.Now()
	s.deletedAt = &now
	s.updatedAt = now
}

// IsActive verifica se o benchmarking está ativo (não deletado)
func (s *SocialBenchmarking) IsActive() bool {
	return s.deletedAt == nil
}
