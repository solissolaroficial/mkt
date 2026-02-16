package flows

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateFlow handles the business logic for updating a flow
type UpdateFlow struct {
	flowGateway gateway.FlowGateway
}

// NewUpdateFlow creates a new UpdateFlow use case
func NewUpdateFlow(flowGateway gateway.FlowGateway) *UpdateFlow {
	return &UpdateFlow{
		flowGateway: flowGateway,
	}
}

// Execute updates a flow
func (uc *UpdateFlow) Execute(ctx context.Context, id string, name string, description *string, color *string, sortOrder int) (*entity.Flow, error) {
	// Get existing flow
	flow, err := uc.flowGateway.GetByUUID(id)
	if err != nil {
		return nil, err
	}

	// Update name
	if name != "" {
		if err := flow.UpdateName(name); err != nil {
			return nil, err
		}
	}

	// Update description
	flow.UpdateDescription(description)

	// Update color
	flow.UpdateColor(color)

	// Update sort order
	if sortOrder >= 0 {
		if err := flow.UpdateSortOrder(sortOrder); err != nil {
			return nil, err
		}
	}

	// Save to database
	updatedFlow, err := uc.flowGateway.Update(flow)
	if err != nil {
		return nil, err
	}

	return updatedFlow, nil
}
