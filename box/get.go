package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/storage"
)

func Get(ctx *fiber.Ctx) error {
	fileID := ctx.Params("id")
	var file db.File

	if err := db.DB.Where("id = ? AND user_id = ?", fileID, auth.UserId(ctx)).First(&file).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "file not found", err)
	}

	reader, err := storage.Box.Get(file.Hash)
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "get file failed", err)
	}
	defer reader.Close()

	// Set headers untuk download
	ctx.Set("Content-Disposition", "attachment; filename="+file.Name)
	ctx.Set("Content-Type", "application/octet-stream")

	return ctx.SendStream(reader)
}
