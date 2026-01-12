package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RepMarketingAction struct {
	id          uuid.UUID
	repName     string
	date        time.Time
	description string
	month       string // Derivado da data (JAN, FEV, MAR, etc.)
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

// NewRepMarketingAction cria uma nova entidade RepMarketingAction
func NewRepMarketingAction(
	repName string,
	date time.Time,
	description string,
) (*RepMarketingAction, error) {
	repName = strings.TrimSpace(repName)
	if repName == "" {
		return nil, errors.New("repName is required")
	}

	description = strings.TrimSpace(description)
	if description == "" {
		return nil, errors.New("description is required")
	}

	// Derivar mês abreviado da data
	month := deriveRepMarketingMonthFromDate(date)

	action := &RepMarketingAction{
		id:          uuid.New(),
		repName:     repName,
		date:        date,
		description: description,
		month:       month,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}

	if err := action.Validate(); err != nil {
		return nil, err
	}

	return action, nil
}

// ReconstructRepMarketingAction reconstrói a entidade do banco de dados
func ReconstructRepMarketingAction(
	id uuid.UUID,
	repName string,
	date time.Time,
	description string,
	month string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *RepMarketingAction {
	return &RepMarketingAction{
		id:          id,
		repName:     repName,
		date:        date,
		description: description,
		month:       month,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
	}
}

// deriveRepMarketingMonthFromDate deriva o mês abreviado da data para RepMarketingAction
func deriveRepMarketingMonthFromDate(date time.Time) string {
	months := []string{"JAN", "FEV", "MAR", "ABR", "MAI", "JUN", "JUL", "AGO", "SET", "OUT", "NOV", "DEZ"}
	return months[date.Month()-1]
}

// Getters
func (r *RepMarketingAction) ID() uuid.UUID         { return r.id }
func (r *RepMarketingAction) RepName() string       { return r.repName }
func (r *RepMarketingAction) Date() time.Time       { return r.date }
func (r *RepMarketingAction) Description() string   { return r.description }
func (r *RepMarketingAction) Month() string         { return r.month }
func (r *RepMarketingAction) CreatedAt() time.Time  { return r.createdAt }
func (r *RepMarketingAction) UpdatedAt() time.Time  { return r.updatedAt }
func (r *RepMarketingAction) DeletedAt() *time.Time { return r.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (r *RepMarketingAction) Validate() error {
	if r.repName == "" {
		return errors.New("repName is required")
	}

	if r.description == "" {
		return errors.New("description is required")
	}

	if len(r.repName) > 100 {
		return errors.New("repName must be at most 100 characters")
	}

	if len(r.description) > 500 {
		return errors.New("description must be at most 500 characters")
	}

	return nil
}

// UpdateRepName atualiza o nome do representante
func (r *RepMarketingAction) UpdateRepName(repName string) error {
	repName = strings.TrimSpace(repName)
	if repName == "" {
		return errors.New("repName is required")
	}
	if len(repName) > 100 {
		return errors.New("repName must be at most 100 characters")
	}
	r.repName = repName
	r.updatedAt = time.Now()
	return nil
}

// UpdateDate atualiza a data da ação
func (r *RepMarketingAction) UpdateDate(date time.Time) error {
	r.date = date
	r.month = deriveRepMarketingMonthFromDate(date)
	r.updatedAt = time.Now()
	return nil
}

// UpdateDescription atualiza a descrição
func (r *RepMarketingAction) UpdateDescription(description string) error {
	description = strings.TrimSpace(description)
	if description == "" {
		return errors.New("description cannot be empty")
	}
	if len(description) > 500 {
		return errors.New("description must be at most 500 characters")
	}
	r.description = description
	r.updatedAt = time.Now()
	return nil
}

// SoftDelete marca a ação como deletada
func (r *RepMarketingAction) SoftDelete() {
	now := time.Now()
	r.deletedAt = &now
	r.updatedAt = now
}

// IsActive verifica se a ação está ativa (não deletada)
func (r *RepMarketingAction) IsActive() bool {
	return r.deletedAt == nil
}
