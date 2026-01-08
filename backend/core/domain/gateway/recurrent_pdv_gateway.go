package gateway

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// RecurrentPdvGateway define a interface para operações de banco de dados de PDVs recorrentes
type RecurrentPdvGateway interface {
	// CRUD
	Save(ctx context.Context, pdv *entity.RecurrentPdv) error
	FindByID(ctx context.Context, id string) (*entity.RecurrentPdv, error)
	Update(ctx context.Context, pdv *entity.RecurrentPdv) error
	Delete(ctx context.Context, id string) error

	// Queries
	FindByCriteria(
		ctx context.Context,
		criteria *domain.RecurrentPdvCriteria,
		pagination *valueobject.Pagination,
		sortOrder *valueobject.SortOrder,
	) ([]*entity.RecurrentPdv, error)
	CountByCriteria(ctx context.Context, criteria *domain.RecurrentPdvCriteria) (int64, error)

	// Utilities
	ExistsByID(ctx context.Context, id string) (bool, error)
}
