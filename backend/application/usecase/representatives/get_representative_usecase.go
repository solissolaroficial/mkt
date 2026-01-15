package representatives

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetRepresentativeInput defines input data for getting a representative
type GetRepresentativeInput struct {
	ID uuid.UUID
}

// GetRepresentativeOutput defines output data for getting a representative
type GetRepresentativeOutput struct {
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
	DeletedAt *string
}

type GetRepresentativeUseCase struct {
	representativeGateway gateway.RepresentativeGateway
}

func NewGetRepresentativeUseCase(representativeGateway gateway.RepresentativeGateway) *GetRepresentativeUseCase {
	return &GetRepresentativeUseCase{
		representativeGateway: representativeGateway,
	}
}

func (uc *GetRepresentativeUseCase) Execute(ctx context.Context, input GetRepresentativeInput) (*GetRepresentativeOutput, error) {
	representative, err := uc.representativeGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if representative == nil {
		return nil, errors.ErrRepresentativeNotFound
	}

	deletedAt := ""
	if representative.DeletedAt() != nil {
		deletedAt = representative.DeletedAt().Format(time.RFC3339)
	}

	return &GetRepresentativeOutput{
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
		DeletedAt: &deletedAt,
	}, nil
}
