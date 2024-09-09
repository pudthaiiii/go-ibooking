package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type middleware struct {
}

type Middleware interface {
	Authenticate(next fiber.Handler) fiber.Handler
	RequiredMerchant(next fiber.Handler) fiber.Handler
}

func NewMiddleware() Middleware {
	return &middleware{}
}
