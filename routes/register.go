package routes

import (
	"bluebird/controllers"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes | @desc: "/cards" api route
func RegisterRoutes(app *fiber.App) {
	registers := app.Group("/v1/register")
	registers.Get("/hello", controllers.HelloRegister)
}
