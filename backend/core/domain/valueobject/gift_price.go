package valueobject

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidGiftPrice  = errors.New("invalid gift price")
	ErrGiftPriceNegative = errors.New("gift price cannot be negative")
)

type GiftPrice struct {
	price float64
}

func NewGiftPrice(price float64) (*GiftPrice, error) {
	if price < 0 {
		return nil, ErrGiftPriceNegative
	}

	return &GiftPrice{price: price}, nil
}

func ReconstructGiftPrice(price float64) *GiftPrice {
	return &GiftPrice{price: price}
}

func (g *GiftPrice) Value() float64 {
	return g.price
}

func (g *GiftPrice) String() string {
	return fmt.Sprintf("R$ %.2f", g.price)
}

// FormatBrazilian retorna o valor no formato brasileiro (R$ 1.234,56)
func (g *GiftPrice) FormatBrazilian() string {
	return fmt.Sprintf("R$ %.2f", g.price)
}
