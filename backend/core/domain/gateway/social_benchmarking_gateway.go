package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type SocialBenchmarkingGateway interface {
	// CRUD
	Save(ctx context.Context, benchmarking *entity.SocialBenchmarking) error
	FindByID(ctx context.Context, id string) (*entity.SocialBenchmarking, error)
	Update(ctx context.Context, benchmarking *entity.SocialBenchmarking) error
	Delete(ctx context.Context, id string) error

	// Queries
	FindByCriteria(
		ctx context.Context,
		criteria *domain.SocialBenchmarkingCriteria,
		pagination *valueobject.Pagination,
		sortOrder *valueobject.SortOrder,
	) ([]*entity.SocialBenchmarking, error)
	CountByCriteria(ctx context.Context, criteria *domain.SocialBenchmarkingCriteria) (int64, error)

	// Utilities
	ExistsByID(ctx context.Context, id string) (bool, error)
	GetByBrandID(brandID uuid.UUID) (*entity.SocialBenchmarking, error)
}
