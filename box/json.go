package box

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

func ExportJSON(ctx *fiber.Ctx) error {
	var files []db.File

	if err := db.DB.Where("user_id = ?", auth.UserId(ctx)).Find(&files).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}

	// Build complete data dengan tags dan metadata
	var exportdata []map[string]interface{}

	for _, file := range files {
		// Get tags
		var tags []string
		db.DB.Model(&db.FileTag{}).Where("file_id = ?", file.ID).Pluck("tag", &tags)

		// Get metadata
		var metadata []db.FileMeta
		db.DB.Where("file_id = ?", file.ID).Find(&metadata)
		metamap := make(map[string]string)
		for _, meta := range metadata {
			metamap[meta.Key] = meta.Value
		}

		exportdata = append(exportdata, map[string]interface{}{
			"file":     file,
			"tags":     tags,
			"metadata": metamap,
		})
	}

	return ctx.JSON(fiber.Map{
		"user_id":     auth.UserId(ctx),
		"count":       len(files),
		"data":        exportdata,
		"exported_at": time.Now(),
	})
}
