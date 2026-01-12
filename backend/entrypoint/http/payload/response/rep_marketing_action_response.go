package response

type RepMarketingActionResponse struct {
	UUID        string  `json:"uuid"`
	RepName     string  `json:"rep_name"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Month       string  `json:"month"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at,omitempty"`
}
