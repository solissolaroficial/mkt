package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type SocialDailyAggregation struct {
	id              uuid.UUID
	brandID         uuid.UUID
	brandName       *valueobject.BrandName
	aggregationDate time.Time
	totalPosts      int
	totalLikes      int
	totalComments   int
	totalShares     *int
	avgLikes        float64
	avgComments     float64
	avgShares       *float64
	followersAtDate *int
	engagementRate  *valueobject.EngagementRate
	createdAt       time.Time
	updatedAt       time.Time
	deletedAt       *time.Time
}

// NewSocialDailyAggregation cria uma nova entidade SocialDailyAggregation
func NewSocialDailyAggregation(
	brandName string,
	aggregationDate time.Time,
	totalPosts int,
	totalLikes int,
	totalComments int,
	totalShares *int,
	avgLikes float64,
	avgComments float64,
	avgShares *float64,
	followersAtDate *int,
) (*SocialDailyAggregation, error) {
	brand, err := valueobject.NewBrandName(brandName)
	if err != nil {
		return nil, err
	}

	if totalPosts < 0 {
		return nil, errors.New("totalPosts cannot be negative")
	}

	if totalLikes < 0 {
		return nil, errors.New("totalLikes cannot be negative")
	}

	if totalComments < 0 {
		return nil, errors.New("totalComments cannot be negative")
	}

	if totalShares != nil && *totalShares < 0 {
		return nil, errors.New("totalShares cannot be negative")
	}

	if avgLikes < 0 {
		return nil, errors.New("avgLikes cannot be negative")
	}

	if avgComments < 0 {
		return nil, errors.New("avgComments cannot be negative")
	}

	if avgShares != nil && *avgShares < 0 {
		return nil, errors.New("avgShares cannot be negative")
	}

	if followersAtDate != nil && *followersAtDate < 0 {
		return nil, errors.New("followersAtDate cannot be negative")
	}

	// Calcular engagement rate
	var engagementRate *valueobject.EngagementRate
	var errRate error
	if followersAtDate != nil && *followersAtDate > 0 {
		engagementRate, errRate = valueobject.CalculateEngagementRate(avgLikes, avgComments, *followersAtDate)
		if errRate != nil {
			return nil, errRate
		}
	} else {
		engagementRate = valueobject.ReconstructEngagementRate(0)
	}

	aggregation := &SocialDailyAggregation{
		id:              uuid.New(),
		brandID:         uuid.New(), // Will be updated when brand is created/fetched
		brandName:       brand,
		aggregationDate: aggregationDate,
		totalPosts:      totalPosts,
		totalLikes:      totalLikes,
		totalComments:   totalComments,
		totalShares:     totalShares,
		avgLikes:        avgLikes,
		avgComments:     avgComments,
		avgShares:       avgShares,
		followersAtDate: followersAtDate,
		engagementRate:  engagementRate,
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
		deletedAt:       nil,
	}

	if err := aggregation.Validate(); err != nil {
		return nil, err
	}

	return aggregation, nil
}

