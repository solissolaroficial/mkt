package pdv

import (
	"context"

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
func (uc *UpdatePdvPostStatusUseCase) Execute(ctx context.Context, input UpdateStatusInput) error {
	// Buscar post existente
	post, err := uc.gateway.FindByID(ctx, input.PostID)
	if err != nil {
		return domainErrors.ErrPdvPostNotFound
	}

	// Validar e converter status
	newStatus, err := valueobject.NewPdvStatus(input.NewStatus)
	if err != nil {
		return err
	}

	// Atualizar status usando método da entity
	if err := post.UpdateStatus(newStatus); err != nil {
		return err
	}

	// Salvar via gateway
	return uc.gateway.Update(ctx, post)
}
