package request

type CreateShowroomItemRequest struct {
	PDV                string  `json:"pdv" validate:"required,max=200"`
	City               *string `json:"city,omitempty" validate:"omitempty,max=100"`
	Contact            *string `json:"contact,omitempty" validate:"omitempty,max=100"`
	RepresentativeUUID string  `json:"representative_uuid" validate:"required,uuid"`
	DeliveryForecast   *string `json:"delivery_forecast,omitempty" validate:"omitempty,datetime=2006-01-02"`
	WorkshopDate       *string `json:"workshop_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}

type UpdateShowroomItemRequest struct {
	ID                 string  `json:"-" validate:"required,uuid"`
	City               *string `json:"city,omitempty" validate:"omitempty,max=100"`
	Contact            *string `json:"contact,omitempty" validate:"omitempty,max=100"`
	RepresentativeUUID *string `json:"representative_uuid,omitempty" validate:"omitempty,uuid"`
	DeliveryForecast   *string `json:"delivery_forecast,omitempty" validate:"omitempty,datetime=2006-01-02"`
	WorkshopDate       *string `json:"workshop_date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	Delivered          *bool   `json:"delivered,omitempty"`
	PDV                *string `json:"pdv,omitempty" validate:"omitempty,max=200"`
}

type ListShowroomItemsQuery struct {
	RepresentativeUUID *string `query:"representative_uuid" validate:"omitempty,uuid"`
	Delivered          *bool   `query:"delivered,omitempty"`
	City               *string `query:"city,omitempty" validate:"omitempty,max=100"`
	Page               int     `query:"page" validate:"omitempty,min=1"`
	Limit              int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy             *string `query:"sort_by" validate:"omitempty,oneof=created_at delivery_forecast"`
	SortOrder          *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
