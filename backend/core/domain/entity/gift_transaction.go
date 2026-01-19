package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type GiftTransaction struct {
	id                 uuid.UUID
	itemName           string
	quantity           *valueobject.TransactionQuantity
	transactionType    *valueobject.TransactionType
	date               time.Time
	time               string
	price              *valueobject.GiftPrice // Apenas para entradas
	representativeUUID *uuid.UUID             // Apenas para saídas (opcional para entradas)
	unit               string
	createdAt          time.Time
	updatedAt          time.Time
	deletedAt          *time.Time
}

// NewGiftTransaction cria uma nova entidade GiftTransaction
func NewGiftTransaction(
	itemName string,
	quantity int,
	transactionType string,
	date string,
	timeStr string,
	price *float64,
	representativeUUID *uuid.UUID,
	unit string,
) (*GiftTransaction, error) {
	// Validar itemName
	itemName = strings.TrimSpace(itemName)
	if itemName == "" {
		return nil, errors.New("itemName is required")
	}

	// Validar e criar value objects
	txQuantity, err := valueobject.NewTransactionQuantity(quantity)
	if err != nil {
		return nil, err
	}

	txType, err := valueobject.NewTransactionType(transactionType)
	if err != nil {
		return nil, err
	}

	// Validar data (formato esperado: YYYY-MM-DD)
	txDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}

	// Validar time (formato esperado: HH:MM)
	timeStr = strings.TrimSpace(timeStr)
	if timeStr == "" {
		timeStr = "00:00" // Valor padrão
	}

	// Validar price (obrigatório para entradas, opcional para saídas)
	var giftPrice *valueobject.GiftPrice
	if txType.IsEntry() {
		if price == nil {
			return nil, errors.New("price is required for entry transactions")
		}
		giftPrice, err = valueobject.NewGiftPrice(*price)
		if err != nil {
			return nil, err
		}
	} else if price != nil {
		// Se fornecido para saída, ignorar
		giftPrice = nil
	}

	// Validar representativeUUID (obrigatório para saídas, opcional para entradas)
	var repUUID *uuid.UUID
	if txType.IsExit() {
		if representativeUUID == nil {
			return nil, errors.New("representativeUUID is required for exit transactions")
		}
		repUUID = representativeUUID
	} else if representativeUUID != nil {
		// Se fornecido para entrada, usar o UUID fornecido
		repUUID = representativeUUID
	}

	// Validar unit
	unit = strings.TrimSpace(unit)
	if unit == "" {
		unit = "unid." // Valor padrão
	}

	transaction := &GiftTransaction{
		id:                 uuid.New(),
		itemName:           itemName,
		quantity:           txQuantity,
		transactionType:    &txType,
		date:               txDate,
		time:               timeStr,
		price:              giftPrice,
		representativeUUID: repUUID,
		unit:               unit,
		createdAt:          time.Now(),
		updatedAt:          time.Now(),
	}

	if err := transaction.Validate(); err != nil {
		return nil, err
	}

	return transaction, nil
}

