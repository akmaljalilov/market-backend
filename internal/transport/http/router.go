package http

import (
	"market/internal/app/products"
	"market/internal/app/users"
	"market/internal/deps"
	auth2 "market/internal/transport/http/handlers/auth"
	products2 "market/internal/transport/http/handlers/products"
	"market/internal/transport/http/handlers/user"
	"market/internal/transport/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, dependencies *deps.Dependencies) {
	api := app.Group("/api")
	registerAuth(api, dependencies)
	registerUsers(api, dependencies)
	registerProducts(api, dependencies)
}

// Public API
func registerAuth(api fiber.Router, dependencies *deps.Dependencies) {
	userApp := users.New(dependencies.Repo.UserRepo, dependencies.Service.Notify)
	handler := auth2.NewAuthHandler(userApp)
	auth := api.Group("/auth")
	auth.Post("/sign-in", handler.SignIn)
}

// Private API
func registerUsers(api fiber.Router, dependencies *deps.Dependencies) {
	userApp := users.New(dependencies.Repo.UserRepo, dependencies.Service.Notify)
	handler := user.NewUsersHandler(userApp)
	u := api.Group("/users", middleware.JWT())
	u.Get("/", handler.GetCurrentUser)
	u.Post("/", handler.CreateUser)
}

func registerProducts(api fiber.Router, dependencies *deps.Dependencies) {
	productApp := products.New(dependencies.Repo.ProductRepo)
	handler := products2.NewProductsHandler(productApp)
	pr := api.Group("/products", middleware.JWT())
	pr.Post("/", handler.CreateProduct)
}
