package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/mahadevans87/short-url/helpers"
	"github.com/mahadevans87/short-url/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get(":/url", routes.Resolve)
	app.Post("/api/v1", routes.Shorten)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment file.")
	}

	app := fiber.New()
	app.Use(logger.New())

	base62EncodedString := helpers.Base62Encode(9999999)
	fmt.Println(base62EncodedString)
	fmt.Println(helpers.Base62Decode(base62EncodedString))
}
