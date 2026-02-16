package flows

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	flowerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateFlow handles the business logic for creating a new flow
type CreateFlow struct {
	flowGateway gateway.FlowGateway
}

// NewCreateFlow creates a new CreateFlow use case
func NewCreateFlow(flowGateway gateway.FlowGateway) *CreateFlow {
	return &CreateFlow{
		flowGateway: flowGateway,
	}
}

// Execute creates a new flow
func (uc *CreateFlow) Execute(ctx context.Context, name string, description *string, color *string, sortOrder int) (*entity.Flow, error) {
	// Validate name
	if name == "" {
		return nil, &flowerrors.FlowEmptyNameError{}
	}

	// Create flow
	flow, err := entity.NewFlow(name, description, color, sortOrder)
	if err != nil {
		return nil, err
	}

	// Save to database
	createdFlow, err := uc.flowGateway.Create(flow)
	if err != nil {
		return nil, err
	}

	return createdFlow, nil
}
