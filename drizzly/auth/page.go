package auth

import "github.com/gofiber/fiber/v2"

func page(ctx *fiber.Ctx) error {
	return ctx.Render("auth.html", fiber.Map{
		"google": "",
		"github": "",
	})
}
