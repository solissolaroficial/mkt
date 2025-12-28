package infrastructure

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/seu-usuario/solis-backend/core/domain/service"
)

// hasherServiceImpl implements the HasherService interface using bcrypt
type hasherServiceImpl struct {
	cost int
}

// NewHasherService creates a new instance of hasherServiceImpl
func NewHasherService() service.HasherService {
	return &hasherServiceImpl{
		cost: bcrypt.DefaultCost, // Default cost is 10
	}
}

// NewHasherServiceWithCost creates a new instance of hasherServiceImpl with custom cost
func NewHasherServiceWithCost(cost int) service.HasherService {
	return &hasherServiceImpl{
		cost: cost,
	}
}

// Hash generates a bcrypt hash for the provided password
func (h *hasherServiceImpl) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Compare verifies if the provided password matches the hashed password
func (h *hasherServiceImpl) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
