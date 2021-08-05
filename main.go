package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"

	"bluebird/routes"
)

// Routes function
func setupRoutes(app *fiber.App) {
	app.Use(logger.New())
	// //:storeID
	app.Get("/v1/hello/reverse:paramID", func(c *fiber.Ctx) error {
		paramID := c.Params("paramID")
		result := ""
		for i := 0; i < len(paramID); i++ {

		}
		return c.JSON(fiber.Map{"code": 200, "data": map[string]string{"message": result}, "success": true})
		//return c.JSON(fiber.Map{"code": 200, "data": map[string]string{"message": "Hello World"}, "success": true})

		//return c.JSON(fiber.Map{"code": 200, "data": `"message":"Hello World"`, "success": true})
		//return c.JSON(fiber.Map{"code": 200, "message": "Hello World", "success": true})

		/*
					output
					{
			"code": 200,
			"data": {
			"message":"olleH"
			},
			"success": true
			}
		*/
	})

	// Unrestricted routes
	app.Post("/check_service", func(c *fiber.Ctx) error {
		//return c.SendString("Service Karim Card GoFiber MongoDB euy 👋!")
		return c.JSON(fiber.Map{"code": 200, "data": "Service Karim Card GoFiber MongoDB euy 👋!", "success": true})
	})

	routes.RegisterRoutes(app)
}
func main() {

	// Create new fiber instance
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	})

	// Routes
	setupRoutes(app)

	app.Listen(fmt.Sprintf("%s:%s", "localhost", "8080"))
	//app.Listen(fmt.Sprintf("%s:%s", os.Getenv("IP_SERVER"), os.Getenv("PORT_SERVER")))
}
