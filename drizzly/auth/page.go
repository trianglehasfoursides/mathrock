package auth

import "github.com/gofiber/fiber/v2"

func Page(ctx *fiber.Ctx) error {
	return ctx.Render("auth.html", fiber.Map{
		"google": "",
		"github": "",
	})
}
