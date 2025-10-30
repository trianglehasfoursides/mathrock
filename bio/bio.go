// handlers/linkinbio.go
package bio

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/valid"
)

// Setup page user (hanya sekali)
func Setup(ctx *fiber.Ctx) error {
	var input struct {
		Username    string `json:"username"`
		DisplayName string `json:"display_name"`
		Bio         string `json:"bio"`
		Avatar      string `json:"avatar"`
		Theme       string `json:"theme"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	// Check jika user sudah punya page
	var existing db.LinkInBio
	if err := db.DB.Where("user_id = ?", auth.UserId(ctx)).First(&existing).Error; err == nil {
		return internal.Resperr(ctx, fiber.StatusConflict, "page already exists", nil)
	}

	// Check jika username sudah dipakai
	if err := db.DB.Where("username = ?", input.Username).First(&existing).Error; err == nil {
		return internal.Resperr(ctx, fiber.StatusConflict, "username already taken", nil)
	}

	page := db.LinkInBio{
		UserID:      auth.UserId(ctx),
		Username:    input.Username,
		DisplayName: input.DisplayName,
		Bio:         input.Bio,
		Avatar:      input.Avatar,
		Theme:       input.Theme,
		Active:      true,
	}

	if err := valid.Valid.Struct(page); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Create(&page).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "setup failed", err)
	}

	return ctx.JSON(page)
}

// Get page user sendiri
func GetPage(ctx *fiber.Ctx) error {
	var page db.LinkInBio
	if err := db.DB.Where("user_id = ?", auth.UserId(ctx)).First(&page).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "page not found", err)
	}

	return ctx.JSON(page)
}

// Update page user
func UpdatePage(ctx *fiber.Ctx) error {
	var page db.LinkInBio
	if err := db.DB.Where("user_id = ?", auth.UserId(ctx)).First(&page).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "page not found", err)
	}

	var input struct {
		DisplayName string `json:"display_name"`
		Bio         string `json:"bio"`
		Avatar      string `json:"avatar"`
		Theme       string `json:"theme"`
		Active      bool   `json:"active"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	if input.DisplayName != "" {
		page.DisplayName = input.DisplayName
	}
	if input.Bio != "" {
		page.Bio = input.Bio
	}
	if input.Avatar != "" {
		page.Avatar = input.Avatar
	}
	if input.Theme != "" {
		page.Theme = input.Theme
	}
	page.Active = input.Active

	if err := valid.Valid.Struct(page); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Save(&page).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "update failed", err)
	}

	return ctx.JSON(page)
}

// Add link ke page
func AddLink(ctx *fiber.Ctx) error {
	var input struct {
		Title       string `json:"title"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		Color       string `json:"color"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	// Get next position
	var maxPosition int
	db.DB.Model(&db.LinkInBioItem{}).Where("user_id = ?", auth.UserId(ctx)).
		Select("COALESCE(MAX(position), 0)").Scan(&maxPosition)

	link := db.LinkInBioItem{
		UserID:      auth.UserId(ctx),
		Title:       input.Title,
		URL:         input.URL,
		Description: input.Description,
		Icon:        input.Icon,
		Color:       input.Color,
		Position:    maxPosition + 1,
		Active:      true,
	}

	if err := valid.Valid.Struct(link); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Create(&link).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "create failed", err)
	}

	return ctx.JSON(link)
}

// List links di page user
func ListLinks(ctx *fiber.Ctx) error {
	var links []db.LinkInBioItem

	if err := db.DB.Where("user_id = ? AND active = ?", auth.UserId(ctx), true).
		Order("position ASC").
		Find(&links).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}

	return ctx.JSON(links)
}

// Update link
func UpdateLink(ctx *fiber.Ctx) error {
	linkid := ctx.Params("id")

	var link db.LinkInBioItem
	if err := db.DB.Where("id = ? AND user_id = ?", linkid, auth.UserId(ctx)).First(&link).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "link not found", err)
	}

	var input struct {
		Title       string `json:"title"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		Color       string `json:"color"`
		Active      bool   `json:"active"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	if input.Title != "" {
		link.Title = input.Title
	}
	if input.URL != "" {
		link.URL = input.URL
	}
	if input.Description != "" {
		link.Description = input.Description
	}
	if input.Icon != "" {
		link.Icon = input.Icon
	}
	if input.Color != "" {
		link.Color = input.Color
	}
	link.Active = input.Active

	if err := valid.Valid.Struct(link); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Save(&link).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "update failed", err)
	}

	return ctx.JSON(link)
}

// Delete link
func DeleteLink(ctx *fiber.Ctx) error {
	linkid := ctx.Params("id")

	var link db.LinkInBioItem
	if err := db.DB.Where("id = ? AND user_id = ?", linkid, auth.UserId(ctx)).First(&link).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "link not found", err)
	}

	if err := db.DB.Delete(&link).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "delete failed", err)
	}

	return ctx.JSON(fiber.Map{"status": "deleted"})
}

// Reorder links
func ReorderLinks(ctx *fiber.Ctx) error {
	var input struct {
		Links []struct {
			ID       uint `json:"id"`
			Position int  `json:"position"`
		} `json:"links"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	for _, item := range input.Links {
		db.DB.Model(&db.LinkInBioItem{}).
			Where("id = ? AND user_id = ?", item.ID, auth.UserId(ctx)).
			Update("position", item.Position)
	}

	return ctx.JSON(fiber.Map{"status": "reordered"})
}

// Public: Get page by username
func GetPublicPage(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	var page db.LinkInBio
	if err := db.DB.Where("username = ? AND active = ?", username, true).First(&page).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "page not found", err)
	}

	// Get active links
	var links []db.LinkInBioItem
	db.DB.Where("user_id = ? AND active = ?", page.UserID, true).
		Order("position ASC").
		Find(&links)

	return ctx.JSON(fiber.Map{
		"page":  page,
		"links": links,
	})
}
