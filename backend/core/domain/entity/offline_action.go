package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/valueobject"

	"github.com/google/uuid"
)

type OfflineAction struct {
	id               uuid.UUID
	requestedAmount  *valueobject.MonetaryValue
	actionDate       *valueobject.ActionDate
	category         *valueobject.OfflineCategory
	month            string  // Derivado da data (JAN, FEV, MAR, etc.)
	approvedAmount   *string // Pode ser vazio (em análise)
	orderNumber      *string
	departureDate    *string
	deliveryForecast *string
	deliveryDate     *string
	city             *string
	uf               *string
	scored           *valueobject.ScoredStatus
	status           *valueobject.OfflineStatus
	observation      *string // Descrição/status
	pdv              *string
	repName          string
	createdAt        time.Time
	updatedAt        time.Time
	deletedAt        *time.Time
}

// NewOfflineAction cria uma nova entidade OfflineAction
func NewOfflineAction(
	requestedAmount float64,
	actionDate *valueobject.ActionDate,
	category valueobject.OfflineCategory,
	pdv string,
	repName string,
	observation string,
) (*OfflineAction, error) {
	// Validar e criar value object para valor solicitado
	monetary, err := valueobject.NewMonetaryValue(requestedAmount)
	if err != nil {
		return nil, err
	}

	if actionDate == nil {
		return nil, errors.New("actionDate is required")
	}

	// Validar category
	if !category.IsValid() {
		return nil, errors.New("invalid category")
	}

	pdv = strings.TrimSpace(pdv)
	if pdv == "" {
		return nil, errors.New("pdv is required")
	}

	repName = strings.TrimSpace(repName)
	if repName == "" {
		return nil, errors.New("repName is required")
	}

	// Derivar mês da data
	month := deriveOfflineMonthFromDate(actionDate.Value())

	// Criar status inicial
	status, err := valueobject.NewOfflineStatus("pending")
	if err != nil {
		return nil, err
	}

	// Criar scored status inicial
	scored, err := valueobject.NewScoredStatus("AINDA NÃO")
	if err != nil {
		return nil, err
	}

	// Handle observation and pdv as pointers
	var obs *string
	if observation != "" {
		obsTrimmed := strings.TrimSpace(observation)
		obs = &obsTrimmed
	}

	pdvTrimmed := strings.TrimSpace(pdv)

	action := &OfflineAction{
		id:               uuid.New(),
		requestedAmount:  monetary,
		actionDate:       actionDate,
		category:         &category,
		month:            month,
		approvedAmount:   nil, // Inicialmente vazio (em análise)
		orderNumber:      nil,
		departureDate:    nil,
		deliveryForecast: nil,
		deliveryDate:     nil,
		city:             nil,
		uf:               nil,
		scored:           &scored,
		status:           &status,
		observation:      obs,
		pdv:              &pdvTrimmed,
		repName:          repName,
		createdAt:        time.Now(),
		updatedAt:        time.Now(),
	}

	if err := action.Validate(); err != nil {
		return nil, err
	}

	return action, nil
}

