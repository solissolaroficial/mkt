package request

type CreateRepMarketingActionRequest struct {
	RepresentativeUUID string `json:"representative_uuid" validate:"required,uuid"`
	Date               string `json:"date" validate:"required,datetime=2006-01-02"`
	Description        string `json:"description" validate:"required,max=500"`
}

type UpdateRepMarketingActionRequest struct {
	ID                 string  `json:"-" validate:"required,uuid"`
	RepresentativeUUID *string `json:"representative_uuid,omitempty" validate:"omitempty,uuid"`
	Date               *string `json:"date,omitempty" validate:"omitempty,datetime=2006-01-02"`
	Description        *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

type ListRepMarketingActionsQuery struct {
	RepresentativeUUID *string `query:"representative_uuid" validate:"omitempty,uuid"`
	Month              *string `query:"month" validate:"omitempty,oneof=JAN FEV MAR ABR MAI JUN JUL AGO SET OUT NOV DEZ"`
	Page               int     `query:"page" validate:"omitempty,min=1"`
	Limit              int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy             *string `query:"sort_by" validate:"omitempty,oneof=date created_at"`
	SortOrder          *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
