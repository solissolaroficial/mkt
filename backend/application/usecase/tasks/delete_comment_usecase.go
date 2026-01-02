package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteCommentUseCase deleta um comentário existente
type DeleteCommentUseCase struct {
	commentGateway gateway.CommentGateway
}

// NewDeleteCommentUseCase cria um novo DeleteCommentUseCase
func NewDeleteCommentUseCase(commentGateway gateway.CommentGateway) *DeleteCommentUseCase {
	return &DeleteCommentUseCase{
		commentGateway: commentGateway,
	}
}

// Execute deleta um comentário
func (uc *DeleteCommentUseCase) Execute(id string) error {
	// Parse ID
	commentID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid comment ID: %w", err)
	}

	// Delete comment
	return uc.commentGateway.Delete(commentID)
}
