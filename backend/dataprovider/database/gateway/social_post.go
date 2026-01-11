package gateway

import (
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"gorm.io/gorm"
)

// NewSocialPostGateway creates a new SocialPostGateway
func NewSocialPostGateway(db *gorm.DB) gateway.SocialPostGateway {
	return NewSocialPostGatewayImpl(db)
}
