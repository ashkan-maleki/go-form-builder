package forms

import "github.com/gofiber/fiber/v2"

type FormHandler struct {
}

func (h *FormHandler) List(c *fiber.Ctx) error {
	return c.SendString("This is my list of forms")
}

func (h *FormHandler) Get(c *fiber.Ctx) error {
	return c.SendString("This is my form")
}