// ReconstructSocialDailyAggregation reconstrói a entidade do banco de dados
func ReconstructSocialDailyAggregation(
	id uuid.UUID,
	brandID uuid.UUID,
	brandName string,
	aggregationDate time.Time,
	totalPosts int,
	totalLikes int,
	totalComments int,
	totalShares *int,
	avgLikes float64,
	avgComments float64,
	avgShares *float64,
	followersAtDate *int,
	engagementRate float64,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *SocialDailyAggregation {
	brand := valueobject.ReconstructBrandName(brandName)
	engRate := valueobject.ReconstructEngagementRate(engagementRate)

	return &SocialDailyAggregation{
		id:              id,
		brandID:         brandID,
		brandName:       brand,
		aggregationDate: aggregationDate,
		totalPosts:      totalPosts,
		totalLikes:      totalLikes,
		totalComments:   totalComments,
		totalShares:     totalShares,
		avgLikes:        avgLikes,
		avgComments:     avgComments,
		avgShares:       avgShares,
		followersAtDate: followersAtDate,
		engagementRate:  engRate,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
		deletedAt:       deletedAt,
	}
}

// Getters
func (s *SocialDailyAggregation) ID() uuid.UUID                     { return s.id }
func (s *SocialDailyAggregation) BrandID() uuid.UUID                { return s.brandID }
func (s *SocialDailyAggregation) BrandName() *valueobject.BrandName { return s.brandName }
func (s *SocialDailyAggregation) AggregationDate() time.Time        { return s.aggregationDate }
func (s *SocialDailyAggregation) TotalPosts() int                   { return s.totalPosts }
func (s *SocialDailyAggregation) TotalLikes() int                   { return s.totalLikes }
func (s *SocialDailyAggregation) TotalComments() int                { return s.totalComments }
func (s *SocialDailyAggregation) TotalShares() *int                 { return s.totalShares }
func (s *SocialDailyAggregation) AvgLikes() float64                 { return s.avgLikes }
func (s *SocialDailyAggregation) AvgComments() float64              { return s.avgComments }
func (s *SocialDailyAggregation) AvgShares() *float64               { return s.avgShares }
func (s *SocialDailyAggregation) FollowersAtDate() *int             { return s.followersAtDate }
func (s *SocialDailyAggregation) EngagementRate() *valueobject.EngagementRate {
	return s.engagementRate
}
func (s *SocialDailyAggregation) CreatedAt() time.Time  { return s.createdAt }
func (s *SocialDailyAggregation) UpdatedAt() time.Time  { return s.updatedAt }
func (s *SocialDailyAggregation) DeletedAt() *time.Time { return s.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (s *SocialDailyAggregation) Validate() error {
	if s.brandID == uuid.Nil {
		return errors.New("brandID is required")
	}

	if s.brandName == nil {
		return errors.New("brandName is required")
	}

	if s.aggregationDate.IsZero() {
		return errors.New("aggregationDate is required")
	}

	if s.totalPosts < 0 {
		return errors.New("totalPosts cannot be negative")
	}

	if s.totalLikes < 0 {
		return errors.New("totalLikes cannot be negative")
	}

	if s.totalComments < 0 {
		return errors.New("totalComments cannot be negative")
	}

	if s.totalShares != nil && *s.totalShares < 0 {
		return errors.New("totalShares cannot be negative")
	}

	if s.avgLikes < 0 {
		return errors.New("avgLikes cannot be negative")
	}

	if s.avgComments < 0 {
		return errors.New("avgComments cannot be negative")
	}

	if s.avgShares != nil && *s.avgShares < 0 {
		return errors.New("avgShares cannot be negative")
	}

	if s.followersAtDate != nil && *s.followersAtDate < 0 {
		return errors.New("followersAtDate cannot be negative")
	}

	return nil
}

// SoftDelete marca a agregação como deletada
func (s *SocialDailyAggregation) SoftDelete() {
	now := time.Now()
	s.deletedAt = &now
	s.updatedAt = now
}

// IsActive verifica se a agregação está ativa (não deletada)
func (s *SocialDailyAggregation) IsActive() bool {
	return s.deletedAt == nil
}

// UpdateAggregations atualiza as agregações
func (s *SocialDailyAggregation) UpdateAggregations(
	totalPosts int,
	totalLikes int,
	totalComments int,
	totalShares *int,
	avgLikes float64,
	avgComments float64,
	avgShares *float64,
	followersAtDate *int,
) error {
	if totalPosts < 0 {
		return errors.New("totalPosts cannot be negative")
	}

	if totalLikes < 0 {
		return errors.New("totalLikes cannot be negative")
	}

	if totalComments < 0 {
		return errors.New("totalComments cannot be negative")
	}

	if totalShares != nil && *totalShares < 0 {
		return errors.New("totalShares cannot be negative")
	}

	if avgLikes < 0 {
		return errors.New("avgLikes cannot be negative")
	}

	if avgComments < 0 {
		return errors.New("avgComments cannot be negative")
	}

	if avgShares != nil && *avgShares < 0 {
		return errors.New("avgShares cannot be negative")
	}

	if followersAtDate != nil && *followersAtDate < 0 {
		return errors.New("followersAtDate cannot be negative")
	}

	s.totalPosts = totalPosts
	s.totalLikes = totalLikes
	s.totalComments = totalComments
	s.totalShares = totalShares
	s.avgLikes = avgLikes
	s.avgComments = avgComments
	s.avgShares = avgShares
	s.followersAtDate = followersAtDate

	// Recalcular engagement rate
	var engagementRate *valueobject.EngagementRate
	var err error
	if followersAtDate != nil && *followersAtDate > 0 {
		engagementRate, err = valueobject.CalculateEngagementRate(avgLikes, avgComments, *followersAtDate)
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

// UpdateBrandID atualiza o ID da marca
func (s *SocialDailyAggregation) UpdateBrandID(brandID uuid.UUID) error {
	if brandID == uuid.Nil {
		return errors.New("brandID cannot be nil")
	}

	s.brandID = brandID
	s.updatedAt = time.Now()
	return nil
}
