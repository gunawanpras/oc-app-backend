package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CheckAppHealth(c *fiber.Ctx) error
	CheckDatabaseHealth(c *fiber.Ctx) error
}
