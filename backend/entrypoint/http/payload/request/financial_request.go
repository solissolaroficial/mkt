package request

type CreateAccountPayableRequest struct {
	Supplier    string  `json:"supplier" validate:"required,max=200"`
	Description string  `json:"description" validate:"required,max=500"`
	Amount      float64 `json:"amount" validate:"required,min=0"`
	DueDate     string  `json:"due_date" validate:"required,datetime=2006-01-02"`
	Recurrence  string  `json:"recurrence" validate:"required,oneof=none monthly yearly"`
}

type UpdateAccountPayableRequest struct {
	Supplier    *string  `json:"supplier" validate:"omitempty,max=200"`
	Description *string  `json:"description" validate:"omitempty,max=500"`
	Amount      *float64 `json:"amount" validate:"omitempty,min=0"`
	DueDate     *string  `json:"due_date" validate:"omitempty,datetime=2006-01-02"`
	Recurrence  *string  `json:"recurrence" validate:"omitempty,oneof=none monthly yearly"`
}

type ListAccountsPayableQuery struct {
	Supplier   *string  `query:"supplier" validate:"omitempty,max=200"`
	MinAmount  *float64 `query:"min_amount" validate:"omitempty,min=0"`
	MaxAmount  *float64 `query:"max_amount" validate:"omitempty,min=0"`
	Status     *string  `query:"status" validate:"omitempty,oneof=pending sent_to_finance"`
	Recurrence *string  `query:"recurrence" validate:"omitempty,oneof=none monthly yearly"`
	StartDate  *string  `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate    *string  `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
	Page       *int     `query:"page" validate:"omitempty,min=1"`
	Limit      *int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy     *string  `query:"sort_by" validate:"omitempty,oneof=due_date amount supplier created_at"`
	SortOrder  *string  `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
