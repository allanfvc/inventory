package api

import (
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

//StartAPI Inicializa a api
func StartAPI() {
	app := fiber.New()
	app.Use(middleware.Logger("${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n", os.Stdout, "15:04:05"))
	app.Listen(4000)
}
