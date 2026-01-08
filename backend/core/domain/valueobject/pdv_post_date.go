package valueobject

import (
	"errors"
	"time"
)

// PdvPostDate representa a data de um post PDV
// Value Object imutável para encapsular validação de data
type PdvPostDate struct {
	value time.Time
}

// NewPdvPostDate cria um novo PdvPostDate com validação
func NewPdvPostDate(date time.Time) (*PdvPostDate, error) {
	if date.IsZero() {
		return nil, errors.New("data do post não pode ser zero")
	}

	// Usar timezone do Brasil para consistência
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return nil, errors.New("failed to load timezone")
	}

	// Validar range de datas (mínimo 2020, máximo 1 ano no futuro)
	minDate := time.Date(2020, 1, 1, 0, 0, 0, 0, loc)
	maxDate := time.Now().In(loc).AddDate(1, 0, 0) // 1 ano no futuro
	dateInLoc := date.In(loc)

	if dateInLoc.Before(minDate) || dateInLoc.After(maxDate) {
		return nil, errors.New("data do post fora do range permitido (2020 até 1 ano no futuro)")
	}

	return &PdvPostDate{value: date}, nil
}

// Value retorna a data encapsulada
func (d *PdvPostDate) Value() time.Time {
	return d.value
}

// ReconstructPdvPostDate reconstrói um PdvPostDate a partir de uma string sem validação
// Assume que os dados do banco são válidos (padrão dual constructor)
func ReconstructPdvPostDate(dateStr string) *PdvPostDate {
	date, _ := time.Parse("2006-01-02", dateStr)
	return &PdvPostDate{value: date}
}

// String retorna a data formatada como string
func (d *PdvPostDate) String() string {
	return d.value.Format("2006-01-02")
}

// Equals verifica se duas datas são iguais
func (d *PdvPostDate) Equals(other *PdvPostDate) bool {
	if other == nil {
		return false
	}
	return d.value.Equal(other.value)
}

// Before verifica se esta data é anterior a outra
func (d *PdvPostDate) Before(other *PdvPostDate) bool {
	if other == nil {
		return false
	}
	return d.value.Before(other.value)
}

// After verifica se esta data é posterior a outra
func (d *PdvPostDate) After(other *PdvPostDate) bool {
	if other == nil {
		return false
	}
	return d.value.After(other.value)
}
