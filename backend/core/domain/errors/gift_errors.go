package errors

import (
	"errors"
)

var (
	// GiftItem errors
	ErrGiftItemNotFound        = errors.New("gift item not found")
	ErrGiftItemNameRequired    = errors.New("gift item name is required")
	ErrGiftItemNameTooLong     = errors.New("gift item name is too long")
	ErrGiftItemNameInvalid     = errors.New("gift item name contains invalid characters")
	ErrGiftItemPriceNegative   = errors.New("gift item price cannot be negative")
	ErrGiftItemStockNegative   = errors.New("gift item stock cannot be negative")
	ErrGiftItemNameAlreadyUsed = errors.New("gift item name already used")
	ErrGiftItemHasTransactions = errors.New("gift item has associated transactions and cannot be deleted")

	// GiftTransaction errors
	ErrGiftTransactionNotFound               = errors.New("gift transaction not found")
	ErrGiftTransactionInvalidType            = errors.New("gift transaction type must be 'in' or 'out'")
	ErrGiftTransactionQuantityInvalid        = errors.New("gift transaction quantity must be greater than zero")
	ErrGiftTransactionPriceRequired          = errors.New("gift transaction price is required for entry transactions")
	ErrGiftTransactionRepresentativeRequired = errors.New("gift transaction representative is required for exit transactions")
	ErrGiftTransactionInvalidDate            = errors.New("gift transaction date is invalid")
	ErrInvalidFieldForTransactionType        = errors.New("invalid field for transaction type (price for 'out' or representative for 'in')")
)
