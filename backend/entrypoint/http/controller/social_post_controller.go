package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/seu-usuario/solis-backend/application/usecase/social"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
)

type SocialPostController struct {
	createUseCase                  *social.CreateSocialPostUseCase
	updateUseCase                  *social.UpdateSocialPostUseCase
	deleteUseCase                  *social.DeleteSocialPostUseCase
	listUseCase                    *social.ListSocialPostsUseCase
	getUseCase                     *social.GetSocialPostUseCase
	recalculateAggregationsUseCase *social.RecalculateDailyAggregationsUseCase
	listAggregationsUseCase        *social.ListSocialDailyAggregationsUseCase
	getAggregationUseCase          *social.GetSocialDailyAggregationUseCase
}

func NewSocialPostController(
	createUseCase *social.CreateSocialPostUseCase,
	updateUseCase *social.UpdateSocialPostUseCase,
	deleteUseCase *social.DeleteSocialPostUseCase,
	listUseCase *social.ListSocialPostsUseCase,
	getUseCase *social.GetSocialPostUseCase,
	recalculateAggregationsUseCase *social.RecalculateDailyAggregationsUseCase,
	listAggregationsUseCase *social.ListSocialDailyAggregationsUseCase,
	getAggregationUseCase *social.GetSocialDailyAggregationUseCase,
) *SocialPostController {
	return &SocialPostController{
		createUseCase:                  createUseCase,
		updateUseCase:                  updateUseCase,
		deleteUseCase:                  deleteUseCase,
		listUseCase:                    listUseCase,
		getUseCase:                     getUseCase,
		recalculateAggregationsUseCase: recalculateAggregationsUseCase,
		listAggregationsUseCase:        listAggregationsUseCase,
		getAggregationUseCase:          getAggregationUseCase,
	}
}

// CreateSocialPost creates a new social post
func (c *SocialPostController) CreateSocialPost(ctx *fiber.Ctx) error {
	var req request.CreateSocialPostRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	input := social.CreateSocialPostInput{
		BrandName:       req.BrandName,
		Platform:        req.Platform,
		PostDate:        req.PostDate,
		PostTime:        req.PostTime,
		Likes:           req.Likes,
		Comments:        req.Comments,
		Shares:          req.Shares,
		PostType:        req.PostType,
		Caption:         req.Caption,
		FollowersAtPost: req.FollowersAtPost,
	}

	post, err := c.createUseCase.Execute(input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": post,
	})
}

// GetSocialPost gets a social post by ID
func (c *SocialPostController) GetSocialPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	post, err := c.getUseCase.Execute(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": post,
	})
}

// ListSocialPosts lists social posts with filters
func (c *SocialPostController) ListSocialPosts(ctx *fiber.Ctx) error {
	var req request.ListSocialPostsRequest
	if err := ctx.QueryParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	input := social.ListSocialPostsInput{
		BrandName:     req.BrandName,
		Platform:      req.Platform,
		PostType:      req.PostType,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		MinFollowers:  req.MinFollowers,
		MaxFollowers:  req.MaxFollowers,
		MinEngagement: req.MinEngagement,
		MaxEngagement: req.MaxEngagement,
		OnlyActive:    req.OnlyActive,
		Page:          req.Page,
		PageSize:      req.PageSize,
		SortBy:        req.SortBy,
		SortOrder:     req.SortOrder,
	}

	output, err := c.listUseCase.Execute(ctx.Context(), input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": output,
	})
}

// UpdateSocialPost updates a social post
func (c *SocialPostController) UpdateSocialPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var req request.UpdateSocialPostRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	input := social.UpdateSocialPostInput{
		ID:              id,
		Platform:        req.Platform,
		PostDate:        req.PostDate,
		PostTime:        req.PostTime,
		Likes:           req.Likes,
		Comments:        req.Comments,
		Shares:          req.Shares,
		PostType:        req.PostType,
		Caption:         req.Caption,
		FollowersAtPost: req.FollowersAtPost,
	}

	post, err := c.updateUseCase.Execute(ctx.Context(), input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": post,
	})
}

// DeleteSocialPost deletes a social post
func (c *SocialPostController) DeleteSocialPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := c.deleteUseCase.Execute(ctx.Context(), id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// ListSocialDailyAggregations lists social daily aggregations
func (c *SocialPostController) ListSocialDailyAggregations(ctx *fiber.Ctx) error {
	var req request.ListSocialDailyAggregationsRequest
	if err := ctx.QueryParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	input := social.ListSocialDailyAggregationsInput{
		BrandName: req.BrandName,
		Platform:  req.Platform,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Page:      req.Page,
		PageSize:  req.PageSize,
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}

	output, err := c.listAggregationsUseCase.Execute(ctx.Context(), input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": output,
	})
}

// RecalculateDailyAggregations recalculates daily aggregations for a brand and date
func (c *SocialPostController) RecalculateDailyAggregations(ctx *fiber.Ctx) error {
	brandName := ctx.Params("brandName")
	dateStr := ctx.Params("date")

	// Parse date from string
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid date format, expected YYYY-MM-DD",
		})
	}

	aggregation, err := c.recalculateAggregationsUseCase.Execute(brandName, date)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Daily aggregations recalculated successfully",
		"data":    aggregation,
	})
}

// GetSocialDailyAggregation gets a social daily aggregation by ID
func (c *SocialPostController) GetSocialDailyAggregation(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	aggregation, err := c.getAggregationUseCase.Execute(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"data": aggregation,
	})
}
