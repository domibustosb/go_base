package anime

import (
	"log"

	"github.com/domibustosb/go_base/internal/api/v1/validators"
	"github.com/gofiber/fiber/v2"
)

func GetAnime(c *fiber.Ctx) error {
	c.JSON(fiber.Map{"message": "Holi"})
	c.SendStatus(fiber.StatusOK)
	return nil
}

func GetAnimeByID(c *fiber.Ctx) error {
	id := c.Params("id")

	c.JSON(fiber.Map{"id": id, "nombre": "One piece"})
	c.SendStatus(fiber.StatusOK)
	return nil
}

func Create(c *fiber.Ctx) error {
	var animeValidator validators.AnimeValidator

	if err := c.BodyParser(&animeValidator); err != nil {
		c.JSON(fiber.Map{"Error": err})
		c.SendStatus(fiber.StatusBadRequest)
		return nil
	}

	if err := validators.ValidateStruct(animeValidator); err != nil {
		log.Println(err)
		c.JSON(fiber.Map{"Error ValidateStruct": err})
		c.SendStatus(fiber.StatusBadRequest)
		return nil
	}
	c.JSON(animeValidator)
	c.SendStatus(fiber.StatusOK)

	return nil
}
