package request

type CreateShowroomItemRequest struct {
	PDV              string  `json:"pdv" validate:"required,max=200"`
	City             *string `json:"city,omitempty" validate:"omitempty,max=100"`
	Contact          *string `json:"contact,omitempty" validate:"omitempty,max=100"`
	RepName          string  `json:"rep_name" validate:"required,max=100"`
	DeliveryForecast *string `json:"delivery_forecast,omitempty" validate:"omitempty,datetime=2006-01-02"`
	WorkshopDate     *string `json:"workshop_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}

type UpdateShowroomItemRequest struct {
	ID               string  `json:"-" validate:"required,uuid"`
	City             *string `json:"city,omitempty" validate:"omitempty,max=100"`
	Contact          *string `json:"contact,omitempty" validate:"omitempty,max=100"`
	RepName          *string `json:"rep_name,omitempty" validate:"omitempty,max=100"`
	DeliveryForecast *string `json:"delivery_forecast,omitempty" validate:"omitempty,datetime=2006-01-02"`
	WorkshopDate     *string `json:"workshop_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	Delivered        *bool   `json:"delivered,omitempty"`
	PDV              *string `json:"pdv,omitempty" validate:"omitempty,max=200"`
}

type ListShowroomItemsQuery struct {
	RepName   *string `query:"rep_name" validate:"omitempty,max=100"`
	Delivered *bool   `query:"delivered,omitempty"`
	City      *string `query:"city" validate:"omitempty,max=100"`
	Page      int     `query:"page" validate:"omitempty,min=1"`
	Limit     int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sort_by" validate:"omitempty,oneof=created_at delivery_forecast"`
	SortOrder *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
