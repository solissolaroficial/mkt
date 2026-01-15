package gateway

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// RepresentativeGateway defines the interface for representative persistence
type RepresentativeGateway interface {
	// Create creates a new representative
	Create(representative *entity.Representative) error

	// Update updates an existing representative
	Update(representative *entity.Representative) error

	// Delete deletes a representative (soft delete)
	Delete(id uuid.UUID) error

	// FindByID finds a representative by UUID
	FindByID(id uuid.UUID) (*entity.Representative, error)

	// FindAll finds all representatives with pagination
	FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error)

	// FindByCriteria finds representatives using the Criteria pattern
	FindByCriteria(criteria *domain.RepresentativeCriteria, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error)

	// CountByCriteria counts representatives using the Criteria pattern
	CountByCriteria(criteria *domain.RepresentativeCriteria) (int64, error)

	// ExistsByCode checks if a representative with the given code exists
	ExistsByCode(code int) (bool, error)

	// ExistsByID checks if a representative with the given UUID exists
	ExistsByID(id uuid.UUID) (bool, error)

	// FindActive finds all active representatives
	FindActive(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error)

	// FindByRegion finds representatives by region
	FindByRegion(region string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error)

	// FindByCompany finds representatives by company
	FindByCompany(company string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error)
}
