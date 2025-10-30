package note

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/note_create.go
func Add(ctx *fiber.Ctx) error {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	if input.Title == "" {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "title required", nil)
	}

	// Buat note
	note := db.Note{
		UserID:  auth.UserId(ctx),
		Title:   input.Title,
		Content: input.Content,
	}

	if err := db.DB.Create(&note).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "create failed", err)
	}

	return ctx.JSON(note)
}
