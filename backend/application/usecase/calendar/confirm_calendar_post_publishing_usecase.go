package calendar

import (
	"context"

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

func (uc *ConfirmCalendarPostPublishingUseCase) Execute(ctx context.Context, input ConfirmPublishingInput) error {
	// Buscar post existente
	post, err := uc.gateway.FindByID(ctx, input.PostID)
	if err != nil {
		return domainErrors.ErrCalendarPostNotFound
	}

	// Validar plataformas
	if err := valueobject.ValidatePlatforms(input.Platforms); err != nil {
		return err
	}

	// Atualizar plataformas publicadas
	if err := post.UpdatePublishedPlatforms(input.Platforms); err != nil {
		return err
	}

	// Atualizar status para publicado
	if err := post.UpdateStatus(valueobject.StatusPublished); err != nil {
		return err
	}

	// Adicionar evento ao histórico
	historyEvent := valueobject.NewPostHistoryEvent("published", input.User, nil)
	post.AddHistoryEvent(historyEvent)

	// Salvar via gateway
	return uc.gateway.Update(ctx, post)
}
