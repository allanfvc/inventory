package api

import (
	"github.com/allanfvc/inventory/api/controller"
	"github.com/allanfvc/inventory/api/database"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

//StartAPI Inicializa a api
func StartAPI() {
	app := fiber.New()
	database.InitDatabase()
	defer database.CloseDatabase()
	app.Static("/inventory-ui", "./public")
	app.Use(middleware.Logger())
	api := app.Group("/v1")
	controller.RegisterRoutes(api)
	app.Listen(4000)
}
