package response

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

type BrandResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BrandsListResponse struct {
	Brands []BrandResponse `json:"brands"`
}

type BrandPayloadMapper struct{}

func NewBrandPayloadMapper() *BrandPayloadMapper {
	return &BrandPayloadMapper{}
}

func (m *BrandPayloadMapper) ToBrandResponse(brand *entity.Brand) BrandResponse {
	return BrandResponse{
		ID:        brand.UUID().String(),
		Name:      brand.Name(),
		CreatedAt: brand.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: brand.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}
}

func (m *BrandPayloadMapper) ToBrandsListResponse(brands []*entity.Brand) BrandsListResponse {
	brandResponses := make([]BrandResponse, len(brands))
	for i, brand := range brands {
		brandResponses[i] = m.ToBrandResponse(brand)
	}
	return BrandsListResponse{Brands: brandResponses}
}
