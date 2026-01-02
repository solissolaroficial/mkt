package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetCommentUseCase busca um comentário por ID
type GetCommentUseCase struct {
	commentGateway gateway.CommentGateway
}

// NewGetCommentUseCase cria um novo GetCommentUseCase
func NewGetCommentUseCase(commentGateway gateway.CommentGateway) *GetCommentUseCase {
	return &GetCommentUseCase{
		commentGateway: commentGateway,
	}
}

// Execute busca um comentário por ID
func (uc *GetCommentUseCase) Execute(id string) (*entity.Comment, error) {
	// Parse ID
	commentID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid comment ID: %w", err)
	}

	// Find comment
	comment, err := uc.commentGateway.FindByID(commentID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}
