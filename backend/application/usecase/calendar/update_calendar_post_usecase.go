package calendar

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type UpdateCalendarPostUseCase struct {
	gateway gateway.CalendarPostGateway
}

type UpdateCalendarPostInput struct {
	ID          string
	Title       *string
	Description *string
	Caption     *string
	ImageURL    *string
}

func NewUpdateCalendarPost(gateway gateway.CalendarPostGateway) *UpdateCalendarPostUseCase {
	return &UpdateCalendarPostUseCase{gateway: gateway}
}

func (uc *UpdateCalendarPostUseCase) Execute(ctx context.Context, input UpdateCalendarPostInput) (*entity.CalendarPost, error) {
	// Buscar post existente
	post, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domainErrors.ErrCalendarPostNotFound
	}

	// Atualizar campos usando métodos da entity
	if input.Title != nil {
		if err := post.UpdateTitle(*input.Title); err != nil {
			return nil, err
		}
	}

	if input.Description != nil {
		post.UpdateDescription(input.Description)
	}

	if input.Caption != nil {
		post.UpdateCaption(input.Caption)
	}

	if input.ImageURL != nil {
		post.UpdateImageURL(input.ImageURL)
	}

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}
