package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

type ProgramCredentialGateway interface {
	Save(ctx context.Context, credential *entity.ProgramCredential) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.ProgramCredential, error)
	FindAll(ctx context.Context) ([]*entity.ProgramCredential, error)
	FindActive(ctx context.Context) ([]*entity.ProgramCredential, error)
	Update(ctx context.Context, credential *entity.ProgramCredential) error
	Delete(ctx context.Context, id uuid.UUID) error
	ExistsByName(ctx context.Context, name string) (bool, error)
}
