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
	comments, _, err := uc.ExecuteWithPagination(taskID, 1, 10)
	return comments, err
}

// ExecuteWithPagination lista comentários com paginação
func (uc *ListCommentsUseCase) ExecuteWithPagination(taskID string, page, limit int) ([]*entity.Comment, int64, error) {
	// Parse task ID
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, 0, err
	}

	// Validar e calcular paginação
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // valor padrão
	}

	offset := (page - 1) * limit

	// Usar paginação correta
	pagination := valueobject.NewPagination(offset, limit)

	// Default sort order (timestamp descending)
	sortOrder, err := valueobject.NewSortOrder("timestamp", valueobject.SortDirectionDesc)
	if err != nil {
		return nil, 0, err
	}

	// Find comments by task ID
	comments, total, err := uc.commentGateway.FindByTaskID(taskUUID, &pagination, []*valueobject.SortOrder{sortOrder})
	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}
