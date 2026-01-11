package gateway

import (
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"gorm.io/gorm"
)

// NewSocialDailyAggregationGateway creates a new SocialDailyAggregationGateway
func NewSocialDailyAggregationGateway(db *gorm.DB) gateway.SocialDailyAggregationGateway {
	return NewSocialDailyAggregationGatewayImpl(db)
}
