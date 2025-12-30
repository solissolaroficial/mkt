package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// KpiGateway defines the interface for KPI category data operations
type KpiGateway interface {
	// Save creates a new KPI category or updates an existing one
	Save(ctx context.Context, kpi *entity.KpiCategory) error

	// FindByID retrieves a KPI category by its ID
	FindByID(ctx context.Context, id uuid.UUID) (*entity.KpiCategory, error)

	// FindByTitle retrieves a KPI category by its title
	FindByTitle(ctx context.Context, title string) (*entity.KpiCategory, error)

	// FindBySlug retrieves a KPI category by its slug
	FindBySlug(ctx context.Context, slug string) (*entity.KpiCategory, error)

	// FindBySlugs retrieves KPI categories by a list of slugs
	FindBySlugs(ctx context.Context, slugs []string) ([]*entity.KpiCategory, error)

	// FindAll retrieves KPI categories with pagination
	FindAll(ctx context.Context, pagination valueobject.Pagination) ([]*entity.KpiCategory, int64, error)

	// Update modifies an existing KPI category
	Update(ctx context.Context, kpi *entity.KpiCategory) error

	// Delete removes a KPI category by its ID
	Delete(ctx context.Context, id uuid.UUID) error
}
