package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/todo_create.go
func Create(ctx *fiber.Ctx) error {
	var input struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		DueDate     *time.Time `json:"due_date"`
		Priority    string     `json:"priority"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	// Buat todo object untuk validation
	todo := db.Todo{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     input.DueDate,
		Priority:    input.Priority,
		UserID:      auth.UserId(ctx),
	}

	// Validate menggunakan model
	if err := valid.Valid.Struct(todo); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "create failed", err)
	}

	return ctx.JSON(todo)
}
