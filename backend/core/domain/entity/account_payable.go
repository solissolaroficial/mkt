package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type AccountPayable struct {
	id            uuid.UUID
	supplier      *valueobject.SupplierName
	dueDate       *valueobject.DueDate
	description   string
	amount        *valueobject.PaymentAmount
	nfArrived     bool
	boletoArrived bool
	status        valueobject.AccountStatus
	recurrence    valueobject.RecurrenceType
	createdAt     time.Time
	updatedAt     time.Time
	deletedAt     *time.Time
}

// NewAccountPayable cria uma nova entidade AccountPayable
func NewAccountPayable(
	supplier string,
	description string,
	amount float64,
	dueDate string,
	recurrence string,
) (*AccountPayable, error) {
	// Validar e criar value objects
	supplierName, err := valueobject.NewSupplierName(supplier)
	if err != nil {
		return nil, err
	}

	dueDateVO, err := valueobject.NewDueDate(dueDate)
	if err != nil {
		return nil, err
	}

	paymentAmount, err := valueobject.NewPaymentAmount(amount)
	if err != nil {
		return nil, err
	}

	recurrenceType, err := valueobject.NewRecurrenceType(recurrence)
	if err != nil {
		return nil, err
	}

	// Validar descrição
	if description == "" {
		return nil, errors.New("description is required")
	}

	// Criar status inicial
	statusPending, err := valueobject.NewAccountStatus("pending")
	if err != nil {
		return nil, err
	}

	account := &AccountPayable{
		id:            uuid.New(),
		supplier:      supplierName,
		dueDate:       dueDateVO,
		description:   description,
		amount:        paymentAmount,
		nfArrived:     false,
		boletoArrived: false,
		status:        statusPending,
		recurrence:    recurrenceType,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}

	if err := account.Validate(); err != nil {
		return nil, err
	}

	return account, nil
}

// ReconstructAccountPayable reconstrói a entidade do banco de dados
func ReconstructAccountPayable(
	id uuid.UUID,
	supplier string,
	description string,
	amount float64,
	dueDate time.Time,
	nfArrived bool,
	boletoArrived bool,
	status string,
	recurrence string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *AccountPayable {
	supplierName := valueobject.ReconstructSupplierName(supplier)
	dueDateVO := valueobject.ReconstructDueDate(dueDate)
	paymentAmount := valueobject.ReconstructPaymentAmount(amount)
	recurrenceType := valueobject.ReconstructRecurrenceType(recurrence)
	accountStatus := valueobject.ReconstructAccountStatus(status)

	return &AccountPayable{
		id:            id,
		supplier:      supplierName,
		dueDate:       dueDateVO,
		description:   description,
		amount:        paymentAmount,
		nfArrived:     nfArrived,
		boletoArrived: boletoArrived,
		status:        accountStatus,
		recurrence:    recurrenceType,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
		deletedAt:     deletedAt,
	}
}

// Getters
func (a *AccountPayable) ID() uuid.UUID                          { return a.id }
func (a *AccountPayable) Supplier() *valueobject.SupplierName    { return a.supplier }
func (a *AccountPayable) DueDate() *valueobject.DueDate          { return a.dueDate }
func (a *AccountPayable) Description() string                    { return a.description }
func (a *AccountPayable) Amount() *valueobject.PaymentAmount     { return a.amount }
func (a *AccountPayable) NFArrived() bool                        { return a.nfArrived }
func (a *AccountPayable) BoletoArrived() bool                    { return a.boletoArrived }
func (a *AccountPayable) Status() valueobject.AccountStatus      { return a.status }
func (a *AccountPayable) Recurrence() valueobject.RecurrenceType { return a.recurrence }
func (a *AccountPayable) CreatedAt() time.Time                   { return a.createdAt }
func (a *AccountPayable) UpdatedAt() time.Time                   { return a.updatedAt }
func (a *AccountPayable) DeletedAt() *time.Time                  { return a.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (a *AccountPayable) Validate() error {
	if a.supplier == nil {
		return errors.New("supplier is required")
	}

	if a.dueDate == nil {
		return errors.New("dueDate is required")
	}

	if a.amount == nil {
		return errors.New("amount is required")
	}

	if a.description == "" {
		return errors.New("description is required")
	}

	return nil
}

// UpdateSupplier atualiza o fornecedor
func (a *AccountPayable) UpdateSupplier(supplier string) error {
	supplierName, err := valueobject.NewSupplierName(supplier)
	if err != nil {
		return err
	}

	a.supplier = supplierName
	a.updatedAt = time.Now()
	return nil
}

// UpdateDescription atualiza a descrição
func (a *AccountPayable) UpdateDescription(description string) error {
	if description == "" {
		return errors.New("description cannot be empty")
	}

	a.description = description
	a.updatedAt = time.Now()
	return nil
}

// UpdateAmount atualiza o valor
func (a *AccountPayable) UpdateAmount(amount float64) error {
	paymentAmount, err := valueobject.NewPaymentAmount(amount)
	if err != nil {
		return err
	}

	a.amount = paymentAmount
	a.updatedAt = time.Now()
	return nil
}

// UpdateDueDate atualiza a data de vencimento
func (a *AccountPayable) UpdateDueDate(dueDate string) error {
	dueDateVO, err := valueobject.NewDueDate(dueDate)
	if err != nil {
		return err
	}

	a.dueDate = dueDateVO
	a.updatedAt = time.Now()
	return nil
}

// UpdateRecurrence atualiza a recorrência
func (a *AccountPayable) UpdateRecurrence(recurrence string) error {
	recurrenceType, err := valueobject.NewRecurrenceType(recurrence)
	if err != nil {
		return err
	}

	a.recurrence = recurrenceType
	a.updatedAt = time.Now()
	return nil
}

// ToggleNF alterna o status da NF
func (a *AccountPayable) ToggleNF() {
	a.nfArrived = !a.nfArrived
	a.updatedAt = time.Now()
}

// ToggleBoleto alterna o status do Boleto
func (a *AccountPayable) ToggleBoleto() {
	a.boletoArrived = !a.boletoArrived
	a.updatedAt = time.Now()
}

// SendToFinance marca como enviado ao financeiro
func (a *AccountPayable) SendToFinance() {
	a.status = valueobject.AccountStatusSentToFinance
	a.updatedAt = time.Now()
}

// SoftDelete marca a conta como deletada
func (a *AccountPayable) SoftDelete() {
	now := time.Now()
	a.deletedAt = &now
	a.updatedAt = now
}

// IsActive verifica se a conta está ativa (não deletada)
func (a *AccountPayable) IsActive() bool {
	return a.deletedAt == nil
}

// IsPending verifica se está pendente
func (a *AccountPayable) IsPending() bool {
	return a.status.IsPending()
}

// IsSentToFinance verifica se foi enviado ao financeiro
func (a *AccountPayable) IsSentToFinance() bool {
	return a.status.IsSentToFinance()
}
