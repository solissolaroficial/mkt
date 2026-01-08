package pdv

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// UpdatePdvPostStatusUseCase atualiza o status de um post de PDV
type UpdatePdvPostStatusUseCase struct {
	gateway gateway.PdvPostGateway
}

// UpdateStatusInput representa os dados de entrada para atualizar o status
type UpdateStatusInput struct {
	PostID    string
	NewStatus string
	User      string
}

// NewUpdatePdvPostStatus cria uma nova instância do UpdatePdvPostStatusUseCase
func NewUpdatePdvPostStatus(gateway gateway.PdvPostGateway) *UpdatePdvPostStatusUseCase {
	return &UpdatePdvPostStatusUseCase{gateway: gateway}
}

// Execute atualiza o status de um post de PDV
func (uc *UpdatePdvPostStatusUseCase) Execute(ctx context.Context, input UpdateStatusInput) (*entity.PdvPost, error) {
	// Buscar post existente
	post, err := uc.gateway.FindByID(ctx, input.PostID)
	if err != nil {
		return nil, domainErrors.ErrPdvPostNotFound
	}

	// Validar e converter status
	newStatus, err := valueobject.NewPdvStatus(input.NewStatus)
	if err != nil {
		return nil, err
	}

	// Atualizar status usando método da entity
	if err := post.UpdateStatus(newStatus); err != nil {
		return nil, err
	}

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}
