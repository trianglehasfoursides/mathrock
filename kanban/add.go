package kanban

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/valid"
)

// handlers/kanban_create.go
func Add(ctx *fiber.Ctx) error {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Color       string `json:"color"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	// Get next position untuk status ini
	var maxPosition int
	db.DB.Model(&db.Kanban{}).Where("user_id = ? AND status = ?", auth.UserId(ctx), input.Status).
		Select("COALESCE(MAX(position), 0)").Scan(&maxPosition)

	kanban := db.Kanban{
		UserID:      auth.UserId(ctx),
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		Position:    maxPosition + 1,
		Color:       input.Color,
	}

	// Validate
	if err := valid.Valid.Struct(kanban); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Create(&kanban).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "create failed", err)
	}

	return ctx.JSON(kanban)
}
