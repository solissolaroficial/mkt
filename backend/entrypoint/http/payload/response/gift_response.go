package response

// Gift Item Response DTOs
type GiftItemResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// GiftItemListData representa os dados da lista de gift items
type GiftItemListData struct {
	Items []GiftItemResponse `json:"items"`
	Meta  MetaResponse       `json:"meta"`
}

// Gift Transaction Response DTOs
type GiftTransactionResponse struct {
	ID                 string   `json:"id"`
	ItemName           string   `json:"item_name"`
	Quantity           int      `json:"quantity"`
	TransactionType    string   `json:"transaction_type"`
	Date               string   `json:"date"`
	Time               string   `json:"time"`
	Price              *float64 `json:"price"`
	RepresentativeUUID *string  `json:"representative_uuid"`
	Unit               string   `json:"unit"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}

// GiftTransactionListData representa os dados da lista de gift transactions
type GiftTransactionListData struct {
	Transactions []GiftTransactionResponse `json:"transactions"`
	Meta         MetaResponse              `json:"meta"`
}
