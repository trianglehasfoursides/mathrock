package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_count.go
func Count(ctx *fiber.Ctx) error {
	var count int64
	if err := db.DB.Model(&db.File{}).Where("user_id = ? AND hidden = ?", auth.UserId(ctx), false).Count(&count).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "count failed", err)
	}

	return ctx.JSON(fiber.Map{"count": count})
}
