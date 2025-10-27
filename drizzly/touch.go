package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/storage"
)

func Up(ctx *fiber.Ctx) (err error) {
	file, err := ctx.FormFile("")
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "", err)
	}

	storage.Box.Upload(file)
}
