package gateway

import (
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type SocialPostGateway interface {
	Create(post *entity.SocialPost) error
	Update(post *entity.SocialPost) error
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*entity.SocialPost, error)
	List(criteria *domain.SocialPostCriteria, pagination *valueobject.Pagination) ([]*entity.SocialPost, int64, error)
	ListByBrandAndDate(brandName string, date time.Time) ([]*entity.SocialPost, error)
}
