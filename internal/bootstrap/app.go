package bootstrap

import (
	"encoding/json"
	"market/internal/app/users"
	"market/internal/config"
	"market/internal/deps"
	"market/internal/transport/http"
	"market/internal/transport/http/middleware/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Start(cfg *config.Config) {
	services := deps.Init(cfg)
	userService := users.NewService(services.UserRepo)
	userHandler := http.NewUsersHandler(userService)
	app := fiber.New(fiber.Config{
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
		UnescapePath: true,
		BodyLimit:    500 * 1024 * 1024,
	})
	swagger.SwaggerRoute(app, cfg)
	http.RegisterUsersRoutes(app, userHandler)
	logrus.Fatal(app.Listen(":3000"))

}
