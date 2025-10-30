package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/storage"
)

// handlers/file_import.go
func Import(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid form", err)
	}

	files := form.File["files"]
	if len(files) == 0 {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "no files provided", nil)
	}

	var results []map[string]interface{}
	userid := auth.UserId(ctx)

	for _, fileheader := range files {
		hash, uploaderr := storage.Box.Upload(fileheader)
		if uploaderr != nil {
			results = append(results, map[string]interface{}{
				"name":  fileheader.Filename,
				"error": uploaderr.Error(),
			})
			continue
		}

		filedata := db.File{
			Name:   fileheader.Filename,
			Hash:   hash,
			UserID: userid,
		}

		if dberr := db.DB.Create(&filedata).Error; dberr != nil {
			results = append(results, map[string]interface{}{
				"name":  fileheader.Filename,
				"error": dberr.Error(),
			})
		} else {
			results = append(results, map[string]interface{}{
				"name": fileheader.Filename,
				"id":   filedata.ID,
				"hash": hash,
			})
		}
	}

	return ctx.JSON(fiber.Map{"results": results})
}
