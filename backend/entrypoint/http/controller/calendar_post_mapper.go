package controller

import (
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	calendarrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	calendarresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// CalendarPostMapper converte entre requests/responses e entities
type CalendarPostMapper struct{}

// NewCalendarPostMapper cria um novo CalendarPostMapper
func NewCalendarPostMapper() *CalendarPostMapper {
	return &CalendarPostMapper{}
}

// ToCalendarPost converte CreateCalendarPostRequest para entity.CalendarPost
func (m *CalendarPostMapper) ToCalendarPost(req *calendarrequest.CreateCalendarPostRequest, assigneeID uuid.UUID) (*entity.CalendarPost, error) {
	postDate, err := valueobject.NewPostDate(req.Date)
	if err != nil {
		return nil, err
	}

	postTime, err := valueobject.NewPostTime(req.Time)
	if err != nil {
		return nil, err
	}

	category, err := valueobject.NewPostCategory(req.Category)
	if err != nil {
		return nil, err
	}

	postType, err := valueobject.NewPostType(req.Type)
	if err != nil {
		return nil, err
	}

	return entity.NewCalendarPost(
		req.Title,
		req.Description,
		postDate,
		postTime,
		req.Caption,
		category,
		postType,
		&assigneeID,
		req.Platforms,
		req.ImageURL,
	)
}

// ToCalendarPostResponse converte entity.CalendarPost para CalendarPostResponse
func (m *CalendarPostMapper) ToCalendarPostResponse(post *entity.CalendarPost) calendarresponse.CalendarPostResponse {
	history := make([]calendarresponse.PostHistoryEvent, len(post.History()))
	for i, event := range post.History() {
		history[i] = calendarresponse.PostHistoryEvent{
			ID:        event.ID,
			Action:    event.Action,
			User:      event.User,
			Text:      event.Text,
			Timestamp: event.Timestamp.Format(time.RFC3339),
		}
	}

	return calendarresponse.CalendarPostResponse{
		ID:                 post.ID().String(),
		Title:              post.Title(),
		Description:        post.Description(),
		Date:               post.Date().String(),
		Time:               post.PostTime().String(),
		Caption:            post.Caption(),
		Category:           post.Category().String(),
		Type:               post.Type().String(),
		Status:             post.Status().String(),
		AssigneeID:         uuidPtrToString(post.AssigneeUUID()),
		Assignee:           nil, // Will be populated by controller if needed
		Platforms:          post.Platforms(),
		PublishedPlatforms: post.PublishedPlatforms(),
		ImageURL:           post.ImageURL(),
		History:            history,
		CreatedAt:          post.CreatedAt().Format(time.RFC3339),
		UpdatedAt:          post.UpdatedAt().Format(time.RFC3339),
	}
}

// ToCalendarPostsListResponse converte lista de entity.CalendarPost para CalendarPostsListResponse
func (m *CalendarPostMapper) ToCalendarPostsListResponse(posts []*entity.CalendarPost, total int64, page, limit int) *calendarresponse.CalendarPostsListResponse {
	data := make([]calendarresponse.CalendarPostResponse, len(posts))
	for i, post := range posts {
		data[i] = m.ToCalendarPostResponse(post)
	}

	// Calcular total de páginas com arredondamento para cima
	totalPages := (int(total) + limit - 1) / limit
	if totalPages == 0 {
		totalPages = 1
	}

	return &calendarresponse.CalendarPostsListResponse{
		Data:       data,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}
