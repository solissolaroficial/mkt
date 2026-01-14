package valueobject

import (
	"errors"
	"time"
)

var (
	ErrInvalidDueDate = errors.New("invalid due date")
)

type DueDate struct {
	date time.Time
}

func NewDueDate(dateStr string) (*DueDate, error) {
	// Validar data (formato esperado: YYYY-MM-DD)
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, ErrInvalidDueDate
	}

	return &DueDate{date: date}, nil
}

func ReconstructDueDate(date time.Time) *DueDate {
	return &DueDate{date: date}
}

func (d *DueDate) Value() time.Time {
	return d.date
}

func (d *DueDate) String() string {
	return d.date.Format("2006-01-02")
}

// FormatBrazilian retorna a data no formato brasileiro (DD/MM/YYYY)
func (d *DueDate) FormatBrazilian() string {
	return d.date.Format("02/01/2006")
}
