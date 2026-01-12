package request

type CreateRepMarketingActionRequest struct {
	RepName     string `json:"rep_name" validate:"required,max=100"`
	Date        string `json:"date" validate:"required,datetime=2006-01-02"`
	Description string `json:"description" validate:"required,max=500"`
}

type UpdateRepMarketingActionRequest struct {
	ID          string  `json:"-" validate:"required,uuid"`
	RepName     *string `json:"rep_name,omitempty" validate:"omitempty,max=100"`
	Date        *string `json:"date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type ListRepMarketingActionsQuery struct {
	RepName   *string `query:"rep_name" validate:"omitempty,max=100"`
	Month     *string `query:"month" validate:"omitempty,oneof=JAN FEV MAR ABR MAI JUN JUL AGO SET OUT NOV DEZ"`
	Page      int     `query:"page" validate:"omitempty,min=1"`
	Limit     int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sort_by" validate:"omitempty,oneof=date created_at"`
	SortOrder *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
