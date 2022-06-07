package router

import (
	"database/sql"
	"workspace/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetRoute(app *fiber.App, db *sql.DB) {
	// Set path and middleware
	api := app.Group("/api/v1", logger.New())

	// Route
	api.Get("/", func(c *fiber.Ctx) error {
		// Get all data
		return handler.GetAllRows(c, db)
	})

	api.Get("/:manufac", func(c *fiber.Ctx) error {
		// Get by manufacturer
		return handler.GetByManufac(c, db)
	})

	api.Post("/new", func(c *fiber.Ctx) error {
		// Add new mobile data
		return handler.AddNewRow(c, db)
	})

	api.Delete("/delete/:id", func(c *fiber.Ctx) error {
		// Delete mobile data
		return handler.DeleteHandler(c, db)
	})
}
