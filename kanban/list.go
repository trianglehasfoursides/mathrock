package kanban

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/kanban_list.go
func List(ctx *fiber.Ctx) error {
	var kanbans []db.Kanban

	// Group by status dengan order position
	if err := db.DB.Where("user_id = ?", auth.UserId(ctx)).
		Order("status, position ASC").
		Find(&kanbans).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}

	return ctx.JSON(kanbans)
}
