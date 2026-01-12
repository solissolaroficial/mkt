package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ShowroomItem struct {
	id               uuid.UUID
	pdv              string
	city             *string
	contact          *string
	repName          string
	deliveryForecast *string
	workshopDate     *string
	delivered        bool
	createdAt        time.Time
	updatedAt        time.Time
	deletedAt        *time.Time
}

// NewShowroomItem cria uma nova entidade ShowroomItem
func NewShowroomItem(
	pdv string,
	repName string,
) (*ShowroomItem, error) {
	pdv = strings.TrimSpace(pdv)
	if pdv == "" {
		return nil, errors.New("pdv is required")
	}

	repName = strings.TrimSpace(repName)
	if repName == "" {
		return nil, errors.New("repName is required")
	}

	item := &ShowroomItem{
		id:        uuid.New(),
		pdv:       pdv,
		repName:   repName,
		delivered: false, // Inicialmente não entregue
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if err := item.Validate(); err != nil {
		return nil, err
	}

	return item, nil
}

// ReconstructShowroomItem reconstrói a entidade do banco de dados
func ReconstructShowroomItem(
	id uuid.UUID,
	pdv string,
	city *string,
	contact *string,
	repName string,
	deliveryForecast *string,
	workshopDate *string,
	delivered bool,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *ShowroomItem {
	return &ShowroomItem{
		id:               id,
		pdv:              pdv,
		city:             city,
		contact:          contact,
		repName:          repName,
		deliveryForecast: deliveryForecast,
		workshopDate:     workshopDate,
		delivered:        delivered,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
		deletedAt:        deletedAt,
	}
}

// Getters
func (s *ShowroomItem) ID() uuid.UUID             { return s.id }
func (s *ShowroomItem) PDV() string               { return s.pdv }
func (s *ShowroomItem) City() *string             { return s.city }
func (s *ShowroomItem) Contact() *string          { return s.contact }
func (s *ShowroomItem) RepName() string           { return s.repName }
func (s *ShowroomItem) DeliveryForecast() *string { return s.deliveryForecast }
func (s *ShowroomItem) WorkshopDate() *string     { return s.workshopDate }
func (s *ShowroomItem) Delivered() bool           { return s.delivered }
func (s *ShowroomItem) CreatedAt() time.Time      { return s.createdAt }
func (s *ShowroomItem) UpdatedAt() time.Time      { return s.updatedAt }
func (s *ShowroomItem) DeletedAt() *time.Time     { return s.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (s *ShowroomItem) Validate() error {
	if s.pdv == "" {
		return errors.New("pdv is required")
	}

	if s.repName == "" {
		return errors.New("repName is required")
	}

	if s.city != nil && len(*s.city) > 100 {
		return errors.New("city must be at most 100 characters")
	}

	if s.contact != nil && len(*s.contact) > 100 {
		return errors.New("contact must be at most 100 characters")
	}

	return nil
}

// UpdateCity atualiza a cidade
func (s *ShowroomItem) UpdateCity(city *string) error {
	if city != nil && len(*city) > 100 {
		return errors.New("city must be at most 100 characters")
	}

	s.city = city
	s.updatedAt = time.Now()
	return nil
}

// UpdateContact atualiza o contato
func (s *ShowroomItem) UpdateContact(contact *string) error {
	if contact != nil && len(*contact) > 100 {
		return errors.New("contact must be at most 100 characters")
	}

	s.contact = contact
	s.updatedAt = time.Now()
	return nil
}

// UpdateRepName atualiza o nome do representante
func (s *ShowroomItem) UpdateRepName(repName string) error {
	repName = strings.TrimSpace(repName)
	if repName == "" {
		return errors.New("repName is required")
	}
	if len(repName) > 100 {
		return errors.New("repName must be at most 100 characters")
	}
	s.repName = repName
	s.updatedAt = time.Now()
	return nil
}

// UpdateDeliveryForecast atualiza a previsão de entrega
func (s *ShowroomItem) UpdateDeliveryForecast(forecast *string) error {
	s.deliveryForecast = forecast
	s.updatedAt = time.Now()
	return nil
}

// UpdateWorkshopDate atualiza a data do workshop
func (s *ShowroomItem) UpdateWorkshopDate(date *string) error {
	s.workshopDate = date
	s.updatedAt = time.Now()
	return nil
}

// MarkAsDelivered marca o item como entregue
func (s *ShowroomItem) MarkAsDelivered() {
	s.delivered = true
	s.updatedAt = time.Now()
}

// MarkAsNotDelivered marca o item como não entregue
func (s *ShowroomItem) MarkAsNotDelivered() {
	s.delivered = false
	s.updatedAt = time.Now()
}

// UpdatePDV atualiza o PDV
func (s *ShowroomItem) UpdatePDV(pdv string) error {
	pdv = strings.TrimSpace(pdv)
	if pdv == "" {
		return errors.New("pdv cannot be empty")
	}
	if len(pdv) > 200 {
		return errors.New("pdv must be at most 200 characters")
	}
	s.pdv = pdv
	s.updatedAt = time.Now()
	return nil
}

// SoftDelete marca o item como deletado
func (s *ShowroomItem) SoftDelete() {
	now := time.Now()
	s.deletedAt = &now
	s.updatedAt = now
}

// IsActive verifica se o item está ativo (não deletado)
func (s *ShowroomItem) IsActive() bool {
	return s.deletedAt == nil
}
