package my

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/url_list.go
func List(ctx *fiber.Ctx) error {
	var shorturls []db.URL
	if err := db.DB.Where("user_id = ?", auth.UserId(ctx)).Find(&shorturls).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}

	return ctx.JSON(shorturls)
}
