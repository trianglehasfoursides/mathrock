package todo

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/valid"
)

// handlers/todo_update.go
func Update(ctx *fiber.Ctx) error {
	todoid := ctx.Params("id")

	var todo db.Todo
	if err := db.DB.Where("id = ? AND user_id = ?", todoid, auth.UserId(ctx)).First(&todo).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "todo not found", err)
	}

	var input struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Completed   bool       `json:"completed"`
		DueDate     *time.Time `json:"due_date"`
		Priority    string     `json:"priority"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	// Update fields
	if input.Title != "" {
		todo.Title = input.Title
	}
	if input.Description != "" {
		todo.Description = input.Description
	}
	todo.Completed = input.Completed
	todo.DueDate = input.DueDate
	if input.Priority != "" {
		todo.Priority = input.Priority
	}

	// Validate sebelum save
	if err := valid.Valid.Struct(todo); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Save(&todo).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "update failed", err)
	}

	return ctx.JSON(todo)
}
