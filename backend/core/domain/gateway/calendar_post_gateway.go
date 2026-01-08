package gateway

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type CalendarPostGateway interface {
	// CRUD
	Save(ctx context.Context, post *entity.CalendarPost) error
	FindByID(ctx context.Context, id string) (*entity.CalendarPost, error)
	Update(ctx context.Context, post *entity.CalendarPost) error
	Delete(ctx context.Context, id string) error

	// Queries
	FindByCriteria(
		ctx context.Context,
		criteria *domain.CalendarPostCriteria,
		pagination *valueobject.Pagination,
		sortOrder *valueobject.SortOrder,
	) ([]*entity.CalendarPost, error)
	CountByCriteria(ctx context.Context, criteria *domain.CalendarPostCriteria) (int64, error)

	// Utilities
	ExistsByID(ctx context.Context, id string) (bool, error)
}
