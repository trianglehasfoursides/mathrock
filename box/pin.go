package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_pin.go
func Pin(ctx *fiber.Ctx) error {
	fileid := ctx.Params("id")
	var file db.File

	if err := db.DB.Where("id = ? AND user_id = ? AND hidden = ?", fileid, auth.UserId(ctx), false).First(&file).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "file not found", err)
	}

	file.Pinned = true
	return db.DB.Save(&file).Error
}
