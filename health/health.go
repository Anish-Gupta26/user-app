package health

import "github.com/gofiber/fiber/v2"

func Check(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"ping": "pong"})
}
