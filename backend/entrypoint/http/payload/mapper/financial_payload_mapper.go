package mapper

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/application/usecase/financial/accountpayable"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// AccountPayablePayloadMapper converte entre DTOs e Use Case Inputs/Outputs
type AccountPayablePayloadMapper struct{}

func NewAccountPayablePayloadMapper() *AccountPayablePayloadMapper {
	return &AccountPayablePayloadMapper{}
}

// CreateRequestToInput converte CreateAccountPayableRequest para CreateAccountPayableInput
func (m *AccountPayablePayloadMapper) CreateRequestToInput(req *request.CreateAccountPayableRequest) accountpayable.CreateAccountPayableInput {
	return accountpayable.CreateAccountPayableInput{
		Supplier:    req.Supplier,
		Description: req.Description,
		Amount:      req.Amount,
		DueDate:     req.DueDate,
		Recurrence:  req.Recurrence,
	}
}

// UpdateRequestToInput converte UpdateAccountPayableRequest para UpdateAccountPayableInput
func (m *AccountPayablePayloadMapper) UpdateRequestToInput(idStr string, req *request.UpdateAccountPayableRequest) (accountpayable.UpdateAccountPayableInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return accountpayable.UpdateAccountPayableInput{}, err
	}
	return accountpayable.UpdateAccountPayableInput{
		ID:          id,
		Supplier:    req.Supplier,
		Description: req.Description,
		Amount:      req.Amount,
		DueDate:     req.DueDate,
		Recurrence:  req.Recurrence,
	}, nil
}

// GetRequestToInput converte ID string para GetAccountPayableInput
func (m *AccountPayablePayloadMapper) GetRequestToInput(idStr string) (accountpayable.GetAccountPayableInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return accountpayable.GetAccountPayableInput{}, err
	}
	return accountpayable.GetAccountPayableInput{
		ID: id,
	}, nil
}

// DeleteRequestToInput converte ID string para DeleteAccountPayableInput
func (m *AccountPayablePayloadMapper) DeleteRequestToInput(idStr string) (accountpayable.DeleteAccountPayableInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return accountpayable.DeleteAccountPayableInput{}, err
	}
	return accountpayable.DeleteAccountPayableInput{
		ID: id,
	}, nil
}

// ToggleNFRequestToInput converte ID string para ToggleNFInput
func (m *AccountPayablePayloadMapper) ToggleNFRequestToInput(idStr string) (accountpayable.ToggleNFInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return accountpayable.ToggleNFInput{}, err
	}
	return accountpayable.ToggleNFInput{
		ID: id,
	}, nil
}

// ToggleBoletoRequestToInput converte ID string para ToggleBoletoInput
func (m *AccountPayablePayloadMapper) ToggleBoletoRequestToInput(idStr string) (accountpayable.ToggleBoletoInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return accountpayable.ToggleBoletoInput{}, err
	}
	return accountpayable.ToggleBoletoInput{
		ID: id,
	}, nil
}

// SendToFinanceRequestToInput converte ID string para SendToFinanceInput
func (m *AccountPayablePayloadMapper) SendToFinanceRequestToInput(idStr string) (accountpayable.SendToFinanceInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return accountpayable.SendToFinanceInput{}, err
	}
	return accountpayable.SendToFinanceInput{
		ID: id,
	}, nil
}

// ListQueryToInput converte ListAccountsPayableQuery para ListAccountsPayableInput
func (m *AccountPayablePayloadMapper) ListQueryToInput(query *request.ListAccountsPayableQuery) accountpayable.ListAccountsPayableInput {
	return accountpayable.ListAccountsPayableInput{
		Supplier:   query.Supplier,
		MinAmount:  query.MinAmount,
		MaxAmount:  query.MaxAmount,
		Status:     query.Status,
		Recurrence: query.Recurrence,
		StartDate:  query.StartDate,
		EndDate:    query.EndDate,
		Page:       query.Page,
		Limit:      query.Limit,
		SortBy:     query.SortBy,
		SortOrder:  query.SortOrder,
	}
}

// OutputToResponse converte CreateAccountPayableOutput para AccountPayableResponse
func (m *AccountPayablePayloadMapper) OutputToResponse(output *accountpayable.CreateAccountPayableOutput) *response.AccountPayableResponse {
	return &response.AccountPayableResponse{
		ID:            output.ID.String(),
		Supplier:      output.Supplier,
		Description:   output.Description,
		Amount:        output.Amount,
		DueDate:       output.DueDate,
		NFArrived:     output.NFArrived,
		BoletoArrived: output.BoletoArrived,
		Status:        output.Status,
		Recurrence:    output.Recurrence,
		CreatedAt:     output.CreatedAt,
		UpdatedAt:     output.UpdatedAt,
	}
}

