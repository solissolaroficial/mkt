package gateway

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type ShowroomItemGateway interface {
	// CRUD
	Save(ctx context.Context, item *entity.ShowroomItem) error
	FindByID(ctx context.Context, id string) (*entity.ShowroomItem, error)
	Update(ctx context.Context, item *entity.ShowroomItem) error
	Delete(ctx context.Context, id string) error

	// Queries
	FindByCriteria(
		ctx context.Context,
		criteria *domain.ShowroomItemCriteria,
		pagination *valueobject.Pagination,
		sortOrder *valueobject.SortOrder,
	) ([]*entity.ShowroomItem, error)
	CountByCriteria(ctx context.Context, criteria *domain.ShowroomItemCriteria) (int64, error)

	// Utilities
	ExistsByID(ctx context.Context, id string) (bool, error)
}
