package calendar

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/repository/criteria"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type ListCalendarPostsUseCase struct {
	gateway gateway.CalendarPostGateway
}

type ListCalendarPostsInput struct {
	Category   *string
	Type       *string
	Status     *string
	AssigneeID *uuid.UUID
	StartDate  *string
	EndDate    *string
	Platform   *string
	Page       int
	Limit      int
	SortBy     *string
	SortOrder  *string
}

func NewListCalendarPosts(gateway gateway.CalendarPostGateway) *ListCalendarPostsUseCase {
	return &ListCalendarPostsUseCase{gateway: gateway}
}

func (uc *ListCalendarPostsUseCase) Execute(ctx context.Context, input ListCalendarPostsInput) ([]*entity.CalendarPost, int64, error) {
	// Criar criteria
	crit := criteria.NewCalendarPostCriteria()

	if input.Category != nil {
		category, err := valueobject.NewPostCategory(*input.Category)
		if err == nil {
			crit = crit.WithCategory(&category)
		}
	}

	if input.Type != nil {
		postType, err := valueobject.NewPostType(*input.Type)
		if err == nil {
			crit = crit.WithType(&postType)
		}
	}

	if input.Status != nil {
		status, err := valueobject.NewPostStatus(*input.Status)
		if err == nil {
			crit = crit.WithStatus(&status)
		}
	}

	if input.AssigneeID != nil {
		crit = crit.WithAssigneeID(input.AssigneeID)
	}

	if input.StartDate != nil {
		crit = crit.WithStartDate(input.StartDate)
	}

	if input.EndDate != nil {
		crit = crit.WithEndDate(input.EndDate)
	}

	if input.Platform != nil {
		crit = crit.WithPlatform(input.Platform)
	}

	// Criar paginação
	pagination := valueobject.NewPagination(input.Page, input.Limit)

	// Criar ordenação
	var sortBy string
	if input.SortBy != nil {
		sortBy = *input.SortBy
	}
	var sortDirection valueobject.SortDirection
	if input.SortOrder != nil {
		sortDirection = valueobject.SortDirection(strings.ToUpper(*input.SortOrder))
	}
	sortOrder, err := valueobject.NewSortOrder(sortBy, sortDirection)
	if err != nil {
		// Se houver erro, usar ordenação padrão
		sortOrder, _ = valueobject.NewSortOrder("date", valueobject.SortDirectionAsc)
	}

	// Buscar posts
	posts, err := uc.gateway.FindByCriteria(ctx, crit, &pagination, sortOrder)
	if err != nil {
		return nil, 0, err
	}

	// Contar total
	total, err := uc.gateway.CountByCriteria(ctx, crit)
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}
