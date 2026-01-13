package request

// Gift Item Request DTOs
type CreateGiftItemRequest struct {
	Name  string  `json:"name" validate:"required,max=200"`
	Stock int     `json:"stock" validate:"required,min=0"`
	Price float64 `json:"price" validate:"required,min=0"`
}

type UpdateGiftItemRequest struct {
	Name  *string  `json:"name" validate:"omitempty,max=200"`
	Stock *int     `json:"stock" validate:"omitempty,min=0"`
	Price *float64 `json:"price" validate:"omitempty,min=0"`
}

type ListGiftItemsQuery struct {
	Name     *string  `query:"name" validate:"omitempty,max=200"`
	MinStock *int     `query:"min_stock" validate:"omitempty,min=0"`
	MaxStock *int     `query:"max_stock" validate:"omitempty,min=0"`
	MinPrice *float64 `query:"min_price" validate:"omitempty,min=0"`
	MaxPrice *float64 `query:"max_price" validate:"omitempty,min=0"`
	Page     *int     `query:"page" validate:"omitempty,min=1"`
	Limit    *int     `query:"limit" validate:"omitempty,min=1,max=100"`
}

// Gift Transaction Request DTOs
type CreateGiftTransactionRequest struct {
	ItemName        string   `json:"item_name" validate:"required,max=200"`
	Quantity        int      `json:"quantity" validate:"required,min=1"`
	TransactionType string   `json:"transaction_type" validate:"required,oneof=in out"`
	Date            string   `json:"date" validate:"required,datetime=2006-01-02"`
	Time            string   `json:"time" validate:"omitempty,datetime=15:04"`
	Price           *float64 `json:"price" validate:"omitempty,min=0"`
	Representative  *string  `json:"representative" validate:"omitempty,max=200"`
	Unit            string   `json:"unit" validate:"omitempty,max=20"`
}

type UpdateGiftTransactionRequest struct {
	ItemName        *string  `json:"item_name" validate:"omitempty,max=200"`
	Quantity        *int     `json:"quantity" validate:"omitempty,min=1"`
	TransactionType *string  `json:"transaction_type" validate:"omitempty,oneof=in out"`
	Date            *string  `json:"date" validate:"omitempty,datetime=2006-01-02"`
	Time            *string  `json:"time" validate:"omitempty,datetime=15:04"`
	Price           *float64 `json:"price" validate:"omitempty,min=0"`
	Representative  *string  `json:"representative" validate:"omitempty,max=200"`
	Unit            *string  `json:"unit" validate:"omitempty,max=20"`
}

type ListGiftTransactionsQuery struct {
	ItemName        *string `query:"item_name" validate:"omitempty,max=200"`
	TransactionType *string `query:"transaction_type" validate:"omitempty,oneof=in out"`
	Representative  *string `query:"representative" validate:"omitempty,max=200"`
	StartDate       *string `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate         *string `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
	Page            *int    `query:"page" validate:"omitempty,min=1"`
	Limit           *int    `query:"limit" validate:"omitempty,min=1,max=100"`
}
