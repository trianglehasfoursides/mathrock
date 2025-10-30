// handlers/mood.go
package mood

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/valid"
)

func Add(ctx *fiber.Ctx) error {
	var input struct {
		Mood        string    `json:"mood"`
		Intensity   int       `json:"intensity"`
		Description string    `json:"description"`
		Date        time.Time `json:"date"`
		Tags        string    `json:"tags"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	if input.Date.IsZero() {
		input.Date = time.Now()
	}

	mood := db.Mood{
		UserID:      auth.UserId(ctx),
		Mood:        input.Mood,
		Intensity:   input.Intensity,
		Description: input.Description,
		Date:        input.Date,
		Tags:        input.Tags,
	}

	if err := valid.Valid.Struct(mood); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	var existing db.Mood
	if err := db.DB.Where("user_id = ? AND DATE(date) = DATE(?)", auth.UserId(ctx), input.Date).First(&existing).Error; err == nil {
		return internal.Resperr(ctx, fiber.StatusConflict, "mood entry already exists for this date", nil)
	}

	if err := db.DB.Create(&mood).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "create failed", err)
	}

	return ctx.JSON(mood)
}

func List(ctx *fiber.Ctx) error {
	var moods []db.Mood

	query := db.DB.Where("user_id = ?", auth.UserId(ctx))

	if from := ctx.Query("from"); from != "" {
		fromDate, _ := time.Parse("2006-01-02", from)
		query = query.Where("DATE(date) >= ?", fromDate)
	}

	if to := ctx.Query("to"); to != "" {
		toDate, _ := time.Parse("2006-01-02", to)
		query = query.Where("DATE(date) <= ?", toDate)
	}

	if mood := ctx.Query("mood"); mood != "" {
		query = query.Where("mood = ?", mood)
	}

	if err := query.Order("date DESC").Find(&moods).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}

	return ctx.JSON(moods)
}

func Get(ctx *fiber.Ctx) error {
	moodid := ctx.Params("id")

	var mood db.Mood
	if err := db.DB.Where("id = ? AND user_id = ?", moodid, auth.UserId(ctx)).First(&mood).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "mood not found", err)
	}

	return ctx.JSON(mood)
}

func Update(ctx *fiber.Ctx) error {
	moodid := ctx.Params("id")

	var mood db.Mood
	if err := db.DB.Where("id = ? AND user_id = ?", moodid, auth.UserId(ctx)).First(&mood).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "mood not found", err)
	}

	var input struct {
		Mood        string    `json:"mood"`
		Intensity   int       `json:"intensity"`
		Description string    `json:"description"`
		Date        time.Time `json:"date"`
		Tags        string    `json:"tags"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	if input.Mood != "" {
		mood.Mood = input.Mood
	}
	if input.Intensity > 0 {
		mood.Intensity = input.Intensity
	}
	if input.Description != "" {
		mood.Description = input.Description
	}
	if !input.Date.IsZero() {
		mood.Date = input.Date
	}
	if input.Tags != "" {
		mood.Tags = input.Tags
	}

	if err := valid.Valid.Struct(mood); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "validation failed", err)
	}

	if err := db.DB.Save(&mood).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "update failed", err)
	}

	return ctx.JSON(mood)
}

func Delete(ctx *fiber.Ctx) error {
	moodid := ctx.Params("id")

	var mood db.Mood
	if err := db.DB.Where("id = ? AND user_id = ?", moodid, auth.UserId(ctx)).First(&mood).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "mood not found", err)
	}

	if err := db.DB.Delete(&mood).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "delete failed", err)
	}

	return ctx.JSON(fiber.Map{"status": "deleted"})
}

func Stats(ctx *fiber.Ctx) error {
	var stats []struct {
		Mood         string  `json:"mood"`
		Count        int     `json:"count"`
		AvgIntensity float64 `json:"avg_intensity"`
	}

	query := `
		SELECT mood, COUNT(*) as count, AVG(intensity) as avg_intensity
		FROM moods 
		WHERE user_id = ? 
		GROUP BY mood 
		ORDER BY count DESC
	`

	if err := db.DB.Raw(query, auth.UserId(ctx)).Scan(&stats).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "stats failed", err)
	}

	var trend []struct {
		Date  string `json:"date"`
		Mood  string `json:"mood"`
		Count int    `json:"count"`
	}

	trendQuery := `
		SELECT DATE(date) as date, mood, COUNT(*) as count
		FROM moods 
		WHERE user_id = ? AND date >= DATE('now', '-7 days')
		GROUP BY DATE(date), mood
		ORDER BY date DESC
	`

	db.DB.Raw(trendQuery, auth.UserId(ctx)).Scan(&trend)

	return ctx.JSON(fiber.Map{
		"mood_stats":   stats,
		"weekly_trend": trend,
	})
}

func Today(ctx *fiber.Ctx) error {
	today := time.Now().Format("2006-01-02")

	var mood db.Mood
	if err := db.DB.Where("user_id = ? AND DATE(date) = ?", auth.UserId(ctx), today).First(&mood).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusNotFound, "no mood entry for today", nil)
	}

	return ctx.JSON(mood)
}
