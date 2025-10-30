package kanban

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/valid"
)

// handlers/kanban_update.go
func Update(ctx *fiber.Ctx) error {
	kanbanid := ctx.Params("id")

	var kanban db.Kanban
	if err := db.DB.Where("id = ? AND user_id = ?", kanbanid, auth.UserId(ctx)).First(&kanban).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "kanban not found", err)
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Color       string `json:"color"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	// Update fields
	if input.Title != "" {
		kanban.Title = input.Title
	}
	if input.Description != "" {
		kanban.Description = input.Description
	}
	if input.Status != "" {
		kanban.Status = input.Status
	}
	if input.Color != "" {
		kanban.Color = input.Color
	}

	// Validate
	if err := valid.Valid.Struct(kanban); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Save(&kanban).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "update failed", err)
	}

	return ctx.JSON(kanban)
}
