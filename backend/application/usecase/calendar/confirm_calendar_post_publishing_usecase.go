package calendar

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type ConfirmCalendarPostPublishingUseCase struct {
	gateway gateway.CalendarPostGateway
}

type ConfirmPublishingInput struct {
	PostID    string
	Platforms []string
	User      string
}

func NewConfirmCalendarPostPublishing(gateway gateway.CalendarPostGateway) *ConfirmCalendarPostPublishingUseCase {
	return &ConfirmCalendarPostPublishingUseCase{gateway: gateway}
}

func (uc *ConfirmCalendarPostPublishingUseCase) Execute(ctx context.Context, input ConfirmPublishingInput) (*entity.CalendarPost, error) {
	// Buscar post existente
	post, err := uc.gateway.FindByID(ctx, input.PostID)
	if err != nil {
		return nil, domainErrors.ErrCalendarPostNotFound
	}

	// Validar plataformas
	if err := valueobject.ValidatePlatforms(input.Platforms); err != nil {
		return nil, err
	}

	// Atualizar plataformas publicadas
	if err := post.UpdatePublishedPlatforms(input.Platforms); err != nil {
		return nil, err
	}

	// Atualizar status para publicado
	if err := post.UpdateStatus(valueobject.StatusPublished); err != nil {
		return nil, err
	}

	// Adicionar evento ao histórico
	historyEvent := valueobject.NewPostHistoryEvent("published", input.User, nil)
	post.AddHistoryEvent(historyEvent)

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}
