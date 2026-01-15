package gateway

import (
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	representativeerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type representativeGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.RepresentativeMapper
}

// NewRepresentativeGateway creates a new RepresentativeGateway implementation
func NewRepresentativeGateway(db *gorm.DB) gateway.RepresentativeGateway {
	return &representativeGatewayImpl{
		db:     db,
		mapper: mapper.NewRepresentativeMapper(),
	}
}

// Create creates a new representative
func (g *representativeGatewayImpl) Create(representative *entity.Representative) error {
	representativeModel := g.mapper.EntityToModel(representative)

	if err := g.db.Create(representativeModel).Error; err != nil {
		return err
	}

	return nil
}

// Update updates an existing representative
func (g *representativeGatewayImpl) Update(representative *entity.Representative) error {
	representativeModel := g.mapper.EntityToModel(representative)

	result := g.db.Where("uuid = ?", representativeModel.UUID).Save(representativeModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return representativeerrors.ErrRepresentativeNotFound
	}

	return nil
}

// Delete deletes a representative (soft delete)
func (g *representativeGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.RepresentativeModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return representativeerrors.ErrRepresentativeNotFound
	}

	return nil
}

// FindByID finds a representative by UUID
func (g *representativeGatewayImpl) FindByID(id uuid.UUID) (*entity.Representative, error) {
	var representativeModel model.RepresentativeModel

	err := g.db.Where("uuid = ?", id).First(&representativeModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, representativeerrors.ErrRepresentativeNotFound
		}
		return nil, err
	}

	return g.mapper.ModelToEntity(&representativeModel)
}

// FindAll finds all representatives with pagination
func (g *representativeGatewayImpl) FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error) {
	var representativeModels []model.RepresentativeModel
	var total int64

	// Get total count
	if err := g.db.Model(&model.RepresentativeModel{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.RepresentativeModel{})

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&representativeModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	representativeModelPointers := make([]*model.RepresentativeModel, len(representativeModels))
	for i := range representativeModels {
		representativeModelPointers[i] = &representativeModels[i]
	}

	entities, err := g.mapper.ModelsToEntities(representativeModelPointers)
	return entities, total, err
}

// FindByCriteria finds representatives using the Criteria pattern
func (g *representativeGatewayImpl) FindByCriteria(criteria *domain.RepresentativeCriteria, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error) {
	var representativeModels []model.RepresentativeModel
	var total int64

	// Build query
	query := g.db.Model(&model.RepresentativeModel{})

	// Apply filters
	if criteria.GetName() != nil {
		query = query.Where("name ILIKE ?", "%"+*criteria.GetName()+"%")
	}

	if criteria.GetCompany() != nil {
		query = query.Where("company ILIKE ?", "%"+*criteria.GetCompany()+"%")
	}

	if criteria.GetEmail() != nil {
		query = query.Where("email ILIKE ?", "%"+*criteria.GetEmail()+"%")
	}

	if criteria.GetRegion() != nil {
		query = query.Where("region = ?", *criteria.GetRegion())
	}

	if criteria.GetCity() != nil {
		query = query.Where("city ILIKE ?", "%"+*criteria.GetCity()+"%")
	}

	if criteria.GetActive() != nil {
		query = query.Where("active = ?", *criteria.GetActive())
	}

	if criteria.GetCode() != nil {
		query = query.Where("code = ?", *criteria.GetCode())
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&representativeModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	representativeModelPointers := make([]*model.RepresentativeModel, len(representativeModels))
	for i := range representativeModels {
		representativeModelPointers[i] = &representativeModels[i]
	}

	entities, err := g.mapper.ModelsToEntities(representativeModelPointers)
	return entities, total, err
}

// CountByCriteria counts representatives using the Criteria pattern
func (g *representativeGatewayImpl) CountByCriteria(criteria *domain.RepresentativeCriteria) (int64, error) {
	var total int64

	// Build query
	query := g.db.Model(&model.RepresentativeModel{})

	// Apply filters
	if criteria.GetName() != nil {
		query = query.Where("name ILIKE ?", "%"+*criteria.GetName()+"%")
	}

	if criteria.GetCompany() != nil {
		query = query.Where("company ILIKE ?", "%"+*criteria.GetCompany()+"%")
	}

	if criteria.GetEmail() != nil {
		query = query.Where("email ILIKE ?", "%"+*criteria.GetEmail()+"%")
	}

	if criteria.GetRegion() != nil {
		query = query.Where("region = ?", *criteria.GetRegion())
	}

	if criteria.GetCity() != nil {
		query = query.Where("city ILIKE ?", "%"+*criteria.GetCity()+"%")
	}

	if criteria.GetActive() != nil {
		query = query.Where("active = ?", *criteria.GetActive())
	}

	if criteria.GetCode() != nil {
		query = query.Where("code = ?", *criteria.GetCode())
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// ExistsByCode checks if a representative with the given code exists
func (g *representativeGatewayImpl) ExistsByCode(code int) (bool, error) {
	var count int64
	err := g.db.Model(&model.RepresentativeModel{}).Where("code = ?", code).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// ExistsByID checks if a representative with the given UUID exists
func (g *representativeGatewayImpl) ExistsByID(id uuid.UUID) (bool, error) {
	var count int64
	err := g.db.Model(&model.RepresentativeModel{}).Where("uuid = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// FindActive finds all active representatives
func (g *representativeGatewayImpl) FindActive(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error) {
	var representativeModels []model.RepresentativeModel
	var total int64

	// Get total count
	if err := g.db.Model(&model.RepresentativeModel{}).Where("active = ?", true).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.RepresentativeModel{}).Where("active = ?", true)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&representativeModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	representativeModelPointers := make([]*model.RepresentativeModel, len(representativeModels))
	for i := range representativeModels {
		representativeModelPointers[i] = &representativeModels[i]
	}

	entities, err := g.mapper.ModelsToEntities(representativeModelPointers)
	return entities, total, err
}

// FindByRegion finds representatives by region
func (g *representativeGatewayImpl) FindByRegion(region string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error) {
	var representativeModels []model.RepresentativeModel
	var total int64

	// Get total count
	if err := g.db.Model(&model.RepresentativeModel{}).Where("region = ?", region).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.RepresentativeModel{}).Where("region = ?", region)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&representativeModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	representativeModelPointers := make([]*model.RepresentativeModel, len(representativeModels))
	for i := range representativeModels {
		representativeModelPointers[i] = &representativeModels[i]
	}

	entities, err := g.mapper.ModelsToEntities(representativeModelPointers)
	return entities, total, err
}

// FindByCompany finds representatives by company
func (g *representativeGatewayImpl) FindByCompany(company string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Representative, int64, error) {
	var representativeModels []model.RepresentativeModel
	var total int64

	// Get total count
	if err := g.db.Model(&model.RepresentativeModel{}).Where("company = ?", company).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.RepresentativeModel{}).Where("company = ?", company)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&representativeModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	representativeModelPointers := make([]*model.RepresentativeModel, len(representativeModels))
	for i := range representativeModels {
		representativeModelPointers[i] = &representativeModels[i]
	}

	entities, err := g.mapper.ModelsToEntities(representativeModelPointers)
	return entities, total, err
}
