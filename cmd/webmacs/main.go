package main

import (
	"os"
	"webmacs-backend/pkg/adapter/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New(fiber.Config{
    CaseSensitive: true,
    StrictRouting: true,
  })

  http.Router(app)
  listen(app)
}

func listen(app *fiber.App) {
  port := os.Getenv("PORT")

  if (port == "") {
    app.Listen(":3000")
  }

  app.Listen(":" + port)
}
