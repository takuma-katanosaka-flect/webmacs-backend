package main

import (
	"webmacs-backend/pkg/adapter/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New(fiber.Config{
    CaseSensitive: true,
    StrictRouting: true,
  })

  http.Router(app)

  app.Listen(":3000")
}
