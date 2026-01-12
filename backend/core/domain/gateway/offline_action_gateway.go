package gateway

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type OfflineActionGateway interface {
	// CRUD
	Save(ctx context.Context, action *entity.OfflineAction) error
	FindByID(ctx context.Context, id string) (*entity.OfflineAction, error)
	Update(ctx context.Context, action *entity.OfflineAction) error
	Delete(ctx context.Context, id string) error

	// Queries
	FindByCriteria(
		ctx context.Context,
		criteria *domain.OfflineActionCriteria,
		pagination *valueobject.Pagination,
		sortOrder *valueobject.SortOrder,
	) ([]*entity.OfflineAction, error)
	CountByCriteria(ctx context.Context, criteria *domain.OfflineActionCriteria) (int64, error)

	// Utilities
	ExistsByID(ctx context.Context, id string) (bool, error)
}
