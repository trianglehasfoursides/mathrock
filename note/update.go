package note

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/note_update.go
func Update(ctx *fiber.Ctx) error {
	noteid := ctx.Params("id")

	var note db.Note
	if err := db.DB.Where("id = ? AND user_id = ?", noteid, auth.UserId(ctx)).First(&note).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "note not found", err)
	}

	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	// Update fields jika provided
	if input.Title != "" {
		note.Title = input.Title
	}

	if err := db.DB.Save(&note).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "update failed", err)
	}

	return ctx.JSON(note)
}
