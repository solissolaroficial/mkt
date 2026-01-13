package mapper

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/application/usecase/gifts"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// GiftItemPayloadMapper converte entre DTOs e Use Case Inputs/Outputs
type GiftItemPayloadMapper struct{}

func NewGiftItemPayloadMapper() *GiftItemPayloadMapper {
	return &GiftItemPayloadMapper{}
}

// CreateRequestToInput converte CreateGiftItemRequest para CreateGiftItemInput
func (m *GiftItemPayloadMapper) CreateRequestToInput(req *request.CreateGiftItemRequest) gifts.CreateGiftItemInput {
	return gifts.CreateGiftItemInput{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}
}

// UpdateRequestToInput converte UpdateGiftItemRequest para UpdateGiftItemInput
func (m *GiftItemPayloadMapper) UpdateRequestToInput(idStr string, req *request.UpdateGiftItemRequest) (gifts.UpdateGiftItemInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return gifts.UpdateGiftItemInput{}, err
	}
	return gifts.UpdateGiftItemInput{
		ID:    id,
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}, nil
}

// GetRequestToInput converte ID string para GetGiftItemInput
func (m *GiftItemPayloadMapper) GetRequestToInput(idStr string) (gifts.GetGiftItemInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return gifts.GetGiftItemInput{}, err
	}
	return gifts.GetGiftItemInput{
		ID: id,
	}, nil
}

// DeleteRequestToInput converte ID string para DeleteGiftItemInput
func (m *GiftItemPayloadMapper) DeleteRequestToInput(idStr string) (gifts.DeleteGiftItemInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return gifts.DeleteGiftItemInput{}, err
	}
	return gifts.DeleteGiftItemInput{
		ID: id,
	}, nil
}

// ListQueryToInput converte ListGiftItemsQuery para ListGiftItemsInput
func (m *GiftItemPayloadMapper) ListQueryToInput(query *request.ListGiftItemsQuery) gifts.ListGiftItemsInput {
	return gifts.ListGiftItemsInput{
		Name:     query.Name,
		MinStock: query.MinStock,
		MaxStock: query.MaxStock,
		MinPrice: query.MinPrice,
		MaxPrice: query.MaxPrice,
		Page:     query.Page,
		Limit:    query.Limit,
	}
}

// OutputToResponse converte CreateGiftItemOutput para GiftItemResponse
func (m *GiftItemPayloadMapper) OutputToResponse(output *gifts.CreateGiftItemOutput) *response.GiftItemResponse {
	return &response.GiftItemResponse{
		ID:    output.ID.String(),
		Name:  output.Name,
		Stock: output.Stock,
		Price: output.Price,
	}
}

// GetOutputToResponse converte GetGiftItemOutput para GiftItemResponse
func (m *GiftItemPayloadMapper) GetOutputToResponse(output *gifts.GetGiftItemOutput) *response.GiftItemResponse {
	return &response.GiftItemResponse{
		ID:        output.ID.String(),
		Name:      output.Name,
		Stock:     output.Stock,
		Price:     output.Price,
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}
}

// UpdateOutputToResponse converte UpdateGiftItemOutput para GiftItemResponse
func (m *GiftItemPayloadMapper) UpdateOutputToResponse(output *gifts.UpdateGiftItemOutput) *response.GiftItemResponse {
	return &response.GiftItemResponse{
		ID:        output.ID.String(),
		Name:      output.Name,
		Stock:     output.Stock,
		Price:     output.Price,
		CreatedAt: output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}
}

// ListOutputsToResponses converte slice de ListGiftItemsOutput para slice de GiftItemResponse
func (m *GiftItemPayloadMapper) ListOutputsToResponses(outputs []*gifts.ListGiftItemsOutput) []response.GiftItemResponse {
	responses := make([]response.GiftItemResponse, len(outputs))
	for i, output := range outputs {
		responses[i] = response.GiftItemResponse{
			ID:        output.ID.String(),
			Name:      output.Name,
			Stock:     output.Stock,
			Price:     output.Price,
			CreatedAt: output.CreatedAt,
			UpdatedAt: output.UpdatedAt,
		}
	}
	return responses
}

// GiftTransactionPayloadMapper converte entre DTOs e Use Case Inputs/Outputs
type GiftTransactionPayloadMapper struct{}

func NewGiftTransactionPayloadMapper() *GiftTransactionPayloadMapper {
	return &GiftTransactionPayloadMapper{}
}

// CreateRequestToInput converte CreateGiftTransactionRequest para CreateGiftTransactionInput
func (m *GiftTransactionPayloadMapper) CreateRequestToInput(req *request.CreateGiftTransactionRequest) gifts.CreateGiftTransactionInput {
	return gifts.CreateGiftTransactionInput{
		ItemName:        req.ItemName,
		Quantity:        req.Quantity,
		TransactionType: req.TransactionType,
		Date:            req.Date,
		Time:            req.Time,
		Price:           req.Price,
		Representative:  req.Representative,
		Unit:            req.Unit,
	}
}

