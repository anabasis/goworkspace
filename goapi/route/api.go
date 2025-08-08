package route

import (
	"github.com/anabasis/goapi/controller/svc1"
	"github.com/anabasis/goapi/controller/svc2"
	"github.com/gofiber/fiber/v2"
)

func Router() *fiber.App {

	app := fiber.New()

	// Define a simple route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"message": "Hello, World!"})
	})

	// Service 1 routes
	app_svc1 := app.Group("/svc1")
	app_svc1.Get("/req1", svc1.Req1)
	app_svc1.Get("/req2", svc1.Req2)

	// Service 2 routes
	app_svc2 := app.Group("/svc2")
	app_svc2.Get("/req1", svc2.Req1)
	app_svc2.Get("/req2", svc2.Req2)

	return app
}
