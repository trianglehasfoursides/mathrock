package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_tag_add.go
func TagAdd(ctx *fiber.Ctx) error {
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

	// Tambah tags baru
	for _, tag := range input.Tags {
		// Check jika tag sudah ada
		var existingTag db.FileTag
		err := db.DB.Where("file_id = ? AND tag = ?", file.ID, tag).First(&existingTag).Error
		if err == nil {
			continue // Tag sudah ada, skip
		}

		filetag := db.FileTag{
			FileID: file.ID,
			Tag:    tag,
		}
		if err := db.DB.Create(&filetag).Error; err != nil {
			return internal.Resperr(ctx, fiber.StatusInternalServerError, "create tag failed", err)
		}
	}

	return ctx.JSON(fiber.Map{"status": "tags added"})
}
