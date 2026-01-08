package pdv

import (
	"context"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// UpdatePdvPostUseCase atualiza um post de PDV existente
type UpdatePdvPostUseCase struct {
	gateway gateway.PdvPostGateway
}

// UpdatePdvPostInput representa os dados de entrada para atualizar um post de PDV
type UpdatePdvPostInput struct {
	ID       string
	RepName  *string
	PdvName  *string
	PostDate *string
	Platform *string
	Link     *string
	ProofUrl *string
}

// NewUpdatePdvPost cria uma nova instância do UpdatePdvPostUseCase
func NewUpdatePdvPost(gateway gateway.PdvPostGateway) *UpdatePdvPostUseCase {
	return &UpdatePdvPostUseCase{gateway: gateway}
}

// Execute atualiza um post de PDV existente
func (uc *UpdatePdvPostUseCase) Execute(ctx context.Context, input UpdatePdvPostInput) (*entity.PdvPost, error) {
	// Buscar post existente
	post, err := uc.gateway.FindByID(ctx, input.ID)
	if err != nil {
		return nil, domainErrors.ErrPdvPostNotFound
	}

	// Atualizar campos se fornecidos
	if input.RepName != nil {
		if err := post.UpdateRepName(*input.RepName); err != nil {
			return nil, err
		}
	}

	if input.PdvName != nil {
		if err := post.UpdatePdvName(*input.PdvName); err != nil {
			return nil, err
		}
	}

	if input.PostDate != nil {
		date, err := time.Parse("2006-01-02", *input.PostDate)
		if err != nil {
			return nil, err
		}
		postDate, err := valueobject.NewPdvPostDate(date)
		if err != nil {
			return nil, err
		}
		if err := post.UpdatePostDate(postDate); err != nil {
			return nil, err
		}
	}

	if input.Platform != nil {
		if err := post.UpdatePlatform(*input.Platform); err != nil {
			return nil, err
		}
	}

	if input.Link != nil {
		if err := post.UpdateLink(input.Link); err != nil {
			return nil, err
		}
	}

	if input.ProofUrl != nil {
		if err := post.UpdateProofUrl(input.ProofUrl); err != nil {
			return nil, err
		}
	}

	// Salvar via gateway
	if err := uc.gateway.Update(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}
