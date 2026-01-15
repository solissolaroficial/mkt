package representatives

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type ListRepresentativesInput struct {
	Page      int
	PageSize  int
	Name      *string
	Company   *string
	Email     *string
	Region    *string
	City      *string
	Active    *bool
	Code      *int
	SortBy    string
	SortOrder string
}

type ListRepresentativesOutput struct {
	Data       []*entity.Representative
	Total      int64
	Page       int
	PageSize   int
	TotalPages int
}

type ListRepresentativesUseCase struct {
	representativeGateway gateway.RepresentativeGateway
}

func NewListRepresentativesUseCase(representativeGateway gateway.RepresentativeGateway) *ListRepresentativesUseCase {
	return &ListRepresentativesUseCase{
		representativeGateway: representativeGateway,
	}
}

func (uc *ListRepresentativesUseCase) Execute(ctx context.Context, input ListRepresentativesInput) (*ListRepresentativesOutput, error) {
	// Validate pagination parameters
	if input.Page < 1 {
		input.Page = 1
	}
	if input.PageSize < 1 {
		input.PageSize = 10
	}
	if input.PageSize > 100 {
		input.PageSize = 100
	}

	// Create pagination value object
	pagination := valueobject.NewPagination(input.Page, input.PageSize)

	// Create sort order value object
	sortOrder, err := valueobject.NewSortOrder(input.SortBy, valueobject.SortDirection(input.SortOrder))
	if err != nil {
		return nil, err
	}
	sortOrders := []*valueobject.SortOrder{sortOrder}

	// Build criteria
	criteria := domain.NewRepresentativeCriteria()
	if input.Name != nil {
		criteria = criteria.WithName(*input.Name)
	}
	if input.Company != nil {
		criteria = criteria.WithCompany(*input.Company)
	}
	if input.Email != nil {
		criteria = criteria.WithEmail(*input.Email)
	}
	if input.Region != nil {
		criteria = criteria.WithRegion(*input.Region)
	}
	if input.City != nil {
		criteria = criteria.WithCity(*input.City)
	}
	if input.Active != nil {
		criteria = criteria.WithActive(*input.Active)
	}
	if input.Code != nil {
		criteria = criteria.WithCode(*input.Code)
	}

	// Validate criteria
	if err := criteria.Validate(); err != nil {
		return nil, err
	}

	// Query representatives
	representatives, total, err := uc.representativeGateway.FindByCriteria(criteria, &pagination, sortOrders)
	if err != nil {
		return nil, err
	}

	// Calculate total pages
	totalPages := int(total) / input.PageSize
	if int(total)%input.PageSize > 0 {
		totalPages++
	}

	return &ListRepresentativesOutput{
		Data:       representatives,
		Total:      total,
		Page:       input.Page,
		PageSize:   input.PageSize,
		TotalPages: totalPages,
	}, nil
}
