package valueobject

import (
	"errors"
)

var (
	ErrInvalidOfflineCategory = errors.New("invalid offline category")
)

type OfflineCategory string

const (
	CategoryPartnership       OfflineCategory = "PARCERIA"
	CategoryCooperativeAction OfflineCategory = "AÇÃO COOPERADA"
	CategoryExclusiveGifts    OfflineCategory = "ENTREGA DE BRINDES EXCLUSIVOS"
	CategoryMiniatures        OfflineCategory = "MINIATURAS"
	CategoryFairGifts         OfflineCategory = "BRINDES - FEIRA"
	CategoryFairExhibition    OfflineCategory = "FEIRA - EXPOSIÇÃO"
)

func NewOfflineCategory(category string) (OfflineCategory, error) {
	switch OfflineCategory(category) {
	case CategoryPartnership, CategoryCooperativeAction, CategoryExclusiveGifts,
		CategoryMiniatures, CategoryFairGifts, CategoryFairExhibition:
		return OfflineCategory(category), nil
	default:
		return "", ErrInvalidOfflineCategory
	}
}

func (c OfflineCategory) String() string {
	return string(c)
}

func (c OfflineCategory) IsValid() bool {
	switch c {
	case CategoryPartnership, CategoryCooperativeAction, CategoryExclusiveGifts,
		CategoryMiniatures, CategoryFairGifts, CategoryFairExhibition:
		return true
	default:
		return false
	}
}

// GetValidCategories retorna todas as categorias válidas
func GetValidOfflineCategories() []OfflineCategory {
	return []OfflineCategory{
		CategoryPartnership,
		CategoryCooperativeAction,
		CategoryExclusiveGifts,
		CategoryMiniatures,
		CategoryFairGifts,
		CategoryFairExhibition,
	}
}
