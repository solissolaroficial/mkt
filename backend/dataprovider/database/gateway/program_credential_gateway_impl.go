package gateway

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type programCredentialGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.ProgramCredentialMapper
}

func NewProgramCredentialGateway(db *gorm.DB) gateway.ProgramCredentialGateway {
	return &programCredentialGatewayImpl{
		db:     db,
		mapper: mapper.NewProgramCredentialMapper(),
	}
}

func (g *programCredentialGatewayImpl) Save(ctx context.Context, credential *entity.ProgramCredential) error {
	credentialModel := g.mapper.EntityToModel(credential)
	return g.db.WithContext(ctx).Create(credentialModel).Error
}

func (g *programCredentialGatewayImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.ProgramCredential, error) {
	var credentialModel model.ProgramCredentialModel
	err := g.db.WithContext(ctx).
		Where("uuid = ? AND active = ?", id, true).
		First(&credentialModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrProgramCredentialNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&credentialModel), nil
}

func (g *programCredentialGatewayImpl) FindAll(ctx context.Context) ([]*entity.ProgramCredential, error) {
	var credentialModels []model.ProgramCredentialModel
	err := g.db.WithContext(ctx).
		Where("active = ?", true).
		Find(&credentialModels).Error
	if err != nil {
		return nil, err
	}

	return g.mapper.ModelsToEntities(credentialModels), nil
}

func (g *programCredentialGatewayImpl) FindActive(ctx context.Context) ([]*entity.ProgramCredential, error) {
	var credentialModels []model.ProgramCredentialModel
	err := g.db.WithContext(ctx).
		Where("active = ?", true).
		Order("name ASC").
		Find(&credentialModels).Error
	if err != nil {
		return nil, err
	}

	return g.mapper.ModelsToEntities(credentialModels), nil
}

func (g *programCredentialGatewayImpl) Update(ctx context.Context, credential *entity.ProgramCredential) error {
	credentialModel := g.mapper.EntityToModel(credential)
	result := g.db.WithContext(ctx).
		Where("uuid = ? AND active = ?", credentialModel.UUID, true).
		Save(credentialModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrProgramCredentialNotFound
	}
	return nil
}

func (g *programCredentialGatewayImpl) Delete(ctx context.Context, id uuid.UUID) error {
	// Soft delete: set active = false
	result := g.db.WithContext(ctx).
		Model(&model.ProgramCredentialModel{}).
		Where("uuid = ? AND active = ?", id, true).
		Update("active", false)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrProgramCredentialNotFound
	}
	return nil
}

func (g *programCredentialGatewayImpl) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Model(&model.ProgramCredentialModel{}).
		Where("name = ? AND active = ?", name, true).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
