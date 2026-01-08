package request

import (
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// CreateCalendarPostRequest representa a requisição para criar um post no calendário
type CreateCalendarPostRequest struct {
	Title       string   `json:"title" validate:"required,max=500"`
	Description *string  `json:"description"`
	Date        string   `json:"date" validate:"required"`
	Time        string   `json:"time" validate:"required"`
	Caption     *string  `json:"caption"`
	Category    string   `json:"category" validate:"required,oneof=official solis_voce leonardo luiz"`
	Type        string   `json:"type" validate:"required,oneof=video static carousel story article_linkedin article_blog"`
	AssigneeID  string   `json:"assignee_id"`
	Platforms   []string `json:"platforms" validate:"required"`
	ImageURL    *string  `json:"image_url"`
}

// UpdateCalendarPostRequest representa a requisição para atualizar um post no calendário
type UpdateCalendarPostRequest struct {
	Title       *string   `json:"title" validate:"omitempty,max=500"`
	Description *string   `json:"description"`
	Date        *string   `json:"date"`
	Time        *string   `json:"time"`
	Caption     *string   `json:"caption"`
	AssigneeID  *string   `json:"assignee_id"`
	Platforms   *[]string `json:"platforms"`
	ImageURL    *string   `json:"image_url"`
}

// UpdateCalendarPostStatusRequest representa a requisição para atualizar o status de um post
type UpdateCalendarPostStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=in_progress review adjust approved published"`
}

// ConfirmCalendarPostPublishingRequest representa a requisição para confirmar a publicação de um post
type ConfirmCalendarPostPublishingRequest struct {
	PublishedPlatforms []string `json:"published_platforms" validate:"required"`
}

// ListCalendarPostsRequest representa a requisição para listar posts do calendário
type ListCalendarPostsRequest struct {
	Page       int     `json:"page" validate:"min=0"`
	Limit      int     `json:"limit" validate:"min=0"`
	Category   *string `json:"category" validate:"omitempty,oneof=official solis_voce leonardo luiz"`
	Type       *string `json:"type" validate:"omitempty,oneof=video static carousel story article_linkedin article_blog"`
	Status     *string `json:"status" validate:"omitempty,oneof=in_progress review adjust approved published"`
	AssigneeID *string `json:"assignee_id"`
	StartDate  *string `json:"start_date"`
	EndDate    *string `json:"end_date"`
	Platform   *string `json:"platform" validate:"omitempty,oneof=instagram facebook linkedin youtube tiktok"`
	SortBy     *string `json:"sort_by"`
	SortOrder  *string `json:"sort_order" validate:"omitempty,oneof=ASC DESC"`
}

// ToCriteria converte a requisição para CalendarPostCriteria
func (r *ListCalendarPostsRequest) ToCriteria() (*domain.CalendarPostCriteria, error) {
	crit := domain.NewCalendarPostCriteria()

	if r.Category != nil {
		category, err := valueobject.NewPostCategory(*r.Category)
		if err == nil {
			crit = crit.WithCategory(&category)
		}
	}

	if r.Type != nil {
		postType, err := valueobject.NewPostType(*r.Type)
		if err == nil {
			crit = crit.WithType(&postType)
		}
	}

	if r.Status != nil {
		status, err := valueobject.NewPostStatus(*r.Status)
		if err == nil {
			crit = crit.WithStatus(&status)
		}
	}

	if r.StartDate != nil {
		crit = crit.WithStartDate(r.StartDate)
	}

	if r.EndDate != nil {
		crit = crit.WithEndDate(r.EndDate)
	}

	if r.Platform != nil {
		crit = crit.WithPlatform(r.Platform)
	}

	return crit, nil
}
