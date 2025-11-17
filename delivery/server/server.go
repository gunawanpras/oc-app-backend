package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gunawanpras/oc-app-backend/config"
	"github.com/gunawanpras/oc-app-backend/internal/setup"
)

func Up(handler setup.Handler, config config.ServerConfig) {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				fmt.Printf("Error: %v\n", err)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			},
		},
	)

	NewRouter(app, handler)

	err := app.Listen(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Panic(err)
	}
}
