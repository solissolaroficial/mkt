package gateway

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type RepMarketingActionGateway interface {
	// CRUD
	Save(ctx context.Context, action *entity.RepMarketingAction) error
	FindByID(ctx context.Context, id string) (*entity.RepMarketingAction, error)
	Update(ctx context.Context, action *entity.RepMarketingAction) error
	Delete(ctx context.Context, id string) error

	// Queries
	FindByCriteria(
		ctx context.Context,
		criteria *domain.RepMarketingActionCriteria,
		pagination *valueobject.Pagination,
		sortOrder *valueobject.SortOrder,
	) ([]*entity.RepMarketingAction, error)
	CountByCriteria(ctx context.Context, criteria *domain.RepMarketingActionCriteria) (int64, error)

	// Utilities
	ExistsByID(ctx context.Context, id string) (bool, error)
}
