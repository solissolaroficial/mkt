package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// RegisterSocialPostRoutes registers social post routes
func RegisterSocialPostRoutes(router fiber.Router, socialPostController *controller.SocialPostController) {
	// Social Posts routes
	router.Post("/social/posts", socialPostController.CreateSocialPost)
	router.Get("/social/posts/:id", socialPostController.GetSocialPost)
	router.Get("/social/posts", socialPostController.ListSocialPosts)
	router.Put("/social/posts/:id", socialPostController.UpdateSocialPost)
	router.Delete("/social/posts/:id", socialPostController.DeleteSocialPost)

	// Social Daily Aggregations routes
	router.Post("/social/daily-aggregations/recalculate/:brandID/:date", socialPostController.RecalculateDailyAggregations)
	router.Get("/social/daily-aggregations/:id", socialPostController.GetSocialDailyAggregation)
	router.Get("/social/daily-aggregations", socialPostController.ListSocialDailyAggregations)
}
