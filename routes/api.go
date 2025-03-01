package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-auth/app/controllers"
	"go-auth/app/middleware"
)

func ApiRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/register", controllers.RegisterHandler)
	auth.Post("/login", controllers.LoginHandler)

	home := v1.Group("/home", middleware.AuthMiddleware)
	home.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello World!"})
	})
}
