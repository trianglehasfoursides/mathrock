package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/storage"
)

func Add(ctx *fiber.Ctx) (err error) {
	file, err := ctx.FormFile("")
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "", err)
	}

	hash, err := storage.Box.Upload(file)
	if err != nil {
		return
	}

	if encrypt := ctx.FormValue("encrypt", "false"); encrypt == "true" {
		return db.DB.Create(&db.File{
			Name:    file.Filename,
			Hash:    hash,
			Encrypt: true,
			UserID:  auth.UserId(ctx),
			Size:    file.Size,
		}).Error
	}

	return db.DB.Create(&db.File{
		Name:   file.Filename,
		Hash:   hash,
		UserID: auth.UserId(ctx),
		Size:   file.Size,
	}).Error
}
