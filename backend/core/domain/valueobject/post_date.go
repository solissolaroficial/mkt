package valueobject

import (
	"errors"
	"time"
)

var (
	ErrInvalidDate        = errors.New("invalid date format")
	ErrDateCannotBeInPast = errors.New("date cannot be in the past")
)

type PostDate struct {
	date time.Time
}

func NewPostDate(dateStr string) (*PostDate, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, ErrInvalidDate
	}

	// Validar que a data não é no passado
	today := time.Now().Truncate(24 * time.Hour)
	if date.Before(today) {
		return nil, ErrDateCannotBeInPast
	}

	return &PostDate{date: date}, nil
}

// ReconstructPostDate reconstrói a partir de string do banco (sem validação)
func ReconstructPostDate(dateStr string) *PostDate {
	// Assume que dados do banco são válidos
	date, _ := time.Parse("2006-01-02", dateStr)
	return &PostDate{date: date}
}

func (d *PostDate) Value() time.Time {
	return d.date
}

func (d *PostDate) String() string {
	return d.date.Format("2006-01-02")
}
