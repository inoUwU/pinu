package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Route(router fiber.Router) error
}
