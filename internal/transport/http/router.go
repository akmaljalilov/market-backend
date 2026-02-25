package http

import "github.com/gofiber/fiber/v2"

func RegisterUsersRoutes(app *fiber.App, h *UsersHandler) {

	api := app.Group("/api")

	sales := api.Group("/users")
	sales.Post("/", h.CreateUser)
}
