package handlers

import "github.com/gofiber/fiber/v2"

func SampleHandler(c *fiber.Ctx) error {
	id := c.Locals("id").(string)
	email := c.Locals("email").(string)
	return c.SendString("Hello, " + id + " " + email)
}
