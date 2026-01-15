package gateway

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// RepresentativeMonthlyGoalGateway defines interface for representative monthly goal operations
type RepresentativeMonthlyGoalGateway interface {
	// Create creates a new representative monthly goal
	Create(goal *entity.RepresentativeMonthlyGoal) error

	// Update updates an existing representative monthly goal
	Update(goal *entity.RepresentativeMonthlyGoal) error

	// GetByID retrieves a representative monthly goal by ID
	GetByID(id uuid.UUID) (*entity.RepresentativeMonthlyGoal, error)

	// List retrieves representative monthly goals based on criteria
	List(criteria *domain.RepresentativeMonthlyGoalCriteria) ([]*entity.RepresentativeMonthlyGoal, int64, error)

	// Delete soft deletes a representative monthly goal
	Delete(id uuid.UUID) error

	// GetByRepresentativeAndMonth retrieves a goal for a specific representative and month/year
	GetByRepresentativeAndMonth(representativeID uuid.UUID, month int, year int) (*entity.RepresentativeMonthlyGoal, error)

	// GetGoalsByRepresentative retrieves all goals for a specific representative
	GetGoalsByRepresentative(representativeID uuid.UUID) ([]*entity.RepresentativeMonthlyGoal, error)

	// GetGoalsTableData retrieves table data for all representatives with their monthly goals
	GetGoalsTableData(year int, month *int) ([]map[string]interface{}, error)
}
