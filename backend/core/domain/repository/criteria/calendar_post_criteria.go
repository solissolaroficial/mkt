package criteria

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// CalendarPostCriteria representa filtros para busca de posts de calendário
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type CalendarPostCriteria struct {
	category   *valueobject.PostCategory
	postType   *valueobject.PostType
	status     *valueobject.PostStatus
	assigneeID *uuid.UUID
	startDate  *string
	endDate    *string
	platform   *string
}

func NewCalendarPostCriteria() *CalendarPostCriteria {
	return &CalendarPostCriteria{}
}

func (c *CalendarPostCriteria) WithCategory(category *valueobject.PostCategory) *CalendarPostCriteria {
	c.category = category
	return c
}

func (c *CalendarPostCriteria) WithType(postType *valueobject.PostType) *CalendarPostCriteria {
	c.postType = postType
	return c
}

func (c *CalendarPostCriteria) WithStatus(status *valueobject.PostStatus) *CalendarPostCriteria {
	c.status = status
	return c
}

func (c *CalendarPostCriteria) WithAssigneeID(assigneeID *uuid.UUID) *CalendarPostCriteria {
	c.assigneeID = assigneeID
	return c
}

func (c *CalendarPostCriteria) WithStartDate(startDate *string) *CalendarPostCriteria {
	c.startDate = startDate
	return c
}

func (c *CalendarPostCriteria) WithEndDate(endDate *string) *CalendarPostCriteria {
	c.endDate = endDate
	return c
}

func (c *CalendarPostCriteria) WithPlatform(platform *string) *CalendarPostCriteria {
	c.platform = platform
	return c
}

// Getters para o gateway aplicar os filtros
func (c *CalendarPostCriteria) Category() *valueobject.PostCategory { return c.category }
func (c *CalendarPostCriteria) Type() *valueobject.PostType         { return c.postType }
func (c *CalendarPostCriteria) Status() *valueobject.PostStatus     { return c.status }
func (c *CalendarPostCriteria) AssigneeID() *uuid.UUID              { return c.assigneeID }
func (c *CalendarPostCriteria) StartDate() *string                  { return c.startDate }
func (c *CalendarPostCriteria) EndDate() *string                    { return c.endDate }
func (c *CalendarPostCriteria) Platform() *string                   { return c.platform }
