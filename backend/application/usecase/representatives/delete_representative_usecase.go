package representatives

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteRepresentativeInput defines input data for deleting a representative
type DeleteRepresentativeInput struct {
	ID uuid.UUID
}

type DeleteRepresentativeUseCase struct {
	representativeGateway gateway.RepresentativeGateway
}

func NewDeleteRepresentativeUseCase(representativeGateway gateway.RepresentativeGateway) *DeleteRepresentativeUseCase {
	return &DeleteRepresentativeUseCase{
		representativeGateway: representativeGateway,
	}
}

func (uc *DeleteRepresentativeUseCase) Execute(ctx context.Context, input DeleteRepresentativeInput) error {
	// Check if representative exists
	exists, err := uc.representativeGateway.ExistsByID(input.ID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.ErrRepresentativeNotFound
	}

	// Delete representative (soft delete)
	if err := uc.representativeGateway.Delete(input.ID); err != nil {
		return err
	}

	return nil
}
