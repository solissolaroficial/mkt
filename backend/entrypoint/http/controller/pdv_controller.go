package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	usecase "github.com/seu-usuario/solis-backend/application/usecase/pdv"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// PdvController controla as requisições HTTP para PDV
type PdvController struct {
	createPdvPostUseCase       *usecase.CreatePdvPostUseCase
	listPdvPostsUseCase        *usecase.ListPdvPostsUseCase
	getPdvPostUseCase          *usecase.GetPdvPostUseCase
	updatePdvPostUseCase       *usecase.UpdatePdvPostUseCase
	updatePdvPostStatusUseCase *usecase.UpdatePdvPostStatusUseCase
	deletePdvPostUseCase       *usecase.DeletePdvPostUseCase
	createRecurrentPdvUseCase  *usecase.CreateRecurrentPdvUseCase
	getRecurrentPdvUseCase     *usecase.GetRecurrentPdvUseCase
	updateRecurrentPdvUseCase  *usecase.UpdateRecurrentPdvUseCase
	listRecurrentPdvsUseCase   *usecase.ListRecurrentPdvsUseCase
	deleteRecurrentPdvUseCase  *usecase.DeleteRecurrentPdvUseCase
	mapper                     *response.PdvPayloadMapper
}

// NewPdvController cria uma nova instância do PdvController
func NewPdvController(
	createPdvPost *usecase.CreatePdvPostUseCase,
	listPdvPosts *usecase.ListPdvPostsUseCase,
	getPdvPost *usecase.GetPdvPostUseCase,
	updatePdvPost *usecase.UpdatePdvPostUseCase,
	updatePdvPostStatus *usecase.UpdatePdvPostStatusUseCase,
	deletePdvPost *usecase.DeletePdvPostUseCase,
	createRecurrentPdv *usecase.CreateRecurrentPdvUseCase,
	getRecurrentPdv *usecase.GetRecurrentPdvUseCase,
	updateRecurrentPdv *usecase.UpdateRecurrentPdvUseCase,
	listRecurrentPdvs *usecase.ListRecurrentPdvsUseCase,
	deleteRecurrentPdv *usecase.DeleteRecurrentPdvUseCase,
) *PdvController {
	return &PdvController{
		createPdvPostUseCase:       createPdvPost,
		listPdvPostsUseCase:        listPdvPosts,
		getPdvPostUseCase:          getPdvPost,
		updatePdvPostUseCase:       updatePdvPost,
		updatePdvPostStatusUseCase: updatePdvPostStatus,
		deletePdvPostUseCase:       deletePdvPost,
		createRecurrentPdvUseCase:  createRecurrentPdv,
		getRecurrentPdvUseCase:     getRecurrentPdv,
		updateRecurrentPdvUseCase:  updateRecurrentPdv,
		listRecurrentPdvsUseCase:   listRecurrentPdvs,
		deleteRecurrentPdvUseCase:  deleteRecurrentPdv,
		mapper:                     response.NewPdvPayloadMapper(),
	}
}

// CreatePdvPost cria um novo post de PDV
func (c *PdvController) CreatePdvPost(ctx *fiber.Ctx) error {
	var req request.CreatePdvPostRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case (validação está no use case)
	input := usecase.CreatePdvPostInput{
		RepName:  req.RepName,
		PdvName:  req.PdvName,
		PostDate: req.PostDate,
		Platform: req.Platform,
		Link:     req.Link,
		ProofUrl: req.ProofUrl,
	}

	post, err := c.createPdvPostUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating PDV post: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("PDV post created successfully: %s", post.ID())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToPdvPostResponse(post))
}

// ListPdvPosts lista posts de PDV
func (c *PdvController) ListPdvPosts(ctx *fiber.Ctx) error {
	var query request.ListPdvPostsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Valores padrão para paginação
	if query.Page < 1 {
		query.Page = 1
	}
	if query.Limit < 1 || query.Limit > 100 {
		query.Limit = 10
	}

	// Executar use case
	input := usecase.ListPdvPostsInput{
		RepName:   query.RepName,
		Month:     query.Month,
		Platform:  query.Platform,
		Status:    query.Status,
		StartDate: query.StartDate,
		EndDate:   query.EndDate,
		Page:      query.Page,
		Limit:     query.Limit,
		SortBy:    query.SortBy,
		SortOrder: query.SortOrder,
	}

	posts, total, err := c.listPdvPostsUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing PDV posts: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("PDV posts listed successfully: page=%d, limit=%d, total=%d", query.Page, query.Limit, total)
	return ctx.JSON(c.mapper.ToPdvPostsListResponse(posts, total, query.Page, query.Limit))
}

