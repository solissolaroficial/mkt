package flows

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteFlow handles the business logic for deleting a flow
type DeleteFlow struct {
	flowGateway gateway.FlowGateway
}

// NewDeleteFlow creates a new DeleteFlow use case
func NewDeleteFlow(flowGateway gateway.FlowGateway) *DeleteFlow {
	return &DeleteFlow{
		flowGateway: flowGateway,
	}
}

// Execute deletes a flow
func (uc *DeleteFlow) Execute(ctx context.Context, id string) error {
	// Check if flow exists
	_, err := uc.flowGateway.GetByUUID(id)
	if err != nil {
		return err
	}

	// Delete flow
	return uc.flowGateway.Delete(id)
}
