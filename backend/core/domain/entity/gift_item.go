package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type GiftItem struct {
	id        uuid.UUID
	name      *valueobject.GiftName
	stock     int
	price     *valueobject.GiftPrice
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// NewGiftItem cria uma nova entidade GiftItem
func NewGiftItem(
	name string,
	stock int,
	price float64,
) (*GiftItem, error) {
	// Validar e criar value objects
	giftName, err := valueobject.NewGiftName(name)
	if err != nil {
		return nil, err
	}

	giftPrice, err := valueobject.NewGiftPrice(price)
	if err != nil {
		return nil, err
	}

	// Validar stock
	if stock < 0 {
		return nil, errors.New("stock cannot be negative")
	}

	item := &GiftItem{
		id:        uuid.New(),
		name:      giftName,
		stock:     stock,
		price:     giftPrice,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if err := item.Validate(); err != nil {
		return nil, err
	}

	return item, nil
}

// ReconstructGiftItem reconstrói a entidade do banco de dados
func ReconstructGiftItem(
	id uuid.UUID,
	name string,
	stock int,
	price float64,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *GiftItem {
	giftName := valueobject.ReconstructGiftName(name)
	giftPrice := valueobject.ReconstructGiftPrice(price)

	return &GiftItem{
		id:        id,
		name:      giftName,
		stock:     stock,
		price:     giftPrice,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

// Getters
func (g *GiftItem) ID() uuid.UUID                 { return g.id }
func (g *GiftItem) Name() *valueobject.GiftName   { return g.name }
func (g *GiftItem) Stock() int                    { return g.stock }
func (g *GiftItem) Price() *valueobject.GiftPrice { return g.price }
func (g *GiftItem) CreatedAt() time.Time          { return g.createdAt }
func (g *GiftItem) UpdatedAt() time.Time          { return g.updatedAt }
func (g *GiftItem) DeletedAt() *time.Time         { return g.deletedAt }

// Métodos de Negócio

// Validate valida os dados da entidade
func (g *GiftItem) Validate() error {
	if g.name == nil {
		return errors.New("name is required")
	}

	if g.stock < 0 {
		return errors.New("stock cannot be negative")
	}

	if g.price == nil {
		return errors.New("price is required")
	}

	return nil
}

// UpdateName atualiza o nome do brinde
func (g *GiftItem) UpdateName(name string) error {
	giftName, err := valueobject.NewGiftName(name)
	if err != nil {
		return err
	}

	g.name = giftName
	g.updatedAt = time.Now()
	return nil
}

// UpdateStock atualiza o estoque
func (g *GiftItem) UpdateStock(stock int) error {
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	g.stock = stock
	g.updatedAt = time.Now()
	return nil
}

// UpdatePrice atualiza o preço
func (g *GiftItem) UpdatePrice(price float64) error {
	giftPrice, err := valueobject.NewGiftPrice(price)
	if err != nil {
		return err
	}

	g.price = giftPrice
	g.updatedAt = time.Now()
	return nil
}

// AdjustStock ajusta o estoque (usado por transações)
func (g *GiftItem) AdjustStock(delta int) error {
	newStock := g.stock + delta

	if newStock < 0 {
		return errors.New("insufficient stock")
	}

	g.stock = newStock
	g.updatedAt = time.Now()
	return nil
}

// SoftDelete marca o item como deletado
func (g *GiftItem) SoftDelete() {
	now := time.Now()
	g.deletedAt = &now
	g.updatedAt = now
}

// IsActive verifica se o item está ativo (não deletado)
func (g *GiftItem) IsActive() bool {
	return g.deletedAt == nil
}

// HasStock verifica se há estoque suficiente
func (g *GiftItem) HasStock(quantity int) bool {
	return g.stock >= quantity
}
