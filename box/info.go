package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_info.go
func Info(ctx *fiber.Ctx) error {
	fileid := ctx.Params("id")
	var file db.File

	if err := db.DB.Where("id = ? AND user_id = ?", fileid, auth.UserId(ctx)).First(&file).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "file not found", err)
	}

	// Get tags
	var tags []string
	db.DB.Model(&db.FileTag{}).Where("file_id = ?", file.ID).Pluck("tag", &tags)

	// Get metadata
	var metadata []db.FileMeta
	db.DB.Where("file_id = ?", file.ID).Find(&metadata)

	// Convert metadata to map
	metamap := make(map[string]string)
	for _, meta := range metadata {
		metamap[meta.Key] = meta.Value
	}

	return ctx.JSON(fiber.Map{
		"file":     file,
		"tags":     tags,
		"metadata": metamap,
	})
}
