package box

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
)

// handlers/file_duplicates.go
func Duplicates(ctx *fiber.Ctx) error {
	var duplicates []struct {
		Hash      string   `json:"hash"`
		Count     int      `json:"count"`
		FileIDs   []uint   `json:"file_ids"`
		FileNames []string `json:"file_names"`
	}

	// Cari hash yang duplikat
	query := `
		SELECT hash, COUNT(*) as count, 
		       GROUP_CONCAT(id) as file_ids,
		       GROUP_CONCAT(name) as file_names
		FROM files 
		WHERE user_id = ? 
		GROUP BY hash 
		HAVING count > 1
	`

	rows, err := db.DB.Raw(query, auth.UserId(ctx)).Rows()
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "query failed", err)
	}
	defer rows.Close()

	for rows.Next() {
		var dup struct {
			Hash      string
			Count     int
			FileIDs   string
			FileNames string
		}

		rows.Scan(&dup.Hash, &dup.Count, &dup.FileIDs, &dup.FileNames)

		// Parse file IDs dan names
		ids := parseUintSlice(dup.FileIDs)
		names := parseStringSlice(dup.FileNames)

		duplicates = append(duplicates, struct {
			Hash      string   `json:"hash"`
			Count     int      `json:"count"`
			FileIDs   []uint   `json:"file_ids"`
			FileNames []string `json:"file_names"`
		}{
			Hash:      dup.Hash,
			Count:     dup.Count,
			FileIDs:   ids,
			FileNames: names,
		})
	}

	return ctx.JSON(fiber.Map{"duplicates": duplicates})
}

// Helper functions
func parseUintSlice(s string) []uint {
	parts := strings.Split(s, ",")
	var result []uint
	for _, part := range parts {
		if id, err := strconv.ParseUint(part, 10, 32); err == nil {
			result = append(result, uint(id))
		}
	}
	return result
}

func parseStringSlice(s string) []string {
	return strings.Split(s, ",")
}
