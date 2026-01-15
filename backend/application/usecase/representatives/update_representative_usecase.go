package representatives

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateRepresentativeInput defines input data for updating a representative
type UpdateRepresentativeInput struct {
	ID        uuid.UUID
	Name      *string
	Email     *string
	Phone     *string
	Company   *string
	Region    *string
	City      *string
	Attendant *string
	Active    *bool
}

// UpdateRepresentativeOutput defines output data after updating a representative
type UpdateRepresentativeOutput struct {
	UUID      uuid.UUID
	Code      int
	Name      string
	Email     string
	Phone     string
	Company   string
	Region    string
	City      string
	Attendant string
	Active    bool
	CreatedAt string
	UpdatedAt string
}

type UpdateRepresentativeUseCase struct {
	representativeGateway gateway.RepresentativeGateway
}

func NewUpdateRepresentativeUseCase(representativeGateway gateway.RepresentativeGateway) *UpdateRepresentativeUseCase {
	return &UpdateRepresentativeUseCase{
		representativeGateway: representativeGateway,
	}
}

func (uc *UpdateRepresentativeUseCase) Execute(ctx context.Context, input UpdateRepresentativeInput) (*UpdateRepresentativeOutput, error) {
	// Find existing representative
	representative, err := uc.representativeGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if representative == nil {
		return nil, errors.ErrRepresentativeNotFound
	}

	// Update fields if provided
	if input.Name != nil {
		if err := representative.UpdateName(*input.Name); err != nil {
			return nil, err
		}
	}
	if input.Email != nil {
		if err := representative.UpdateEmail(*input.Email); err != nil {
			return nil, err
		}
	}
	if input.Phone != nil {
		if err := representative.UpdatePhone(*input.Phone); err != nil {
			return nil, err
		}
	}
	if input.Company != nil {
		if err := representative.UpdateCompany(*input.Company); err != nil {
			return nil, err
		}
	}
	if input.Region != nil {
		if err := representative.UpdateRegion(*input.Region); err != nil {
			return nil, err
		}
	}
	if input.City != nil {
		if err := representative.UpdateCity(*input.City); err != nil {
			return nil, err
		}
	}
	if input.Attendant != nil {
		if err := representative.UpdateAttendant(*input.Attendant); err != nil {
			return nil, err
		}
	}
	if input.Active != nil {
		if *input.Active {
			representative.Activate()
		} else {
			representative.Deactivate()
		}
	}

	// Save to database
	if err := uc.representativeGateway.Update(representative); err != nil {
		return nil, err
	}

	return &UpdateRepresentativeOutput{
		UUID:      representative.UUID(),
		Code:      representative.Code().Value(),
		Name:      representative.Name(),
		Email:     representative.Email(),
		Phone:     representative.Phone(),
		Company:   representative.Company(),
		Region:    representative.Region(),
		City:      representative.City(),
		Attendant: representative.Attendant(),
		Active:    representative.Active(),
		CreatedAt: representative.CreatedAt().Format(time.RFC3339),
		UpdatedAt: representative.UpdatedAt().Format(time.RFC3339),
	}, nil
}
