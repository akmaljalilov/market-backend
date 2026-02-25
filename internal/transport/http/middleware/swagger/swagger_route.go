package swagger

import (
	"market/docs"
	"market/internal/config"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a fiber.Router, cfg *config.Config) {
	docs.SwaggerInfo.Title += " Build version: " + os.Getenv("VERSION")
	//u, err := url.Parse(cfg.Info.ServiceHost)
	//if err == nil && u.Host != "" {
	//	docs.SwaggerInfo.Host = u.Host
	//}
	// Create routes group.
	route := a.Group("/swagger")

	// Routes for GET method:
	route.Get("*", swagger.HandlerDefault) // get one user by ID
}
