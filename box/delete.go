package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/storage"
)

// handlers/file_delete.go
func Delete(ctx *fiber.Ctx) error {
	fileid := ctx.Params("id")
	var file db.File

	if err := db.DB.Where("id = ? AND user_id = ?", fileid, auth.UserId(ctx)).First(&file).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "file not found", err)
	}

	// Hapus dari storage
	if err := storage.Box.Delete(file.Hash); err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "delete failed", err)
	}

	// Hapus dari database
	return db.DB.Delete(&file).Error
}
