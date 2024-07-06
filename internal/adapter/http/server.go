package http

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Serve() {
	app := fiber.New()
	port := 3001

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
