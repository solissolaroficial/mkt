package response

// AccountPayableResponse representa a resposta de uma conta a pagar
type AccountPayableResponse struct {
	ID            string  `json:"id"`
	Supplier      string  `json:"supplier"`
	Description   string  `json:"description"`
	Amount        float64 `json:"amount"`
	DueDate       string  `json:"due_date"`
	NFArrived     bool    `json:"nf_arrived"`
	BoletoArrived bool    `json:"boleto_arrived"`
	Status        string  `json:"status"`
	Recurrence    string  `json:"recurrence"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}

// AccountPayableListData representa os dados da lista de contas a pagar
type AccountPayableListData struct {
	Accounts []AccountPayableResponse `json:"accounts"`
	Meta     MetaResponse             `json:"meta"`
}
