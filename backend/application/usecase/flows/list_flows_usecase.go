package flows

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// ListFlows handles the business logic for listing all flows
type ListFlows struct {
	flowGateway gateway.FlowGateway
}

// NewListFlows creates a new ListFlows use case
func NewListFlows(flowGateway gateway.FlowGateway) *ListFlows {
	return &ListFlows{
		flowGateway: flowGateway,
	}
}

// Execute lists all flows
func (uc *ListFlows) Execute(ctx context.Context) ([]*entity.Flow, error) {
	return uc.flowGateway.GetAll()
}
