package request

type CreateOfflineActionRequest struct {
	RequestedAmount float64 `json:"requested_amount" validate:"required,min=0"`
	ActionDate      string  `json:"action_date" validate:"required,datetime=2006-01-02"`
	Category        string  `json:"category" validate:"required,offline_category"`
	PDV             string  `json:"pdv" validate:"required,max=200"`
	RepName         string  `json:"rep_name" validate:"required,max=100"`
	Observation     string  `json:"observation" validate:"required,max=500"`
}

type UpdateOfflineActionRequest struct {
	ID               string  `json:"-" validate:"required,uuid"`
	ApprovedAmount   *string `json:"approved_amount,omitempty" validate:"omitempty,max=50"`
	OrderNumber      *string `json:"order_number,omitempty" validate:"omitempty,max=50"`
	DepartureDate    *string `json:"departure_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	DeliveryForecast *string `json:"delivery_forecast,omitempty" validate:"omitempty,datetime=2006-01-02"`
	DeliveryDate     *string `json:"delivery_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	City             *string `json:"city,omitempty" validate:"omitempty,max=100"`
	UF               *string `json:"uf,omitempty" validate:"omitempty,max=2"`
	Scored           *string `json:"scored,omitempty" validate:"omitempty,oneof=SIM NÃO AINDA NÃO"`
	Status           *string `json:"status,omitempty" validate:"omitempty,oneof=pending approved rejected completed"`
	Observation      *string `json:"observation,omitempty" validate:"omitempty,max=500"`
	PDV              *string `json:"pdv,omitempty" validate:"omitempty,max=200"`
	RepName          *string `json:"rep_name,omitempty" validate:"omitempty,max=100"`
}

type ListOfflineActionsQuery struct {
	Category  *string `query:"category" validate:"omitempty,offline_category"`
	RepName   *string `query:"rep_name" validate:"omitempty,max=100"`
	Month     *string `query:"month" validate:"omitempty,oneof=JAN FEV MAR ABR MAI JUN JUL AGO SET OUT NOV DEZ"`
	StartDate *string `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate   *string `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
	Status    *string `query:"status" validate:"omitempty,oneof=pending approved rejected completed"`
	Page      int     `query:"page" validate:"omitempty,min=1"`
	Limit     int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sort_by" validate:"omitempty,oneof=action_date created_at"`
	SortOrder *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
