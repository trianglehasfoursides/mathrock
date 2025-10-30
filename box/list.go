package box

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_list.go
func List(ctx *fiber.Ctx) error {
	var files []db.File
	query := db.DB.Where("user_id = ? AND hidden = ?", auth.UserId(ctx), false)

	// Filter by tags jika ada
	if tags := ctx.Query("tags"); tags != "" {
		// Cari file IDs yang punya tags tersebut
		var fileids []uint
		db.DB.Model(&db.FileTag{}).Where("tag IN ?", strings.Split(tags, ",")).Pluck("file_id", &fileids)
		query = query.Where("id IN ?", fileids)
	}

	// Filter pinned jika ada
	if pinned := ctx.Query("pinned"); pinned != "" {
		query = query.Where("pinned = ?", pinned == "true")
	}

	if err := query.Find(&files).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}

	return ctx.JSON(files)
}
