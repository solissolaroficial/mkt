package social

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type CreateSocialPostInput struct {
	BrandID         uuid.UUID
	Platform        string
	PostDate        time.Time
	PostTime        *time.Time
	Likes           int
	Comments        int
	Shares          *int
	PostType        string
	Caption         *string
	FollowersAtPost *int
}

type CreateSocialPostOutput struct {
	ID              uuid.UUID
	BrandName       string
	Platform        string
	PostDate        time.Time
	PostTime        *time.Time
	Likes           int
	Comments        int
	Shares          *int
	PostType        string
	Caption         *string
	FollowersAtPost *int
	CreatedAt       time.Time
}

type CreateSocialPostUseCase struct {
	postGateway                    gateway.SocialPostGateway
	brandGateway                   gateway.BrandGateway
	recalculateAggregationsUseCase *RecalculateDailyAggregationsUseCase
}

func NewCreateSocialPostUseCase(
	postGateway gateway.SocialPostGateway,
	brandGateway gateway.BrandGateway,
	recalculateAggregationsUseCase *RecalculateDailyAggregationsUseCase,
) *CreateSocialPostUseCase {
	return &CreateSocialPostUseCase{
		postGateway:                    postGateway,
		brandGateway:                   brandGateway,
		recalculateAggregationsUseCase: recalculateAggregationsUseCase,
	}
}

func (uc *CreateSocialPostUseCase) Execute(input CreateSocialPostInput) (*CreateSocialPostOutput, error) {
	// Validar se o brand existe
	brand, err := uc.brandGateway.FindByID(context.Background(), input.BrandID)
	if err != nil {
		return nil, fmt.Errorf("brand not found")
	}

	// Criar entidade com o nome da marca
	post, err := entity.NewSocialPost(
		brand.Name(),
		input.Platform,
		input.PostDate,
		input.PostTime,
		input.Likes,
		input.Comments,
		input.Shares,
		input.PostType,
		input.Caption,
		input.FollowersAtPost,
	)
	if err != nil {
		return nil, err
	}

	// Atualizar BrandID com o ID correto
	post.UpdateBrandID(input.BrandID)

	// Salvar no banco
	if err := uc.postGateway.Create(post); err != nil {
		return nil, err
	}

	// Recalcular agregações diárias
	_, err = uc.recalculateAggregationsUseCase.Execute(input.BrandID, input.PostDate)
	if err != nil {
		// Log error mas não falhar a criação do post
		// Em produção, isso deveria ser assíncrono
		return nil, fmt.Errorf("failed to recalculate daily aggregations: %w", err)
	}

	// Retornar output
	return &CreateSocialPostOutput{
		ID:              post.ID(),
		BrandName:       post.BrandName().String(),
		Platform:        post.Platform().String(),
		PostDate:        post.PostDate(),
		PostTime:        post.PostTime(),
		Likes:           post.Likes(),
		Comments:        post.Comments(),
		Shares:          post.Shares(),
		PostType:        post.PostType().String(),
		Caption:         post.Caption(),
		FollowersAtPost: post.FollowersAtPost(),
		CreatedAt:       post.CreatedAt(),
	}, nil
}
