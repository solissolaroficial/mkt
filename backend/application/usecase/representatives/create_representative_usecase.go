package representatives

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// CreateRepresentativeInput defines input data for creating a representative
type CreateRepresentativeInput struct {
	Code      int
	Name      string
	Email     string
	Phone     string
	Company   string
	Region    string
	City      string
	Attendant string
}

// CreateRepresentativeOutput defines output data after creating a representative
type CreateRepresentativeOutput struct {
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

type CreateRepresentativeUseCase struct {
	representativeGateway gateway.RepresentativeGateway
}

func NewCreateRepresentativeUseCase(representativeGateway gateway.RepresentativeGateway) *CreateRepresentativeUseCase {
	return &CreateRepresentativeUseCase{
		representativeGateway: representativeGateway,
	}
}

func (uc *CreateRepresentativeUseCase) Execute(ctx context.Context, input CreateRepresentativeInput) (*CreateRepresentativeOutput, error) {
	// Validate input data
	if input.Name == "" {
		return nil, errors.ErrRepresentativeNameRequired
	}
	if input.Email == "" {
		return nil, errors.ErrRepresentativeEmailRequired
	}
	if input.Company == "" {
		return nil, errors.ErrRepresentativeCompanyRequired
	}
	if input.Region == "" {
		return nil, errors.ErrRepresentativeRegionRequired
	}

	// Create value object for code
	code, err := valueobject.NewRepresentativeCode(input.Code)
	if err != nil {
		return nil, err
	}

	// Check if code already exists
	exists, err := uc.representativeGateway.ExistsByCode(input.Code)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.ErrRepresentativeAlreadyExists
	}

	// Create entity
	representative, err := entity.NewRepresentative(
		code,
		input.Name,
		input.Email,
		input.Phone,
		input.Company,
		input.Region,
		input.City,
		input.Attendant,
	)
	if err != nil {
		return nil, err
	}

	// Save to database
	if err := uc.representativeGateway.Create(representative); err != nil {
		return nil, err
	}

	return &CreateRepresentativeOutput{
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