// UpdateRequestToInput converte UpdateGiftTransactionRequest para UpdateGiftTransactionInput
func (m *GiftTransactionPayloadMapper) UpdateRequestToInput(idStr string, req *request.UpdateGiftTransactionRequest) (gifts.UpdateGiftTransactionInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return gifts.UpdateGiftTransactionInput{}, err
	}
	return gifts.UpdateGiftTransactionInput{
		ID:              id,
		ItemName:        req.ItemName,
		Quantity:        req.Quantity,
		TransactionType: req.TransactionType,
		Date:            req.Date,
		Time:            req.Time,
		Price:           req.Price,
		Representative:  req.Representative,
		Unit:            req.Unit,
	}, nil
}

// GetRequestToInput converte ID string para GetGiftTransactionInput
func (m *GiftTransactionPayloadMapper) GetRequestToInput(idStr string) (gifts.GetGiftTransactionInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return gifts.GetGiftTransactionInput{}, err
	}
	return gifts.GetGiftTransactionInput{
		ID: id,
	}, nil
}

// DeleteRequestToInput converte ID string para DeleteGiftTransactionInput
func (m *GiftTransactionPayloadMapper) DeleteRequestToInput(idStr string) (gifts.DeleteGiftTransactionInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return gifts.DeleteGiftTransactionInput{}, err
	}
	return gifts.DeleteGiftTransactionInput{
		ID: id,
	}, nil
}

// ListQueryToInput converte ListGiftTransactionsQuery para ListGiftTransactionsInput
func (m *GiftTransactionPayloadMapper) ListQueryToInput(query *request.ListGiftTransactionsQuery) gifts.ListGiftTransactionsInput {
	return gifts.ListGiftTransactionsInput{
		ItemName:        query.ItemName,
		TransactionType: query.TransactionType,
		Representative:  query.Representative,
		StartDate:       query.StartDate,
		EndDate:         query.EndDate,
		Page:            query.Page,
		Limit:           query.Limit,
	}
}

// OutputToResponse converte CreateGiftTransactionOutput para GiftTransactionResponse
func (m *GiftTransactionPayloadMapper) OutputToResponse(output *gifts.CreateGiftTransactionOutput) *response.GiftTransactionResponse {
	return &response.GiftTransactionResponse{
		ID:              output.ID.String(),
		ItemName:        output.ItemName,
		Quantity:        output.Quantity,
		TransactionType: output.TransactionType,
		Date:            output.Date,
		Time:            output.Time,
		Price:           output.Price,
		Representative:  output.Representative,
		Unit:            output.Unit,
		CreatedAt:       "",
		UpdatedAt:       "",
	}
}

// GetOutputToResponse converte GetGiftTransactionOutput para GiftTransactionResponse
func (m *GiftTransactionPayloadMapper) GetOutputToResponse(output *gifts.GetGiftTransactionOutput) *response.GiftTransactionResponse {
	return &response.GiftTransactionResponse{
		ID:              output.ID.String(),
		ItemName:        output.ItemName,
		Quantity:        output.Quantity,
		TransactionType: output.TransactionType,
		Date:            output.Date,
		Time:            output.Time,
		Price:           output.Price,
		Representative:  output.Representative,
		Unit:            output.Unit,
		CreatedAt:       output.CreatedAt,
		UpdatedAt:       output.UpdatedAt,
	}
}

// UpdateOutputToResponse converte UpdateGiftTransactionOutput para GiftTransactionResponse
func (m *GiftTransactionPayloadMapper) UpdateOutputToResponse(output *gifts.UpdateGiftTransactionOutput) *response.GiftTransactionResponse {
	return &response.GiftTransactionResponse{
		ID:              output.ID.String(),
		ItemName:        output.ItemName,
		Quantity:        output.Quantity,
		TransactionType: output.TransactionType,
		Date:            output.Date,
		Time:            output.Time,
		Price:           output.Price,
		Representative:  output.Representative,
		Unit:            output.Unit,
		CreatedAt:       output.CreatedAt,
		UpdatedAt:       output.UpdatedAt,
	}
}

// ListOutputsToResponses converte slice de ListGiftTransactionsOutput para slice de GiftTransactionResponse
func (m *GiftTransactionPayloadMapper) ListOutputsToResponses(outputs []*gifts.ListGiftTransactionsOutput) []response.GiftTransactionResponse {
	responses := make([]response.GiftTransactionResponse, len(outputs))
	for i, output := range outputs {
		responses[i] = response.GiftTransactionResponse{
			ID:              output.ID.String(),
			ItemName:        output.ItemName,
			Quantity:        output.Quantity,
			TransactionType: output.TransactionType,
			Date:            output.Date,
			Time:            output.Time,
			Price:           output.Price,
			Representative:  output.Representative,
			Unit:            output.Unit,
		}
	}
	return responses
}
