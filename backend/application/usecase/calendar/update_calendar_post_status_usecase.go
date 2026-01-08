package calendar

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type UpdateCalendarPostStatusUseCase struct {
	gateway gateway.CalendarPostGateway
}

type UpdateStatusInput struct {
	PostID    string
	NewStatus string
	User      string
	Comment   *string
}

func NewUpdateCalendarPostStatus(gateway gateway.CalendarPostGateway) *UpdateCalendarPostStatusUseCase {
	return &UpdateCalendarPostStatusUseCase{gateway: gateway}
}

func (uc *UpdateCalendarPostStatusUseCase) Execute(ctx context.Context, input UpdateStatusInput) (*entity.CalendarPost, error) {
	// Buscar post existente
	post, err := uc.gateway.FindByID(ctx, input.PostID)
	if err != nil {
		return nil, domainErrors.ErrCalendarPostNotFound
	}

	// Validar e converter status
	newStatus, err := valueobject.NewPostStatus(input.NewStatus)
	if err != nil {
		return nil, err
	}

	// Atualizar status usando método da entity
	if err := post.UpdateStatus(newStatus); err != nil {
		return nil, err
	}

	// Adicionar evento ao histórico
	action := uc.mapStatusToAction(newStatus)
	historyEvent := valueobject.NewPostHistoryEvent(action, input.User, input.Comment)
	post.AddHistoryEvent(historyEvent)

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}

func (uc *UpdateCalendarPostStatusUseCase) mapStatusToAction(status valueobject.PostStatus) string {
	actionMap := map[valueobject.PostStatus]string{
		valueobject.StatusInProgress: "status_change",
		valueobject.StatusReview:     "status_change",
		valueobject.StatusAdjust:     "adjust_request",
		valueobject.StatusApproved:   "approved",
		valueobject.StatusPublished:  "published",
	}
	return actionMap[status]
}
