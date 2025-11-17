package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetProfile(c *fiber.Ctx) error
	CreateProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
	DeleteProfile(c *fiber.Ctx) error
}
