package auth

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/shareed2k/goth_fiber"
)

func redirect(ctx *fiber.Ctx) error {
	return auth.BeginAuthHandler(ctx)
}
