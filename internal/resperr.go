package internal

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Resperr(ctx *fiber.Ctx, code int, msg string, err error) error {
	switch os.Getenv("") {
	case "debug":
		return ctx.Status(code).JSON(fiber.Map{
			"error": err.Error(),
		})
	default:
		return ctx.Status(code).JSON(fiber.Map{
			"error": msg,
		})
	}
}
