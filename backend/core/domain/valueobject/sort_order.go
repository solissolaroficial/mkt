package valueobject

import (
	"errors"
	"strings"
)

// SortDirection define a direção da ordenação
type SortDirection string

const (
	// SortDirectionAsc ordenação ascendente
	SortDirectionAsc SortDirection = "ASC"
	// SortDirectionDesc ordenação descendente
	SortDirectionDesc SortDirection = "DESC"
)

// SortOrder define a ordenação de uma query
type SortOrder struct {
	field     string
	direction SortDirection
}

// NewSortOrder cria um novo SortOrder
func NewSortOrder(field string, direction SortDirection) (*SortOrder, error) {
	if field == "" {
		return nil, errors.New("field cannot be empty")
	}

	// Normalizar direção para maiúsculas
	normalizedDirection := SortDirection(strings.ToUpper(string(direction)))
	if normalizedDirection != SortDirectionAsc && normalizedDirection != SortDirectionDesc {
		return nil, errors.New("invalid sort direction, must be ASC or DESC")
	}

	return &SortOrder{
		field:     field,
		direction: normalizedDirection,
	}, nil
}

// GetField retorna o campo de ordenação
func (s *SortOrder) GetField() string {
	return s.field
}

// GetDirection retorna a direção da ordenação
func (s *SortOrder) GetDirection() SortDirection {
	return s.direction
}

// IsAsc retorna true se a ordenação for ascendente
func (s *SortOrder) IsAsc() bool {
	return s.direction == SortDirectionAsc
}

// IsDesc retorna true se a ordenação for descendente
func (s *SortOrder) IsDesc() bool {
	return s.direction == SortDirectionDesc
}

// ToSQLString retorna a string SQL da ordenação (ex: "created_at DESC")
func (s *SortOrder) ToSQLString() string {
	return s.field + " " + string(s.direction)
}
