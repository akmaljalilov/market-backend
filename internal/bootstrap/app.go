package bootstrap

import (
	"encoding/json"
	"market/internal/config"
	"market/internal/deps"
	"market/internal/transport/http"
	"market/internal/transport/http/middleware/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Start(cfg *config.Config) {
	services := deps.Init(cfg)
	app := fiber.New(fiber.Config{
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
		UnescapePath: true,
		BodyLimit:    500 * 1024 * 1024,
	})
	swagger.SwaggerRoute(app, cfg)
	http.RegisterRoutes(app, services)
	logrus.Fatal(app.Listen(":3000"))

}
