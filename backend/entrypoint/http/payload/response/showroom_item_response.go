package response

type ShowroomItemResponse struct {
	UUID             string  `json:"uuid"`
	PDV              string  `json:"pdv"`
	City             *string `json:"city,omitempty"`
	Contact          *string `json:"contact,omitempty"`
	RepName          string  `json:"rep_name"`
	DeliveryForecast *string `json:"delivery_forecast,omitempty"`
	WorkshopDate     *string `json:"workshop_date,omitempty"`
	Delivered        bool    `json:"delivered"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	DeletedAt        *string `json:"deleted_at,omitempty"`
}
