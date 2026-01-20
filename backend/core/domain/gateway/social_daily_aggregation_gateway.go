package gateway

import (
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type SocialDailyAggregationGateway interface {
	Create(aggregation *entity.SocialDailyAggregation) error
	Update(aggregation *entity.SocialDailyAggregation) error
	Delete(id uuid.UUID) error
	GetByID(id uuid.UUID) (*entity.SocialDailyAggregation, error)
	GetByBrandAndDate(brandID uuid.UUID, date time.Time) (*entity.SocialDailyAggregation, error)
	List(criteria *domain.SocialDailyAggregationCriteria, pagination *valueobject.Pagination) ([]*entity.SocialDailyAggregation, int64, error)
	GetByBrand(brandID uuid.UUID) ([]*entity.SocialDailyAggregation, error)
}
