package tasks

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListCommentsUseCase lista todos os comentários de uma tarefa
type ListCommentsUseCase struct {
	commentGateway gateway.CommentGateway
}

// NewListCommentsUseCase cria um novo ListCommentsUseCase
func NewListCommentsUseCase(commentGateway gateway.CommentGateway) *ListCommentsUseCase {
	return &ListCommentsUseCase{
		commentGateway: commentGateway,
	}
}

// Execute lista todos os comentários de uma tarefa
func (uc *ListCommentsUseCase) Execute(taskID string) ([]*entity.Comment, error) {
	// Parse task ID
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, err
	}

	// Default pagination (no limit, no offset)
	pagination := valueobject.NewPagination(0, 0)

	// Default sort order (timestamp descending)
	sortOrder, err := valueobject.NewSortOrder("timestamp", valueobject.SortDirectionDesc)
	if err != nil {
		return nil, err
	}

	// Find comments by task ID
	comments, _, err := uc.commentGateway.FindByTaskID(taskUUID, &pagination, []*valueobject.SortOrder{sortOrder})
	if err != nil {
		return nil, err
	}

	return comments, nil
}
