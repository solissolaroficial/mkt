package mapper

import (
	"context"

	"github.com/seu-usuario/solis-backend/application/usecase/budget"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	budgetrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	budgetresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type BudgetPayloadMapper struct{}

func NewBudgetPayloadMapper() *BudgetPayloadMapper {
	return &BudgetPayloadMapper{}
}

// ToBudgetItemResponse converte uma entidade BudgetItem para BudgetItemResponse
func (m *BudgetPayloadMapper) ToBudgetItemResponse(budget *entity.BudgetItem) *budgetresponse.BudgetItemResponse {
	if budget == nil {
		return nil
	}

	return &budgetresponse.BudgetItemResponse{
		UUID:         budget.ID(),
		ObjectUUID:   budget.ObjectUUID(),
		ObjectName:   budget.ObjectName(),
		GroupUUID:    budget.GroupUUID(),
		GroupName:    budget.GroupName(),
		Cod:          budget.Cod(),
		Desc:         budget.Desc(),
		Vals:         budget.Vals(),
		RealizedVals: budget.RealizedVals(),
		Year:         budget.Year(),
		CreatedAt:    budget.CreatedAt(),
		UpdatedAt:    budget.UpdatedAt(),
		DeletedAt:    budget.DeletedAt(),
	}
}

// ToBudgetItemResponseList converte uma lista de entidades para BudgetItemResponse
func (m *BudgetPayloadMapper) ToBudgetItemResponseList(budgets []*entity.BudgetItem) []*budgetresponse.BudgetItemResponse {
	if budgets == nil {
		return []*budgetresponse.BudgetItemResponse{}
	}

	responses := make([]*budgetresponse.BudgetItemResponse, 0, len(budgets))
	for _, b := range budgets {
		responses = append(responses, m.ToBudgetItemResponse(b))
	}

	return responses
}

// ToBudgetSummaryResponse converte um BudgetSummary do gateway para BudgetSummaryResponse
func (m *BudgetPayloadMapper) ToBudgetSummaryResponse(summary *gateway.BudgetSummary) *budgetresponse.BudgetSummaryResponse {
	if summary == nil {
		return nil
	}

	return &budgetresponse.BudgetSummaryResponse{
		ObjectUUID:    summary.ObjectUUID,
		ObjectName:    summary.ObjectName,
		GroupUUID:     summary.GroupUUID,
		GroupName:     summary.GroupName,
		TotalBudget:   summary.TotalBudget,
		TotalRealized: summary.TotalRealized,
		Variance:      summary.Variance,
	}
}

// ToBudgetSummaryResponseList converte uma lista de BudgetSummary para BudgetSummaryResponse
func (m *BudgetPayloadMapper) ToBudgetSummaryResponseList(summaries []*gateway.BudgetSummary) []*budgetresponse.BudgetSummaryResponse {
	if summaries == nil {
		return []*budgetresponse.BudgetSummaryResponse{}
	}

	responses := make([]*budgetresponse.BudgetSummaryResponse, 0, len(summaries))
	for _, summary := range summaries {
		responses = append(responses, m.ToBudgetSummaryResponse(summary))
	}

	return responses
}

// CreateRequestToInput converte CreateBudgetItemRequest para o input do use case
func (m *BudgetPayloadMapper) CreateRequestToInput(req *budgetrequest.CreateBudgetItemRequest) budget.CreateBudgetItemInput {
	return budget.CreateBudgetItemInput{
		ObjectUUID:   req.ObjectUUID,
		ObjectName:   req.ObjectName,
		GroupUUID:    req.GroupUUID,
		GroupName:    req.GroupName,
		Cod:          req.Cod,
		Desc:         req.Desc,
		Vals:         req.Vals,
		RealizedVals: req.RealizedVals,
		Year:         req.Year,
	}
}

// UpdateRequestToInput converte UpdateBudgetItemRequest para o input do use case
func (m *BudgetPayloadMapper) UpdateRequestToInput(id string, req *budgetrequest.UpdateBudgetItemRequest) (string, budget.UpdateBudgetItemInput) {
	// Converter string para *string para ObjectName e GroupName
	var objectName *string
	if req.ObjectName != "" {
		objectName = &req.ObjectName
	}

	var groupName *string
	if req.GroupName != "" {
		groupName = &req.GroupName
	}

	return id, budget.UpdateBudgetItemInput{
		ObjectUUID:   req.ObjectUUID,
		ObjectName:   objectName,
		GroupUUID:    req.GroupUUID,
		GroupName:    groupName,
		Cod:          req.Cod,
		Desc:         req.Desc,
		Vals:         req.Vals,
		RealizedVals: req.RealizedVals,
		Year:         req.Year,
	}
}

// BatchCreateRequestToInput converte BatchCreateBudgetItemsRequest para o input do use case
func (m *BudgetPayloadMapper) BatchCreateRequestToInput(req *budgetrequest.BatchCreateBudgetItemsRequest) budget.BatchCreateBudgetItemsInput {
	items := make([]budget.CreateBudgetItemInput, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, m.CreateRequestToInput(&item))
	}

	return budget.BatchCreateBudgetItemsInput{
		Items: items,
	}
}

// OutputToResponse converte uma entidade para BudgetItemResponse
func (m *BudgetPayloadMapper) OutputToResponse(budget *entity.BudgetItem) *budgetresponse.BudgetItemResponse {
	return m.ToBudgetItemResponse(budget)
}

// ListOutputsToResponses converte uma lista de entidades para BudgetItemResponse
func (m *BudgetPayloadMapper) ListOutputsToResponses(budgets []*entity.BudgetItem) []*budgetresponse.BudgetItemResponse {
	return m.ToBudgetItemResponseList(budgets)
}

// SummaryOutputsToResponses converte uma lista de BudgetSummary para BudgetSummaryResponse
func (m *BudgetPayloadMapper) SummaryOutputsToResponses(summaries []*gateway.BudgetSummary) []*budgetresponse.BudgetSummaryResponse {
	return m.ToBudgetSummaryResponseList(summaries)
}

// GetDistinctYears retorna os anos disponíveis (necessita do gateway)
func (m *BudgetPayloadMapper) GetDistinctYears(ctx context.Context, gateway gateway.BudgetGateway) ([]int, error) {
	return gateway.GetDistinctYears(ctx)
}
