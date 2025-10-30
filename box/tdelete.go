package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_tag_delete.go
func TagDelete(ctx *fiber.Ctx) error {
	fileid := ctx.Params("id")
	var file db.File

	if err := db.DB.Where("id = ? AND user_id = ?", fileid, auth.UserId(ctx)).First(&file).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "file not found", err)
	}

	var input struct {
		Tags []string `json:"tags"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	if len(input.Tags) == 0 {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "tags required", nil)
	}

	// Hapus tags
	result := db.DB.Where("file_id = ? AND tag IN ?", file.ID, input.Tags).Delete(&db.FileTag{})
	if result.Error != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "delete tags failed", result.Error)
	}

	return ctx.JSON(fiber.Map{
		"status": "tags deleted",
		"count":  result.RowsAffected,
	})
}
