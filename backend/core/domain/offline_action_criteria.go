package domain

import (
	"errors"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// OfflineActionCriteria representa filtros para busca de ações offline
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type OfflineActionCriteria struct {
	category  *valueobject.OfflineCategory
	repName   *string
	month     *string
	startDate *string
	endDate   *string
	status    *valueobject.OfflineStatus
}

func NewOfflineActionCriteria() *OfflineActionCriteria {
	return &OfflineActionCriteria{}
}

func (c *OfflineActionCriteria) WithCategory(category *valueobject.OfflineCategory) *OfflineActionCriteria {
	c.category = category
	return c
}

func (c *OfflineActionCriteria) WithRepName(repName *string) *OfflineActionCriteria {
	c.repName = repName
	return c
}

func (c *OfflineActionCriteria) WithMonth(month *string) *OfflineActionCriteria {
	c.month = month
	return c
}

func (c *OfflineActionCriteria) WithStartDate(startDate *string) (*OfflineActionCriteria, error) {
	if startDate != nil {
		_, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			return nil, errors.New("invalid start_date format, expected YYYY-MM-DD")
		}
	}
	c.startDate = startDate
	return c, nil
}

func (c *OfflineActionCriteria) WithEndDate(endDate *string) (*OfflineActionCriteria, error) {
	if endDate != nil {
		_, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			return nil, errors.New("invalid end_date format, expected YYYY-MM-DD")
		}
	}
	c.endDate = endDate
	return c, nil
}

func (c *OfflineActionCriteria) WithStatus(status *valueobject.OfflineStatus) *OfflineActionCriteria {
	c.status = status
	return c
}

// Getters para o gateway aplicar os filtros
func (c *OfflineActionCriteria) Category() *valueobject.OfflineCategory { return c.category }
func (c *OfflineActionCriteria) RepName() *string                       { return c.repName }
func (c *OfflineActionCriteria) Month() *string                         { return c.month }
func (c *OfflineActionCriteria) StartDate() *string                     { return c.startDate }
func (c *OfflineActionCriteria) EndDate() *string                       { return c.endDate }
func (c *OfflineActionCriteria) Status() *valueobject.OfflineStatus     { return c.status }

// Validate valida os critérios
func (c *OfflineActionCriteria) Validate() error {
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
