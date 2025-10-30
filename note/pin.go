package note

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/note_pin.go
func Pin(ctx *fiber.Ctx) error {
	noteid := ctx.Params("id")

	var note db.Note
	if err := db.DB.Where("id = ? AND user_id = ?", noteid, auth.UserId(ctx)).First(&note).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "note not found", err)
	}

	note.Pinned = true
	if err := db.DB.Save(&note).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "pin failed", err)
	}

	return ctx.JSON(fiber.Map{"status": "pinned"})
}
