package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetSubtaskUseCase busca uma subtarefa pelo ID
type GetSubtaskUseCase struct {
	subtaskGateway gateway.SubtaskGateway
}

// NewGetSubtaskUseCase cria um novo GetSubtaskUseCase
func NewGetSubtaskUseCase(subtaskGateway gateway.SubtaskGateway) *GetSubtaskUseCase {
	return &GetSubtaskUseCase{
		subtaskGateway: subtaskGateway,
	}
}

// Execute busca uma subtarefa
func (uc *GetSubtaskUseCase) Execute(id string) (*entity.Subtask, error) {
	// Parse ID
	subtaskID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid subtask ID: %w", err)
	}

	// Find subtask
	return uc.subtaskGateway.FindByID(subtaskID)
}
