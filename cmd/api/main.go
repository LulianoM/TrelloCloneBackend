package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// database.Connect()
	// database.AutoMigrate()
	// database.SetupRedis()
	// database.SetupCacheChannel()

	app := fiber.New()

	//app.Use(cors.New(cors.Config{
	//	AllowCredentials: true,
	//}))

	//routes.Setup(app)

	app.Listen(":8000")
}
