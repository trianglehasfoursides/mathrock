package my

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/url_redirect.go
func Redirect(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var shorturl db.URL
	if err := db.DB.Where("name = ?", name).First(&shorturl).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "url not found", err)
	}

	// Update click count
	db.DB.Model(&shorturl).Update("clicks", shorturl.Clicks+1)

	// Redirect ke original URL
	return ctx.Redirect(shorturl.OriginalURL, fiber.StatusFound)
}
