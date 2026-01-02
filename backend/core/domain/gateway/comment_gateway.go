package gateway

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"

	"github.com/google/uuid"
)

// CommentGateway define a interface para persistência de comentários
type CommentGateway interface {
	// Create cria um novo comentário
	Create(comment *entity.Comment) error

	// Update atualiza um comentário existente
	Update(comment *entity.Comment) error

	// Delete deleta um comentário
	Delete(id uuid.UUID) error

	// FindByID busca um comentário pelo ID
	FindByID(id uuid.UUID) (*entity.Comment, error)

	// FindAll busca todos os comentários com paginação
	FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Comment, int64, error)

	// FindByTaskID busca comentários por task ID
	FindByTaskID(taskID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Comment, int64, error)

	// FindByUserID busca comentários por user ID
	FindByUserID(userID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Comment, int64, error)
}
