package my

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/url_delete.go
func Delete(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var shorturl db.URL
	if err := db.DB.Where("name = ? AND user_id = ?", name, auth.UserId(ctx)).First(&shorturl).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "url not found", err)
	}

	if err := db.DB.Delete(&shorturl).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "delete failed", err)
	}

	return ctx.JSON(fiber.Map{"status": "deleted"})
}
