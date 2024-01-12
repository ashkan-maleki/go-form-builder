package forms

import "github.com/gofiber/fiber/v2"

func ApiSetup(app *fiber.App) {
	group := app.Group("forms")
	formHandler := FormHandler{}
	group.Get("/", formHandler.List)
	group.Get("/1", formHandler.Get)
}
