package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_recent.go
func Recent(ctx *fiber.Ctx) error {
	var files []db.File
	limit := ctx.QueryInt("limit", 10)

	if err := db.DB.Where("user_id = ?", auth.UserId(ctx)).
		Order("updated_at DESC").
		Limit(limit).
		Find(&files).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}

	return ctx.JSON(files)
}
