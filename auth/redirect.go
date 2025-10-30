package auth

import (
	"github.com/gofiber/fiber/v2"
	aoth "github.com/shareed2k/goth_fiber"
)

func Redirect(ctx *fiber.Ctx) error {
	return aoth.BeginAuthHandler(ctx)
}