// UpdatePdvPostStatus atualiza o status de um post de PDV
func (c *PdvController) UpdatePdvPostStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdatePdvPostStatusRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Obter usuário do contexto de autenticação
	user := ""
	if userID := ctx.Locals("user"); userID != nil {
		if userIDStr, ok := userID.(string); ok {
			user = userIDStr
		}
	}

	// Executar use case
	input := usecase.UpdateStatusInput{
		PostID:    id,
		NewStatus: req.Status,
		User:      user,
	}

	if err := c.updatePdvPostStatusUseCase.Execute(ctx.Context(), input); err != nil {
		if errors.Is(err, domainErrors.ErrPdvPostNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "PDV post not found",
			})
		}
		log.Printf("Error updating status for PDV post %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("PDV post status updated successfully: %s to %s by user %s", id, req.Status, user)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Status updated successfully",
	})
}

// CreateRecurrentPdv cria um novo PDV recorrente
func (c *PdvController) CreateRecurrentPdv(ctx *fiber.Ctx) error {
	var req request.CreateRecurrentPdvRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case (validação está no use case)
	input := usecase.CreateRecurrentPdvInput{
		Name:    req.Name,
		RepName: req.RepName,
		City:    req.City,
	}

	pdv, err := c.createRecurrentPdvUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating recurrent PDV: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Recurrent PDV created successfully: %s", pdv.ID())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToRecurrentPdvResponse(pdv))
}

// ListRecurrentPdvs lista PDVs recorrentes
func (c *PdvController) ListRecurrentPdvs(ctx *fiber.Ctx) error {
	var query request.ListRecurrentPdvsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Valores padrão para paginação
	if query.Page < 1 {
		query.Page = 1
	}
	if query.Limit < 1 || query.Limit > 100 {
		query.Limit = 10
	}

	// Executar use case
	input := usecase.ListRecurrentPdvsInput{
		RepName:   query.RepName,
		City:      query.City,
		Page:      query.Page,
		Limit:     query.Limit,
		SortBy:    query.SortBy,
		SortOrder: query.SortOrder,
	}

	pdvs, total, err := c.listRecurrentPdvsUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing recurrent PDVs: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Recurrent PDVs listed successfully: page=%d, limit=%d, total=%d", query.Page, query.Limit, total)
	return ctx.JSON(c.mapper.ToRecurrentPdvsListResponse(pdvs, total, query.Page, query.Limit))
}

// GetRecurrentPdv retorna um PDV recorrente por ID
func (c *PdvController) GetRecurrentPdv(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	pdv, err := c.getRecurrentPdvUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrRecurrentPdvNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Recurrent PDV not found",
			})
		}
		log.Printf("Error getting recurrent PDV %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToRecurrentPdvResponse(pdv))
}

// DeleteRecurrentPdv exclui um PDV recorrente
func (c *PdvController) DeleteRecurrentPdv(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	if err := c.deleteRecurrentPdvUseCase.Execute(ctx.Context(), id); err != nil {
		if errors.Is(err, domainErrors.ErrRecurrentPdvNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Recurrent PDV not found",
			})
		}
		log.Printf("Error deleting recurrent PDV %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Recurrent PDV deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// GetPdvPost retorna um post por ID
func (c *PdvController) GetPdvPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	post, err := c.getPdvPostUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrPdvPostNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "PDV post not found",
			})
		}
		log.Printf("Error getting PDV post %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToPdvPostResponse(post))
}

// UpdatePdvPost atualiza um post existente
func (c *PdvController) UpdatePdvPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdatePdvPostRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := usecase.UpdatePdvPostInput{
		ID:       id,
		RepName:  req.RepName,
		PdvName:  req.PdvName,
		PostDate: req.PostDate,
		Platform: req.Platform,
		Link:     req.Link,
		ProofUrl: req.ProofUrl,
	}

	post, err := c.updatePdvPostUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainErrors.ErrPdvPostNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "PDV post not found",
			})
		}
		log.Printf("Error updating PDV post %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("PDV post updated successfully: %s", id)
	return ctx.JSON(c.mapper.ToPdvPostResponse(post))
}

// DeletePdvPost deleta um post
func (c *PdvController) DeletePdvPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err := c.deletePdvPostUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrPdvPostNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "PDV post not found",
			})
		}
		log.Printf("Error deleting PDV post %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("PDV post deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// UpdateRecurrentPdv atualiza um PDV recorrente existente
func (c *PdvController) UpdateRecurrentPdv(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateRecurrentPdvRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := usecase.UpdateRecurrentPdvInput{
		ID:               id,
		Name:             req.Name,
		RepName:          req.RepName,
		City:             req.City,
		Followers:        req.Followers,
		InstagramProfile: req.InstagramProfile,
	}

	pdv, err := c.updateRecurrentPdvUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainErrors.ErrRecurrentPdvNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Recurrent PDV not found",
			})
		}
		log.Printf("Error updating recurrent PDV %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Recurrent PDV updated successfully: %s", id)
	return ctx.JSON(c.mapper.ToRecurrentPdvResponse(pdv))
}
