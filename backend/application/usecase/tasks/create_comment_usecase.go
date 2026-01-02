package tasks

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateCommentUseCase cria um novo comentário
type CreateCommentUseCase struct {
	commentGateway gateway.CommentGateway
}

// NewCreateCommentUseCase cria um novo CreateCommentUseCase
func NewCreateCommentUseCase(commentGateway gateway.CommentGateway) *CreateCommentUseCase {
	return &CreateCommentUseCase{
		commentGateway: commentGateway,
	}
}

// Execute cria um novo comentário
func (uc *CreateCommentUseCase) Execute(comment *entity.Comment) error {
	if err := comment.Validate(); err != nil {
		return err
	}

	return uc.commentGateway.Create(comment)
}
