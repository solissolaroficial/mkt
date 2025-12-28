package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type UserMapper struct{}

// ToDomain converte Model (GORM) para Entity (Domínio)
func (m *UserMapper) ToDomain(userModel *model.User) *entity.User {
	return entity.ReconstructUser(
		userModel.UUID,
		userModel.Email,
		userModel.Password,
		userModel.Name,
		userModel.Role,
		userModel.Active,
		userModel.CreatedAt,
		userModel.UpdatedAt,
	)
}

// ToModel converte Entity (Domínio) para Model (GORM)
func (m *UserMapper) ToModel(user *entity.User) *model.User {
	return &model.User{
		UUID:      user.ID(),
		Email:     user.Email(),
		Password:  user.Password(),
		Name:      user.Name(),
		Role:      user.Role(),
		Active:    user.IsActive(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}
