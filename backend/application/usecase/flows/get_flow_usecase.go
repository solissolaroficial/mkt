package flows

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetFlow handles the business logic for getting a flow
type GetFlow struct {
	flowGateway gateway.FlowGateway
}

// NewGetFlow creates a new GetFlow use case
func NewGetFlow(flowGateway gateway.FlowGateway) *GetFlow {
	return &GetFlow{
		flowGateway: flowGateway,
	}
}

// Execute gets a flow by ID
func (uc *GetFlow) Execute(ctx context.Context, id string) (*entity.Flow, error) {
	flow, err := uc.flowGateway.GetByUUID(id)
	if err != nil {
		return nil, err
	}

	return flow, nil
}
