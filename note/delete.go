package note

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/note_delete.go
func NoteDelete(ctx *fiber.Ctx) error {
	noteid := ctx.Params("id")

	var note db.Note
	if err := db.DB.Where("id = ? AND user_id = ?", noteid, auth.UserId(ctx)).First(&note).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "note not found", err)
	}

	if err := db.DB.Delete(&note).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "delete failed", err)
	}

	return ctx.JSON(fiber.Map{"status": "deleted"})
}
