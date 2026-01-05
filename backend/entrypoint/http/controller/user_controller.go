package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/application/usecase/users"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// UserController handles HTTP requests for users
type UserController struct {
	listUsersUseCase *users.ListUsersUseCase
	mapper           *UserMapper
}

// NewUserController creates a new UserController instance
func NewUserController(listUsersUseCase *users.ListUsersUseCase) *UserController {
	return &UserController{
		listUsersUseCase: listUsersUseCase,
		mapper:           &UserMapper{},
	}
}

// ListUsers returns all active users with pagination
// GET /api/users?page=1&limit=20
func (c *UserController) ListUsers(ctx *fiber.Ctx) error {
	// Parse pagination parameters
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	users, err := c.listUsersUseCase.Execute(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Failed to list users",
		})
	}

	// Apply pagination
	total := len(users)
	totalPages := (total + limit - 1) / limit
	start := (page - 1) * limit
	end := start + limit

	if start >= total {
		// Return empty result for out of range pages
		return ctx.JSON(response.PaginationResponse{
			Page:       page,
			PageSize:   limit,
			Total:      int64(total),
			TotalPages: totalPages,
			Data:       []response.PublicUserResponse{},
		})
	}

	if end > total {
		end = total
	}

	paginatedUsers := users[start:end]
	userResponses := c.mapper.ToPublicResponseList(paginatedUsers)

	return ctx.JSON(response.PaginationResponse{
		Page:       page,
		PageSize:   limit,
		Total:      int64(total),
		TotalPages: totalPages,
		Data:       userResponses,
	})
}

// ListAllUsers returns all active users without pagination
// GET /api/users/all
func (c *UserController) ListAllUsers(ctx *fiber.Ctx) error {
	users, err := c.listUsersUseCase.Execute(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Failed to list users",
		})
	}

	userResponses := c.mapper.ToPublicResponseList(users)
	return ctx.JSON(userResponses)
}
