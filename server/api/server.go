package api

import (
	"log"
	"project/vnexpress/api/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	router.SetupRouter(app)

	log.Fatal(app.Listen(":8000"))
}
