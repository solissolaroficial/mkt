package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	accountpayable "github.com/seu-usuario/solis-backend/application/usecase/financial/accountpayable"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/mapper"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type AccountPayableController struct {
	createUseCase        *accountpayable.CreateAccountPayableUseCase
	listUseCase          *accountpayable.ListAccountsPayableUseCase
	getUseCase           *accountpayable.GetAccountPayableUseCase
	updateUseCase        *accountpayable.UpdateAccountPayableUseCase
	deleteUseCase        *accountpayable.DeleteAccountPayableUseCase
	toggleNFUseCase      *accountpayable.ToggleNFUseCase
	toggleBoletoUseCase  *accountpayable.ToggleBoletoUseCase
	sendToFinanceUseCase *accountpayable.SendToFinanceUseCase
	mapper               *mapper.AccountPayablePayloadMapper
}

func NewAccountPayableController(
	create *accountpayable.CreateAccountPayableUseCase,
	list *accountpayable.ListAccountsPayableUseCase,
	get *accountpayable.GetAccountPayableUseCase,
	update *accountpayable.UpdateAccountPayableUseCase,
	delete *accountpayable.DeleteAccountPayableUseCase,
	toggleNF *accountpayable.ToggleNFUseCase,
	toggleBoleto *accountpayable.ToggleBoletoUseCase,
	sendToFinance *accountpayable.SendToFinanceUseCase,
) *AccountPayableController {
	return &AccountPayableController{
		createUseCase:        create,
		listUseCase:          list,
		getUseCase:           get,
		updateUseCase:        update,
		deleteUseCase:        delete,
		toggleNFUseCase:      toggleNF,
		toggleBoletoUseCase:  toggleBoleto,
		sendToFinanceUseCase: sendToFinance,
		mapper:               mapper.NewAccountPayablePayloadMapper(),
	}
}

// Create
func (c *AccountPayableController) Create(ctx *fiber.Ctx) error {
	var req request.CreateAccountPayableRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := c.mapper.CreateRequestToInput(&req)
	output, err := c.createUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating account payable: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Account payable created successfully: %s", output.ID)
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.OutputToResponse(output))
}

// List
func (c *AccountPayableController) List(ctx *fiber.Ctx) error {
	var query request.ListAccountsPayableQuery
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Executar use case
	input := c.mapper.ListQueryToInput(&query)
	outputs, total, err := c.listUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing accounts payable: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	responses := c.mapper.ListOutputsToResponses(outputs)

	// Determinar página e limite para o meta
	page := 1
	limit := len(responses)
	if query.Page != nil {
		page = *query.Page
	}
	if query.Limit != nil {
		limit = *query.Limit
	}

	log.Printf("Accounts payable listed successfully: count=%d, total=%d", len(responses), total)
	return ctx.JSON(response.AccountPayableListData{
		Accounts: responses,
		Meta: response.MetaResponse{
			Total: total,
			Page:  page,
			Limit: limit,
		},
	})
}

// GetByID
func (c *AccountPayableController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.GetRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.getUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrAccountPayableNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Account payable not found",
			})
		}
		log.Printf("Error getting account payable %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.GetOutputToResponse(output))
}

// Update
func (c *AccountPayableController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateAccountPayableRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	input, err := c.mapper.UpdateRequestToInput(id, &req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.updateUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrAccountPayableNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Account payable not found",
			})
		}
		log.Printf("Error updating account payable %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Account payable updated successfully: %s", id)
	return ctx.JSON(c.mapper.UpdateOutputToResponse(output))
}

// Delete
func (c *AccountPayableController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.DeleteRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err = c.deleteUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrAccountPayableNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Account payable not found",
			})
		}
		if errors.Is(err, domainerrors.ErrCannotDeleteSentAccount) {
			return ctx.Status(fiber.StatusConflict).JSON(response.ErrorResponse{
				Error: "Cannot delete account that has been sent to finance",
			})
		}
		log.Printf("Error deleting account payable %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Account payable deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// ToggleNF
func (c *AccountPayableController) ToggleNF(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.ToggleNFRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.toggleNFUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrAccountPayableNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Account payable not found",
			})
		}
		log.Printf("Error toggling NF for account payable %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("NF toggled successfully for account payable: %s", id)
	return ctx.JSON(c.mapper.ToggleNFOutputToResponse(output))
}

// ToggleBoleto
func (c *AccountPayableController) ToggleBoleto(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.ToggleBoletoRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.toggleBoletoUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrAccountPayableNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Account payable not found",
			})
		}
		log.Printf("Error toggling boleto for account payable %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Boleto toggled successfully for account payable: %s", id)
	return ctx.JSON(c.mapper.ToggleBoletoOutputToResponse(output))
}

// SendToFinance
func (c *AccountPayableController) SendToFinance(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.SendToFinanceRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.sendToFinanceUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrAccountPayableNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Account payable not found",
			})
		}
		log.Printf("Error sending account payable %s to finance: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Account payable sent to finance successfully: %s", id)
	return ctx.JSON(c.mapper.SendToFinanceOutputToResponse(output))
}
