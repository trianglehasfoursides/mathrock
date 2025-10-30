package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/todo_get.go
func Get(ctx *fiber.Ctx) error {
	todoid := ctx.Params("id")

	var todo db.Todo
	if err := db.DB.Where("id = ? AND user_id = ?", todoid, auth.UserId(ctx)).First(&todo).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "todo not found", err)
	}

	return ctx.JSON(todo)
}