// GetOutputToResponse converte GetAccountPayableOutput para AccountPayableResponse
func (m *AccountPayablePayloadMapper) GetOutputToResponse(output *accountpayable.GetAccountPayableOutput) *response.AccountPayableResponse {
	return &response.AccountPayableResponse{
		ID:            output.ID.String(),
		Supplier:      output.Supplier,
		Description:   output.Description,
		Amount:        output.Amount,
		DueDate:       output.DueDate,
		NFArrived:     output.NFArrived,
		BoletoArrived: output.BoletoArrived,
		Status:        output.Status,
		Recurrence:    output.Recurrence,
		CreatedAt:     output.CreatedAt,
		UpdatedAt:     output.UpdatedAt,
	}
}

// UpdateOutputToResponse converte UpdateAccountPayableOutput para AccountPayableResponse
func (m *AccountPayablePayloadMapper) UpdateOutputToResponse(output *accountpayable.UpdateAccountPayableOutput) *response.AccountPayableResponse {
	return &response.AccountPayableResponse{
		ID:            output.ID.String(),
		Supplier:      output.Supplier,
		Description:   output.Description,
		Amount:        output.Amount,
		DueDate:       output.DueDate,
		NFArrived:     output.NFArrived,
		BoletoArrived: output.BoletoArrived,
		Status:        output.Status,
		Recurrence:    output.Recurrence,
		CreatedAt:     output.CreatedAt,
		UpdatedAt:     output.UpdatedAt,
	}
}

// ToggleNFOutputToResponse converte ToggleNFOutput para AccountPayableResponse
func (m *AccountPayablePayloadMapper) ToggleNFOutputToResponse(output *accountpayable.ToggleNFOutput) *response.AccountPayableResponse {
	return &response.AccountPayableResponse{
		ID:            output.ID.String(),
		Supplier:      output.Supplier,
		Description:   output.Description,
		Amount:        output.Amount,
		DueDate:       output.DueDate,
		NFArrived:     output.NFArrived,
		BoletoArrived: output.BoletoArrived,
		Status:        output.Status,
		Recurrence:    output.Recurrence,
		CreatedAt:     output.CreatedAt,
		UpdatedAt:     output.UpdatedAt,
	}
}

// ToggleBoletoOutputToResponse converte ToggleBoletoOutput para AccountPayableResponse
func (m *AccountPayablePayloadMapper) ToggleBoletoOutputToResponse(output *accountpayable.ToggleBoletoOutput) *response.AccountPayableResponse {
	return &response.AccountPayableResponse{
		ID:            output.ID.String(),
		Supplier:      output.Supplier,
		Description:   output.Description,
		Amount:        output.Amount,
		DueDate:       output.DueDate,
		NFArrived:     output.NFArrived,
		BoletoArrived: output.BoletoArrived,
		Status:        output.Status,
		Recurrence:    output.Recurrence,
		CreatedAt:     output.CreatedAt,
		UpdatedAt:     output.UpdatedAt,
	}
}

// SendToFinanceOutputToResponse converte SendToFinanceOutput para AccountPayableResponse
func (m *AccountPayablePayloadMapper) SendToFinanceOutputToResponse(output *accountpayable.SendToFinanceOutput) *response.AccountPayableResponse {
	return &response.AccountPayableResponse{
		ID:            output.ID.String(),
		Supplier:      output.Supplier,
		Description:   output.Description,
		Amount:        output.Amount,
		DueDate:       output.DueDate,
		NFArrived:     output.NFArrived,
		BoletoArrived: output.BoletoArrived,
		Status:        output.Status,
		Recurrence:    output.Recurrence,
		CreatedAt:     output.CreatedAt,
		UpdatedAt:     output.UpdatedAt,
	}
}

// ListOutputsToResponses converte slice de ListAccountsPayableOutput para slice de AccountPayableResponse
func (m *AccountPayablePayloadMapper) ListOutputsToResponses(outputs []*accountpayable.ListAccountsPayableOutput) []response.AccountPayableResponse {
	responses := make([]response.AccountPayableResponse, len(outputs))
	for i, output := range outputs {
		responses[i] = response.AccountPayableResponse{
			ID:            output.ID.String(),
			Supplier:      output.Supplier,
			Description:   output.Description,
			Amount:        output.Amount,
			DueDate:       output.DueDate,
			NFArrived:     output.NFArrived,
			BoletoArrived: output.BoletoArrived,
			Status:        output.Status,
			Recurrence:    output.Recurrence,
			CreatedAt:     output.CreatedAt,
			UpdatedAt:     output.UpdatedAt,
		}
	}
	return responses
}
