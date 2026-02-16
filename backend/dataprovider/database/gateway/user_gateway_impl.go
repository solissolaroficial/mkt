package gateway

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	authErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type userGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.UserMapper
}

// NewUserGateway creates a new UserGateway implementation
func NewUserGateway(db *gorm.DB) gateway.UserGateway {
	return &userGatewayImpl{
		db:     db,
		mapper: &mapper.UserMapper{},
	}
}

// Save creates a new user or updates an existing one
func (g *userGatewayImpl) Save(ctx context.Context, user *entity.User) error {
	userModel := g.mapper.ToModel(user)

	// Check if user already exists to decide between Create and Update
	var existingUser model.User
	err := g.db.WithContext(ctx).Where("uuid = ?", userModel.UUID).First(&existingUser).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new user
		if err := g.db.WithContext(ctx).Create(userModel).Error; err != nil {
			return err
		}
	} else {
		// Update existing user
		if err := g.db.WithContext(ctx).Save(userModel).Error; err != nil {
			return err
		}
	}

	return nil
}

// FindByID retrieves a user by their ID
func (g *userGatewayImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var userModel model.User

	err := g.db.WithContext(ctx).Where("uuid = ?", id).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErrors.ErrUserNotFound
		}
		return nil, err
	}

	return g.mapper.ToDomain(&userModel), nil
}

// FindByEmail retrieves a user by their email address
func (g *userGatewayImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var userModel model.User

	err := g.db.WithContext(ctx).Where("email = ?", email).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErrors.ErrUserNotFound
		}
		return nil, err
	}

	return g.mapper.ToDomain(&userModel), nil
}

// Update modifies an existing user
func (g *userGatewayImpl) Update(ctx context.Context, user *entity.User) error {
	userModel := g.mapper.ToModel(user)

	result := g.db.WithContext(ctx).Where("uuid = ?", userModel.UUID).Save(userModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return authErrors.ErrUserNotFound
	}

	return nil
}

// ExistsByEmail checks if a user with the given email already exists
func (g *userGatewayImpl) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64

	err := g.db.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// FindByName retrieves a user by their name (case-insensitive search)
func (g *userGatewayImpl) FindByName(ctx context.Context, name string) (*entity.User, error) {
	var userModel model.User

	// Use ILIKE for case-insensitive search in PostgreSQL
	err := g.db.WithContext(ctx).Where("name ILIKE ?", name).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, authErrors.ErrUserNotFound
		}
		return nil, err
	}

	return g.mapper.ToDomain(&userModel), nil
}

// ListAll retrieves all active users
func (g *userGatewayImpl) ListAll(ctx context.Context) ([]*entity.User, error) {
	var userModels []model.User

	err := g.db.WithContext(ctx).Where("active = ?", true).Find(&userModels).Error
	if err != nil {
		return nil, err
	}

	users := make([]*entity.User, len(userModels))
	for i, userModel := range userModels {
		users[i] = g.mapper.ToDomain(&userModel)
	}

	return users, nil
}

// UpdateProfilePhotoKey atualiza a key da foto de perfil do usuário
func (g *userGatewayImpl) UpdateProfilePhotoKey(ctx context.Context, userID uuid.UUID, profilePhotoKey string) error {
	result := g.db.WithContext(ctx).
		Model(&model.User{}).
		Where("uuid = ?", userID).
		Update("profile_photo_key", profilePhotoKey)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
