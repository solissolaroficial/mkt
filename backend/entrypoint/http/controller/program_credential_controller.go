package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	usecase "github.com/seu-usuario/solis-backend/application/usecase/credentials"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type ProgramCredentialController struct {
	listUseCase   *usecase.ListCredentialsUseCase
	getUseCase    *usecase.GetCredentialUseCase
	createUseCase *usecase.CreateCredentialUseCase
	updateUseCase *usecase.UpdateCredentialUseCase
	deleteUseCase *usecase.DeleteCredentialUseCase
	mapper        *ProgramCredentialPayloadMapper
}

func NewProgramCredentialController(
	list *usecase.ListCredentialsUseCase,
	get *usecase.GetCredentialUseCase,
	create *usecase.CreateCredentialUseCase,
	update *usecase.UpdateCredentialUseCase,
	delete *usecase.DeleteCredentialUseCase,
) *ProgramCredentialController {
	return &ProgramCredentialController{
		listUseCase:   list,
		getUseCase:    get,
		createUseCase: create,
		updateUseCase: update,
		deleteUseCase: delete,
		mapper:        NewProgramCredentialPayloadMapper(),
	}
}

// List returns all active credentials
func (c *ProgramCredentialController) List(ctx *fiber.Ctx) error {
	credentials, err := c.listUseCase.Execute(ctx.Context())
	if err != nil {
		log.Printf("Error listing credentials: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(response.CredentialsListResponse(c.mapper.ToResponseList(credentials)))
}

// Get returns a single credential by ID
func (c *ProgramCredentialController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	credential, err := c.getUseCase.Execute(ctx.Context(), usecase.GetCredentialInput{
		ID: parsedID,
	})
	if err != nil {
		log.Printf("Error getting credential: %v", err)
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(response.CredentialResponse(c.mapper.ToResponse(credential)))
}

// Create creates a new credential
func (c *ProgramCredentialController) Create(ctx *fiber.Ctx) error {
	var req request.CreateCredentialRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	credential, err := c.createUseCase.Execute(ctx.Context(), usecase.CreateCredentialInput{
		Name:     req.Name,
		User:     req.User,
		Password: req.Password,
		Access:   req.Access,
		Notes:    req.Notes,
	})
	if err != nil {
		log.Printf("Error creating credential: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.CredentialResponse(c.mapper.ToResponse(credential)))
}

// Update updates an existing credential
func (c *ProgramCredentialController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateCredentialRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	credential, err := c.updateUseCase.Execute(ctx.Context(), usecase.UpdateCredentialInput{
		ID:       parsedID,
		Name:     req.Name,
		User:     req.User,
		Password: req.Password,
		Access:   req.Access,
		Notes:    req.Notes,
	})
	if err != nil {
		log.Printf("Error updating credential: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(response.CredentialResponse(c.mapper.ToResponse(credential)))
}

// Delete soft-deletes a credential
func (c *ProgramCredentialController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	parsedID, err := uuid.Parse(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err = c.deleteUseCase.Execute(ctx.Context(), usecase.DeleteCredentialInput{
		ID: parsedID,
	})
	if err != nil {
		log.Printf("Error deleting credential: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
