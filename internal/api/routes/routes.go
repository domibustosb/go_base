package routes

import (
	"github.com/domibustosb/go_base/internal/api/v1/controllers/anime"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes is a route management of the app
func SetupRoutes(r *fiber.App) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hola team")
	})
	v1 := r.Group("/v1")
	{
		v1.Get("/health", func(c *fiber.Ctx) error {
			return c.SendString("I am alive")
		})
		a := v1.Group("/anime")
		{
			a.Get("/", anime.GetAnime)
			a.Get("/:id", anime.GetAnimeByID)
			a.Post("/", anime.Create)
		}
	}
}