// ReconstructGiftTransaction reconstrói a entidade do banco de dados
func ReconstructGiftTransaction(
	id uuid.UUID,
	itemName string,
	quantity int,
	transactionType string,
	date time.Time,
	timeStr string,
	price *float64,
	representativeUUID *uuid.UUID,
	unit string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *GiftTransaction {
	txQuantity := valueobject.ReconstructTransactionQuantity(quantity)
	txType := valueobject.ReconstructTransactionType(transactionType)

	var giftPrice *valueobject.GiftPrice
	if price != nil {
		giftPrice = valueobject.ReconstructGiftPrice(*price)
	}

	return &GiftTransaction{
		id:                 id,
		itemName:           itemName,
		quantity:           txQuantity,
		transactionType:    &txType,
		date:               date,
		time:               timeStr,
		price:              giftPrice,
		representativeUUID: representativeUUID,
		unit:               unit,
		createdAt:          createdAt,
		updatedAt:          updatedAt,
		deletedAt:          deletedAt,
	}
}

// Getters
func (g *GiftTransaction) ID() uuid.UUID                                 { return g.id }
func (g *GiftTransaction) ItemName() string                              { return g.itemName }
func (g *GiftTransaction) Quantity() *valueobject.TransactionQuantity    { return g.quantity }
func (g *GiftTransaction) TransactionType() *valueobject.TransactionType { return g.transactionType }
func (g *GiftTransaction) Date() time.Time                               { return g.date }
func (g *GiftTransaction) Time() string                                  { return g.time }
func (g *GiftTransaction) Price() *valueobject.GiftPrice                 { return g.price }
func (g *GiftTransaction) RepresentativeUUID() *uuid.UUID                { return g.representativeUUID }
func (g *GiftTransaction) Unit() string                                  { return g.unit }
func (g *GiftTransaction) CreatedAt() time.Time                          { return g.createdAt }
func (g *GiftTransaction) UpdatedAt() time.Time                          { return g.updatedAt }
func (g *GiftTransaction) DeletedAt() *time.Time                         { return g.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (g *GiftTransaction) Validate() error {
	if g.itemName == "" {
		return errors.New("itemName is required")
	}

	if g.quantity == nil {
		return errors.New("quantity is required")
	}

	if g.transactionType == nil {
		return errors.New("transactionType is required")
	}

	// Para entradas, preço é obrigatório
	if g.transactionType.IsEntry() && g.price == nil {
		return errors.New("price is required for entry transactions")
	}

	// Para saídas, representanteUUID é obrigatório
	if g.transactionType.IsExit() && g.representativeUUID == nil {
		return errors.New("representativeUUID is required for exit transactions")
	}

	return nil
}

// UpdateItemName atualiza o nome do item
func (g *GiftTransaction) UpdateItemName(itemName string) error {
	itemName = strings.TrimSpace(itemName)
	if itemName == "" {
		return errors.New("itemName cannot be empty")
	}

	g.itemName = itemName
	g.updatedAt = time.Now()
	return nil
}

// UpdateQuantity atualiza a quantidade
func (g *GiftTransaction) UpdateQuantity(quantity int) error {
	txQuantity, err := valueobject.NewTransactionQuantity(quantity)
	if err != nil {
		return err
	}

	g.quantity = txQuantity
	g.updatedAt = time.Now()
	return nil
}

// UpdateTransactionType atualiza o tipo de transação
func (g *GiftTransaction) UpdateTransactionType(transactionType string) error {
	txType, err := valueobject.NewTransactionType(transactionType)
	if err != nil {
		return err
	}

	g.transactionType = &txType
	g.updatedAt = time.Now()
	return nil
}

// UpdatePrice atualiza o preço
func (g *GiftTransaction) UpdatePrice(price *float64) error {
	if price == nil {
		g.price = nil
		g.updatedAt = time.Now()
		return nil
	}

	giftPrice, err := valueobject.NewGiftPrice(*price)
	if err != nil {
		return err
	}

	g.price = giftPrice
	g.updatedAt = time.Now()
	return nil
}

// UpdateRepresentativeUUID atualiza o representante
func (g *GiftTransaction) UpdateRepresentativeUUID(representativeUUID *uuid.UUID) error {
	g.representativeUUID = representativeUUID
	g.updatedAt = time.Now()
	return nil
}

// SoftDelete marca a transação como deletada
func (g *GiftTransaction) SoftDelete() {
	now := time.Now()
	g.deletedAt = &now
	g.updatedAt = now
}

// IsActive verifica se a transação está ativa (não deletada)
func (g *GiftTransaction) IsActive() bool {
	return g.deletedAt == nil
}

// IsEntry verifica se é uma transação de entrada
func (g *GiftTransaction) IsEntry() bool {
	return g.transactionType != nil && g.transactionType.IsEntry()
}

// IsExit verifica se é uma transação de saída
func (g *GiftTransaction) IsExit() bool {
	return g.transactionType != nil && g.transactionType.IsExit()
}