// ReconstructOfflineAction reconstrói a entidade do banco de dados
func ReconstructOfflineAction(
	id uuid.UUID,
	requestedAmount float64,
	actionDate *valueobject.ActionDate,
	category *valueobject.OfflineCategory,
	month string,
	approvedAmount *string,
	orderNumber *string,
	departureDate *string,
	deliveryForecast *string,
	deliveryDate *string,
	city *string,
	uf *string,
	scored *valueobject.ScoredStatus,
	status *valueobject.OfflineStatus,
	observation *string,
	pdv *string,
	repName string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *OfflineAction {
	monetary := valueobject.ReconstructMonetaryValue(requestedAmount)

	return &OfflineAction{
		id:               id,
		requestedAmount:  monetary,
		actionDate:       actionDate,
		category:         category,
		month:            month,
		approvedAmount:   approvedAmount,
		orderNumber:      orderNumber,
		departureDate:    departureDate,
		deliveryForecast: deliveryForecast,
		deliveryDate:     deliveryDate,
		city:             city,
		uf:               uf,
		scored:           scored,
		status:           status,
		observation:      observation,
		pdv:              pdv,
		repName:          repName,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
		deletedAt:        deletedAt,
	}
}

// deriveOfflineMonthFromDate deriva o mês abreviado da data para OfflineAction
func deriveOfflineMonthFromDate(date time.Time) string {
	months := []string{"JAN", "FEV", "MAR", "ABR", "MAI", "JUN", "JUL", "AGO", "SET", "OUT", "NOV", "DEZ"}
	return months[date.Month()-1]
}

// Getters
func (o *OfflineAction) ID() uuid.UUID                               { return o.id }
func (o *OfflineAction) RequestedAmount() *valueobject.MonetaryValue { return o.requestedAmount }
func (o *OfflineAction) ActionDate() *valueobject.ActionDate         { return o.actionDate }
func (o *OfflineAction) Category() *valueobject.OfflineCategory      { return o.category }
func (o *OfflineAction) Month() string                               { return o.month }
func (o *OfflineAction) ApprovedAmount() *string                     { return o.approvedAmount }
func (o *OfflineAction) OrderNumber() *string                        { return o.orderNumber }
func (o *OfflineAction) DepartureDate() *string                      { return o.departureDate }
func (o *OfflineAction) DeliveryForecast() *string                   { return o.deliveryForecast }
func (o *OfflineAction) DeliveryDate() *string                       { return o.deliveryDate }
func (o *OfflineAction) City() *string                               { return o.city }
func (o *OfflineAction) UF() *string                                 { return o.uf }
func (o *OfflineAction) Scored() *valueobject.ScoredStatus           { return o.scored }
func (o *OfflineAction) Status() *valueobject.OfflineStatus          { return o.status }
func (o *OfflineAction) Observation() *string                        { return o.observation }
func (o *OfflineAction) PDV() *string                                { return o.pdv }
func (o *OfflineAction) RepName() string                             { return o.repName }
func (o *OfflineAction) CreatedAt() time.Time                        { return o.createdAt }
func (o *OfflineAction) UpdatedAt() time.Time                        { return o.updatedAt }
func (o *OfflineAction) DeletedAt() *time.Time                       { return o.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (o *OfflineAction) Validate() error {
	if o.requestedAmount == nil {
		return errors.New("requestedAmount is required")
	}

	if o.actionDate == nil {
		return errors.New("actionDate is required")
	}

	if o.status == nil {
		return errors.New("status is required")
	}

	if o.scored == nil {
		return errors.New("scored is required")
	}

	if o.repName == "" {
		return errors.New("repName is required")
	}

	return nil
}

// UpdateApprovedAmount atualiza o valor aprovado
func (o *OfflineAction) UpdateApprovedAmount(amount *string) error {
	o.approvedAmount = amount
	o.updatedAt = time.Now()
	return nil
}

// UpdateOrderNumber atualiza o número do pedido
func (o *OfflineAction) UpdateOrderNumber(orderNumber *string) error {
	if orderNumber != nil && *orderNumber != "" {
		if len(*orderNumber) > 50 {
			return errors.New("orderNumber must be at most 50 characters")
		}
	}

	o.orderNumber = orderNumber
	o.updatedAt = time.Now()
	return nil
}

// UpdateDepartureDate atualiza a data de saída
func (o *OfflineAction) UpdateDepartureDate(date *string) error {
	o.departureDate = date
	o.updatedAt = time.Now()
	return nil
}

// UpdateDeliveryForecast atualiza a previsão de entrega
func (o *OfflineAction) UpdateDeliveryForecast(forecast *string) error {
	o.deliveryForecast = forecast
	o.updatedAt = time.Now()
	return nil
}

// UpdateDeliveryDate atualiza a data de entrega
func (o *OfflineAction) UpdateDeliveryDate(date *string) error {
	o.deliveryDate = date
	o.updatedAt = time.Now()
	return nil
}

// UpdateLocation atualiza cidade e UF
func (o *OfflineAction) UpdateLocation(city *string, uf *string) error {
	if city != nil && len(*city) > 100 {
		return errors.New("city must be at most 100 characters")
	}

	if uf != nil && len(*uf) > 2 {
		return errors.New("uf must be at most 2 characters")
	}

	o.city = city
	o.uf = uf
	o.updatedAt = time.Now()
	return nil
}

// UpdateScored atualiza o status de pontuação
func (o *OfflineAction) UpdateScored(scored *valueobject.ScoredStatus) error {
	if scored != nil && !scored.IsValid() {
		return errors.New("invalid scored status")
	}

	o.scored = scored
	o.updatedAt = time.Now()
	return nil
}

// UpdateStatus atualiza o status com validação de transição
func (o *OfflineAction) UpdateStatus(newStatus *valueobject.OfflineStatus) error {
	if newStatus == nil {
		return errors.New("new status is required")
	}

	// Validar transição de status
	if !o.status.CanTransitionTo(*newStatus) {
		return errors.New("invalid status transition")
	}

	o.status = newStatus
	o.updatedAt = time.Now()
	return nil
}

// UpdateObservation atualiza a observação
func (o *OfflineAction) UpdateObservation(observation *string) error {
	if observation != nil && len(*observation) > 500 {
		return errors.New("observation must be at most 500 characters")
	}

	o.observation = observation
	o.updatedAt = time.Now()
	return nil
}

// UpdatePDV atualiza o PDV
func (o *OfflineAction) UpdatePDV(pdv *string) error {
	if pdv != nil {
		pdvTrimmed := strings.TrimSpace(*pdv)
		if pdvTrimmed == "" {
			return errors.New("pdv cannot be empty")
		}
		if len(pdvTrimmed) > 200 {
			return errors.New("pdv must be at most 200 characters")
		}
		o.pdv = &pdvTrimmed
	}
	o.updatedAt = time.Now()
	return nil
}

// UpdateRepName atualiza o nome do representante
func (o *OfflineAction) UpdateRepName(repName string) error {
	repName = strings.TrimSpace(repName)
	if repName == "" {
		return errors.New("repName is required")
	}
	if len(repName) > 100 {
		return errors.New("repName must be at most 100 characters")
	}
	o.repName = repName
	o.updatedAt = time.Now()
	return nil
}

// SoftDelete marca a ação como deletada
func (o *OfflineAction) SoftDelete() {
	now := time.Now()
	o.deletedAt = &now
	o.updatedAt = now
}

// IsActive verifica se a ação está ativa (não deletada)
func (o *OfflineAction) IsActive() bool {
	return o.deletedAt == nil
}

// IsApproved verifica se a ação está aprovada
func (o *OfflineAction) IsApproved() bool {
	return o.status != nil && o.status.String() == "approved"
}

// IsCompleted verifica se a ação está concluída
func (o *OfflineAction) IsCompleted() bool {
	return o.status != nil && o.status.String() == "completed"
}
