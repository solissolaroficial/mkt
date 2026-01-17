package pdv

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// CreatePdvPostUseCase cria um novo post de PDV
type CreatePdvPostUseCase struct {
	gateway gateway.PdvPostGateway
}

// CreatePdvPostInput representa os dados de entrada para criar um post de PDV
type CreatePdvPostInput struct {
	RepresentativeUUID string
	PdvName            string
	PostDate           string
	Platform           string
	Link               *string
	ProofUrl           *string
}

// NewCreatePdvPost cria uma nova instância do CreatePdvPostUseCase
func NewCreatePdvPost(gateway gateway.PdvPostGateway) *CreatePdvPostUseCase {
	return &CreatePdvPostUseCase{
		gateway: gateway,
	}
}

// Execute cria um novo post de PDV
func (uc *CreatePdvPostUseCase) Execute(ctx context.Context, input CreatePdvPostInput) (*entity.PdvPost, error) {
	// Validar e criar value objects
	date, err := time.Parse("2006-01-02", input.PostDate)
	if err != nil {
		return nil, errors.New("invalid post date format")
	}

	postDate, err := valueobject.NewPdvPostDate(date)
	if err != nil {
		return nil, err
	}

	// Parse representative UUID
	representativeUUID, err := uuid.Parse(input.RepresentativeUUID)
	if err != nil {
		return nil, err
	}

	// Validar plataforma
	if !valueobject.IsValidPlatform(input.Platform) {
		return nil, errors.New("invalid platform")
	}

	// Validar link se fornecido
	if input.Link != nil && *input.Link != "" {
		if len(*input.Link) > 500 {
			return nil, errors.New("link must be at most 500 characters")
		}
	}

	// Validar proofUrl se fornecido
	if input.ProofUrl != nil && *input.ProofUrl != "" {
		if len(*input.ProofUrl) > 500 {
			return nil, errors.New("proofUrl must be at most 500 characters")
		}
	}

	// Criar entidade
	post, err := entity.NewPdvPost(
		representativeUUID,
		input.PdvName,
		postDate,
		input.Platform,
		input.Link,
		input.ProofUrl,
	)
	if err != nil {
		return nil, err
	}

	// Salvar via gateway
	if err := uc.gateway.Save(ctx, post); err != nil {
		return nil, err
	}

	return post, nil
}
