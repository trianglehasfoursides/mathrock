package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/todo_list.go
func List(ctx *fiber.Ctx) error {
	var todos []db.Todo

	query := db.DB.Where("user_id = ?", auth.UserId(ctx))

	// Filter by status
	if completed := ctx.Query("completed"); completed != "" {
		query = query.Where("completed = ?", completed == "true")
	}

	// Filter by priority
	if priority := ctx.Query("priority"); priority != "" {
		query = query.Where("priority = ?", priority)
	}

	// Sort by priority and due date
	if err := query.Order("completed ASC, priority DESC, due_date ASC").Find(&todos).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}

	return ctx.JSON(todos)
}
