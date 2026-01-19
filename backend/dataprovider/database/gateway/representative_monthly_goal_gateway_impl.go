package gateway

import (
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	representativemonthlygoalerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type representativeMonthlyGoalGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.RepresentativeMonthlyGoalMapper
}

// NewRepresentativeMonthlyGoalGateway creates a new RepresentativeMonthlyGoalGateway implementation
func NewRepresentativeMonthlyGoalGateway(db *gorm.DB) gateway.RepresentativeMonthlyGoalGateway {
	return &representativeMonthlyGoalGatewayImpl{
		db:     db,
		mapper: mapper.NewRepresentativeMonthlyGoalMapper(),
	}
}

// Create creates a new representative monthly goal
func (g *representativeMonthlyGoalGatewayImpl) Create(goal *entity.RepresentativeMonthlyGoal) error {
	goalModel := g.mapper.EntityToModel(goal)

	if err := g.db.Create(goalModel).Error; err != nil {
		return err
	}

	return nil
}

// Update updates an existing representative monthly goal
func (g *representativeMonthlyGoalGatewayImpl) Update(goal *entity.RepresentativeMonthlyGoal) error {
	goalModel := g.mapper.EntityToModel(goal)

	result := g.db.Where("id = ?", goalModel.ID).Save(goalModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return representativemonthlygoalerrors.ErrRepresentativeGoalNotFound
	}

	return nil
}

// GetByID retrieves a representative monthly goal by ID
func (g *representativeMonthlyGoalGatewayImpl) GetByID(id uuid.UUID) (*entity.RepresentativeMonthlyGoal, error) {
	var goalModel model.RepresentativeMonthlyGoalModel

	err := g.db.Where("id = ?", id).First(&goalModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, representativemonthlygoalerrors.ErrRepresentativeGoalNotFound
		}
		return nil, err
	}

	return g.mapper.ModelToEntity(&goalModel)
}

// List retrieves representative monthly goals based on criteria
func (g *representativeMonthlyGoalGatewayImpl) List(criteria *domain.RepresentativeMonthlyGoalCriteria) ([]*entity.RepresentativeMonthlyGoal, int64, error) {
	var goalModels []model.RepresentativeMonthlyGoalModel
	var total int64

	// Build query
	query := g.db.Model(&model.RepresentativeMonthlyGoalModel{})

	// Apply filters
	if criteria.RepresentativeID != nil {
		query = query.Where("representative_uuid = ?", *criteria.RepresentativeID)
	}

	if criteria.Month != nil {
		query = query.Where("month = ?", *criteria.Month)
	}

	if criteria.Year != nil {
		query = query.Where("year = ?", *criteria.Year)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	if criteria.SortOrder != nil {
		query = query.Order(criteria.SortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(criteria.Pagination.Offset()).
		Limit(criteria.Pagination.Limit()).
		Find(&goalModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	goalModelPointers := make([]*model.RepresentativeMonthlyGoalModel, len(goalModels))
	for i := range goalModels {
		goalModelPointers[i] = &goalModels[i]
	}

	entities, err := g.mapper.ModelsToEntities(goalModelPointers)
	return entities, total, err
}

// Delete soft deletes a representative monthly goal
func (g *representativeMonthlyGoalGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.RepresentativeMonthlyGoalModel{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return representativemonthlygoalerrors.ErrRepresentativeGoalNotFound
	}

	return nil
}

// GetByRepresentativeAndMonth retrieves a goal for a specific representative and month/year
func (g *representativeMonthlyGoalGatewayImpl) GetByRepresentativeAndMonth(representativeID uuid.UUID, month int, year int) (*entity.RepresentativeMonthlyGoal, error) {
	var goalModel model.RepresentativeMonthlyGoalModel

	err := g.db.Where("representative_uuid = ? AND month = ? AND year = ?", representativeID, month, year).First(&goalModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, representativemonthlygoalerrors.ErrRepresentativeGoalNotFound
		}
		return nil, err
	}

	return g.mapper.ModelToEntity(&goalModel)
}

// GetGoalsByRepresentative retrieves all goals for a specific representative
func (g *representativeMonthlyGoalGatewayImpl) GetGoalsByRepresentative(representativeID uuid.UUID) ([]*entity.RepresentativeMonthlyGoal, error) {
	var goalModels []model.RepresentativeMonthlyGoalModel

	err := g.db.Where("representative_uuid = ?", representativeID).Order("year DESC, month DESC").Find(&goalModels).Error

	if err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	goalModelPointers := make([]*model.RepresentativeMonthlyGoalModel, len(goalModels))
	for i := range goalModels {
		goalModelPointers[i] = &goalModels[i]
	}

	return g.mapper.ModelsToEntities(goalModelPointers)
}

// GetGoalsTableData retrieves table data for all representatives with their monthly goals
// This returns data in a transposed format matching the old implementation
func (g *representativeMonthlyGoalGatewayImpl) GetGoalsTableData(year int, month *int) ([]map[string]interface{}, error) {
	type GoalRow struct {
		RepresentativeUUID uuid.UUID
		RepresentativeName string
		Company            string
		Region             string
		City               string
		Month              int
		Target             float64
		Realized           float64
	}

	query := `
		SELECT 
			r.uuid as representative_uuid,
			r.name as representative_name,
			r.company,
			r.region,
			r.city,
			g.month,
			g.target,
			g.realized
		FROM representatives r
		LEFT JOIN representative_monthly_goals g ON r.uuid = g.representative_uuid
		WHERE r.deleted_at IS NULL
	`

	args := []interface{}{year}

	if month != nil {
		query += " AND g.year = ? AND g.month = ?"
		args = append(args, *month)
	} else {
		query += " AND g.year = ?"
	}

	query += " ORDER BY r.name ASC, g.month ASC"

	var rows []GoalRow
	if err := g.db.Raw(query, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}

	// Convert to map format for easy consumption
	result := make([]map[string]interface{}, 0, len(rows))
	for _, row := range rows {
		rowMap := map[string]interface{}{
			"representative_uuid": row.RepresentativeUUID.String(),
			"representative_name": row.RepresentativeName,
			"company":             row.Company,
			"region":              row.Region,
			"city":                row.City,
			"month":               row.Month,
			"target":              row.Target,
			"realized":            row.Realized,
		}
		result = append(result, rowMap)
	}

	return result, nil
}
