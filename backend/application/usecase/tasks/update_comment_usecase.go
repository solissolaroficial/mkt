package tasks

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateCommentUseCase atualiza um comentário existente
type UpdateCommentUseCase struct {
	commentGateway gateway.CommentGateway
}

// NewUpdateCommentUseCase cria um novo UpdateCommentUseCase
func NewUpdateCommentUseCase(commentGateway gateway.CommentGateway) *UpdateCommentUseCase {
	return &UpdateCommentUseCase{
		commentGateway: commentGateway,
	}
}

// Execute atualiza um comentário
func (uc *UpdateCommentUseCase) Execute(id, content string) error {
	// Parse ID
	commentID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid comment ID: %w", err)
	}

	// Find existing comment
	comment, err := uc.commentGateway.FindByID(commentID)
	if err != nil {
		return err
	}

	// Update comment content
	if err := comment.Validate(); err != nil {
		return err
	}

	// Update content via Reconstruct
	updatedComment := entity.ReconstructComment(
		comment.ID(),
		comment.TaskID(),
		comment.UserID(),
		content,
		time.Now(), // timestamp
	)

	// Save comment
	return uc.commentGateway.Update(updatedComment)
}
