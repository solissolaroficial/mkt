package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	usecase "github.com/seu-usuario/solis-backend/application/usecase/brand"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type BrandController struct {
	createUseCase *usecase.CreateBrandUseCase
	listUseCase   *usecase.ListBrandsUseCase
	deleteUseCase *usecase.DeleteBrandUseCase
	mapper        *response.BrandPayloadMapper
}

func NewBrandController(
	create *usecase.CreateBrandUseCase,
	list *usecase.ListBrandsUseCase,
	delete *usecase.DeleteBrandUseCase,
) *BrandController {
	return &BrandController{
		createUseCase: create,
		listUseCase:   list,
		deleteUseCase: delete,
		mapper:        response.NewBrandPayloadMapper(),
	}
}

// Create
func (c *BrandController) Create(ctx *fiber.Ctx) error {
	var req request.CreateBrandRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	input := usecase.CreateBrandInput{
		Name: req.Name,
	}

	brand, err := c.createUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating brand: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToBrandResponse(brand))
}

// List
func (c *BrandController) List(ctx *fiber.Ctx) error {
	brands, err := c.listUseCase.Execute(ctx.Context())
	if err != nil {
		log.Printf("Error listing brands: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToBrandsListResponse(brands))
}

// Delete
func (c *BrandController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	uuidID, _ := uuid.Parse(id)
	err := c.deleteUseCase.Execute(ctx.Context(), uuidID)
	if err != nil {
		log.Printf("Error deleting brand: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
