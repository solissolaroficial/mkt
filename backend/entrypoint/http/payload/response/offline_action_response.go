package response

type OfflineActionResponse struct {
	UUID               string  `json:"uuid"`
	RequestedAmount    float64 `json:"requested_amount"`
	ApprovedAmount     *string `json:"approved_amount"`
	OrderNumber        *string `json:"order_number"`
	ActionDate         string  `json:"action_date"`
	DepartureDate      *string `json:"departure_date"`
	DeliveryForecast   *string `json:"delivery_forecast"`
	DeliveryDate       *string `json:"delivery_date"`
	City               *string `json:"city"`
	UF                 *string `json:"uf"`
	Category           string  `json:"category"`
	PDV                string  `json:"pdv"`
	RepresentativeUUID string  `json:"representative_uuid"`
	Observation        *string `json:"observation"`
	Scored             *string `json:"scored"`
	Status             string  `json:"status"`
	Month              string  `json:"month"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	DeletedAt          *string `json:"deleted_at,omitempty"`
}
