package gateway

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// PdvPostGateway define a interface para operações de banco de dados de posts de PDV
// Nota: Criteria está no package domain, não em repository/criteria
type PdvPostGateway interface {
	// CRUD
	Save(ctx context.Context, post *entity.PdvPost) error
	FindByID(ctx context.Context, id string) (*entity.PdvPost, error)
	Update(ctx context.Context, post *entity.PdvPost) error
	Delete(ctx context.Context, id string) error

	// Queries
	FindByCriteria(
		ctx context.Context,
		criteria *domain.PdvPostCriteria,
		pagination *valueobject.Pagination,
		sortOrder *valueobject.SortOrder,
	) ([]*entity.PdvPost, error)
	CountByCriteria(ctx context.Context, criteria *domain.PdvPostCriteria) (int64, error)

	// Utilities
	ExistsByID(ctx context.Context, id string) (bool, error)
}
