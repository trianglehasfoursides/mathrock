package kanban

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/kanban_delete.go
func Delete(ctx *fiber.Ctx) error {
	kanbanid := ctx.Params("id")

	var kanban db.Kanban
	if err := db.DB.Where("id = ? AND user_id = ?", kanbanid, auth.UserId(ctx)).First(&kanban).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "kanban not found", err)
	}

	if err := db.DB.Delete(&kanban).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "delete failed", err)
	}

	return ctx.JSON(fiber.Map{"status": "deleted"})
}
