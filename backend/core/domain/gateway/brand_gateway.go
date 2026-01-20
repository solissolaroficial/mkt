package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

type BrandGateway interface {
	Save(ctx context.Context, brand *entity.Brand) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Brand, error)
	FindByName(ctx context.Context, name string) (*entity.Brand, error)
	List(ctx context.Context) ([]*entity.Brand, error)
	Update(ctx context.Context, brand *entity.Brand) error
	Delete(ctx context.Context, id uuid.UUID) error
	ExistsByName(ctx context.Context, name string) (bool, error)
}
