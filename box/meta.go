package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_meta.go
func Meta(ctx *fiber.Ctx) error {
	fileid := ctx.Params("id")
	var file db.File

	if err := db.DB.Where("id = ? AND user_id = ?", fileid, auth.UserId(ctx)).First(&file).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "file not found", err)
	}

	var metadata []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := ctx.BodyParser(&metadata); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid metadata", err)
	}

	for _, meta := range metadata {
		db.DB.Where("file_id = ? AND key = ?", file.ID, meta.Key).Delete(&db.FileMeta{})
		filemeta := db.FileMeta{
			FileID: file.ID,
			Key:    meta.Key,
			Value:  meta.Value,
		}
		db.DB.Create(&filemeta)
	}

	return ctx.JSON(fiber.Map{"status": "success"})
}
