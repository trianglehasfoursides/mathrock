package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/todo_delete.go
func Delete(ctx *fiber.Ctx) error {
	todoid := ctx.Params("id")

	var todo db.Todo
	if err := db.DB.Where("id = ? AND user_id = ?", todoid, auth.UserId(ctx)).First(&todo).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "todo not found", err)
	}

	if err := db.DB.Delete(&todo).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "delete failed", err)
	}

	return ctx.JSON(fiber.Map{"status": "deleted"})
}
