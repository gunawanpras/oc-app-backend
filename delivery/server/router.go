package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gunawanpras/oc-app-backend/internal/setup"
)

func NewRouter(app *fiber.App, handler setup.Handler) {
	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e any) {
			// stack := debug.Stack()

			fmt.Printf("\n=== Panic recovered ===\n")
			fmt.Printf("Error: %v\n", e)
			fmt.Printf("Path: %s\n", c.Path())
			// fmt.Printf("Stack trace:\n%+v\n", stack)
			fmt.Printf("==========================\n\n")
		},
	}))

	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return nil })

	health := app.Group("/health")
	health.Get("/app", handler.HealthHandler.CheckAppHealth)
	health.Get("/database", handler.HealthHandler.CheckDatabaseHealth)

	profile := app.Group("/profiles")
	profile.Get("/", handler.ProfileHandler.GetProfile)
	profile.Post("/", handler.ProfileHandler.CreateProfile)
	profile.Put("/:id", handler.ProfileHandler.UpdateProfile)
	profile.Delete("/:id", handler.ProfileHandler.DeleteProfile)
}
