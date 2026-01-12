package valueobject

import (
	"errors"
	"time"
)

var (
	ErrInvalidActionDate  = errors.New("invalid action date format")
	ErrActionDateInFuture = errors.New("action date cannot be in future")
)

type ActionDate struct {
	date time.Time
}

func NewActionDate(dateStr string) (*ActionDate, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, ErrInvalidActionDate
	}

	// Validar que a data não é no futuro (ações são feitas no passado ou hoje)
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	today := time.Now().In(loc).Truncate(24 * time.Hour)
	dateInLoc := date.In(loc)

	if dateInLoc.After(today) {
		return nil, ErrActionDateInFuture
	}

	return &ActionDate{date: date}, nil
}

// ReconstructActionDate reconstrói a partir de string do banco (sem validação)
func ReconstructActionDate(dateStr string) *ActionDate {
	date, _ := time.Parse("2006-01-02", dateStr)
	return &ActionDate{date: date}
}

func (d *ActionDate) Value() time.Time {
	return d.date
}

func (d *ActionDate) String() string {
	return d.date.Format("2006-01-02")
}

// FormatBrazilian retorna a data no formato brasileiro (DD/MM/YYYY)
func (d *ActionDate) FormatBrazilian() string {
	return d.date.Format("02/01/2006")
}
