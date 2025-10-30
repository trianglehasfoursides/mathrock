package box

import (
	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/storage"
)

func Addx(ctx *fiber.Ctx) (err error) {
	file, err := ctx.FormFile("")
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "", err)
	}

	hash, err := storage.Box.Upload(file)
	if err != nil {
		return
	}

	if _, err = squirrel.Insert("boxes").
		Columns("name", "hash", "encrypt", "user_id").
		Values(file.Filename, hash, auth.UserId(ctx)).
		RunWith(db.DB).
		Exec(); err != nil {
		return
	}

	return
}
