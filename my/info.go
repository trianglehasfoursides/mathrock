package my

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/url_info.go
func Info(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var shorturl db.URL
	if err := db.DB.Where("name = ? AND user_id = ?", name, auth.UserId(ctx)).First(&shorturl).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "url not found", err)
	}

	return ctx.JSON(shorturl)
}
