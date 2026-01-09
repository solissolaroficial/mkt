package valueobject

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrInvalidEngagementRate  = errors.New("invalid engagement rate")
	ErrEngagementRateNegative = errors.New("engagement rate cannot be negative")
)

type EngagementRate struct {
	rate float64
}

// CalculateEngagementRate calcula a taxa de engajamento
// Fórmula: ((Likes + Comments) / Followers) * 100
func CalculateEngagementRate(avgLikes, avgComments float64, followers int) (*EngagementRate, error) {
	if followers <= 0 {
		// Se não há seguidores, retorna 0
		return &EngagementRate{rate: 0}, nil
	}

	rate := ((avgLikes + avgComments) / float64(followers)) * 100

	// Verificar se o resultado é válido (não é Infinity ou NaN)
	if math.IsInf(rate, 0) || math.IsNaN(rate) {
		return nil, ErrInvalidEngagementRate
	}

	if rate < 0 {
		return nil, ErrEngagementRateNegative
	}

	// Arredondar para 2 casas decimais
	rounded := math.Round(rate*100) / 100

	return &EngagementRate{rate: rounded}, nil
}

func ReconstructEngagementRate(rate float64) *EngagementRate {
	// Assume que dados do banco são válidos
	return &EngagementRate{rate: rate}
}

func (e *EngagementRate) Value() float64 {
	return e.rate
}

func (e *EngagementRate) String() string {
	return fmt.Sprintf("%.2f", e.rate)
}
