package calendar

import (
	"context"

	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type CreateCalendarPostUseCase struct {
	gateway gateway.CalendarPostGateway
}

type CreateCalendarPostInput struct {
	Title       string
	Description *string
	Date        string
	Time        string
	Caption     *string
	Category    string
	Type        string
	AssigneeID  *uuid.UUID
	Platforms   []string
	ImageURL    *string
}

func NewCreateCalendarPost(gateway gateway.CalendarPostGateway) *CreateCalendarPostUseCase {
	return &CreateCalendarPostUseCase{gateway: gateway}
}

func (uc *CreateCalendarPostUseCase) Execute(ctx context.Context, input CreateCalendarPostInput) (*entity.CalendarPost, error) {
	// Validar e criar value objects
	postDate, err := valueobject.NewPostDate(input.Date)
	if err != nil {
		return nil, err
	}

	postTime, err := valueobject.NewPostTime(input.Time)
	if err != nil {
		return nil, err
	}

	category, err := valueobject.NewPostCategory(input.Category)
	if err != nil {
		return nil, err
	}

	postType, err := valueobject.NewPostType(input.Type)
	if err != nil {
		return nil, err
	}

	// Validar plataformas
	if err := valueobject.ValidatePlatforms(input.Platforms); err != nil {
		return nil, err
	}

	// Criar entidade
	post, err := entity.NewCalendarPost(
		input.Title,
		input.Description,
		postDate,
		postTime,
		input.Caption,
		category,
		postType,
		input.AssigneeID,
		input.Platforms,
		input.ImageURL,
	)
	if err != nil {
		return nil, err
	}

	// Adicionar evento de histórico inicial
	assigneeIDStr := (*input.AssigneeID).String()
	historyEvent := valueobject.NewPostHistoryEvent("upload", assigneeIDStr, nil)
	post.AddHistoryEvent(historyEvent)

	// Salvar via gateway
	if err := uc.gateway.Save(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}
