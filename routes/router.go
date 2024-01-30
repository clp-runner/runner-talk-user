package routes

import (
	"github.com/clp-runner/runner-user/users"
	"github.com/gofiber/fiber/v3"
)

func Router() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	google := users.NewGoogle()
	app.Get("/login", google.Login)
	app.Get("/auth/google", google.Callback)
	return app
}