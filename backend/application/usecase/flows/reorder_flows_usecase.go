package flows

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// ReorderFlows handles the business logic for reordering flows
type ReorderFlows struct {
	flowGateway gateway.FlowGateway
}

// NewReorderFlows creates a new ReorderFlows use case
func NewReorderFlows(flowGateway gateway.FlowGateway) *ReorderFlows {
	return &ReorderFlows{
		flowGateway: flowGateway,
	}
}

// Execute reorders flows by updating their sort_order
func (uc *ReorderFlows) Execute(ctx context.Context, flowIDs []string) error {
	return uc.flowGateway.Reorder(flowIDs)
}
