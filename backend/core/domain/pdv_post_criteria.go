package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// PdvPostCriteria representa filtros para busca de posts de PDV
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type PdvPostCriteria struct {
	representativeUUID *uuid.UUID
	month              *string
	platform           *string
	status             *valueobject.PdvStatus
	startDate          *string
	endDate            *string
}

// NewPdvPostCriteria cria um novo PdvPostCriteria vazio
func NewPdvPostCriteria() *PdvPostCriteria {
	return &PdvPostCriteria{}
}

// WithRepresentativeUUID adiciona filtro por UUID do representante
func (c *PdvPostCriteria) WithRepresentativeUUID(representativeUUID *uuid.UUID) *PdvPostCriteria {
	c.representativeUUID = representativeUUID
	return c
}

// WithMonth adiciona filtro por mês
func (c *PdvPostCriteria) WithMonth(month *string) *PdvPostCriteria {
	c.month = month
	return c
}

// WithPlatform adiciona filtro por plataforma
func (c *PdvPostCriteria) WithPlatform(platform *string) *PdvPostCriteria {
	c.platform = platform
	return c
}

// WithStatus adiciona filtro por status
func (c *PdvPostCriteria) WithStatus(status *valueobject.PdvStatus) *PdvPostCriteria {
	c.status = status
	return c
}

// WithStartDate adiciona filtro por data inicial
func (c *PdvPostCriteria) WithStartDate(startDate *string) (*PdvPostCriteria, error) {
	if startDate != nil {
		_, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			return nil, errors.New("invalid start_date format, expected YYYY-MM-DD")
		}
	}
	c.startDate = startDate
	return c, nil
}

// WithEndDate adiciona filtro por data final
func (c *PdvPostCriteria) WithEndDate(endDate *string) (*PdvPostCriteria, error) {
	if endDate != nil {
		_, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			return nil, errors.New("invalid end_date format, expected YYYY-MM-DD")
		}
	}
	c.endDate = endDate
	return c, nil
}

// Getters para o gateway aplicar os filtros

// RepresentativeUUID retorna o filtro de UUID do representante
func (c *PdvPostCriteria) RepresentativeUUID() *uuid.UUID {
	return c.representativeUUID
}

// Month retorna o filtro de mês
func (c *PdvPostCriteria) Month() *string {
	return c.month
}

// Platform retorna o filtro de plataforma
func (c *PdvPostCriteria) Platform() *string {
	return c.platform
}

// Status retorna o filtro de status
func (c *PdvPostCriteria) Status() *valueobject.PdvStatus {
	return c.status
}

// StartDate retorna o filtro de data inicial
func (c *PdvPostCriteria) StartDate() *string {
	return c.startDate
}

// EndDate retorna o filtro de data final
func (c *PdvPostCriteria) EndDate() *string {
	return c.endDate
}

// Validate valida os critérios
func (c *PdvPostCriteria) Validate() error {
	if c.startDate != nil && c.endDate != nil {
		start, err := time.Parse("2006-01-02", *c.startDate)
		if err != nil {
			return errors.New("invalid start_date format")
		}
		end, err := time.Parse("2006-01-02", *c.endDate)
		if err != nil {
			return errors.New("invalid end_date format")
		}
		if start.After(end) {
			return errors.New("start_date must be before or equal to end_date")
		}
	}
	return nil
}
