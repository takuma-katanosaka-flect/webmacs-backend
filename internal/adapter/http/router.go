func router(app) {
	app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })
}