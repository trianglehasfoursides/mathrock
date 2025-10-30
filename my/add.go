package my

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/url_create.go
func Add(ctx *fiber.Ctx) error {
	var input struct {
		Name string `json:"name"` // nama URL (my.mathrock.xyz/namaini)
		URL  string `json:"url"`  // URL tujuan
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	if input.Name == "" {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "name required", nil)
	}

	if input.URL == "" {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "url required", nil)
	}

	// Check jika name sudah ada
	var existing db.URL
	if err := db.DB.Where("name = ?", input.Name).First(&existing).Error; err == nil {
		return internal.Resperr(ctx, fiber.StatusConflict, "name already exists", nil)
	}

	// Buat short URL
	shorturl := db.URL{
		UserID:      auth.UserId(ctx),
		Name:        input.Name,
		OriginalURL: input.URL,
		Clicks:      0,
	}

	if err := db.DB.Create(&shorturl).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "create failed", err)
	}

	return ctx.JSON(fiber.Map{
		"id":     shorturl.ID,
		"name":   shorturl.Name,
		"url":    "https://my.mathrock.xyz/" + shorturl.Name,
		"target": shorturl.OriginalURL,
		"clicks": shorturl.Clicks,
	})
}
