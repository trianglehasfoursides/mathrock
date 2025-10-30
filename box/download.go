package box

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/trianglehasfoursides/mathrock/auth"
	"github.com/trianglehasfoursides/mathrock/db"
	"github.com/trianglehasfoursides/mathrock/internal"
	"github.com/trianglehasfoursides/mathrock/storage"
)

// handlers/file_download_url.go
func DownloadURL(ctx *fiber.Ctx) error {
	var input struct {
		URL      string `json:"url"`
		Filename string `json:"filename"`
		Encrypt  bool   `json:"encrypt"`
	}

	if err := ctx.BodyParser(&input); err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "invalid input", err)
	}

	if input.URL == "" {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "url required", nil)
	}

	if input.Filename == "" {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "filename required", nil)
	}

	// Check jika filename sudah ada
	var existingFile db.File
	if err := db.DB.Where("user_id = ? AND name = ?", auth.UserId(ctx), input.Filename).First(&existingFile).Error; err == nil {
		return internal.Resperr(ctx, fiber.StatusConflict, "filename already exists", nil)
	}

	// Download dari URL
	resp, err := http.Get(input.URL)
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "download failed", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return internal.Resperr(ctx, fiber.StatusBadRequest, "download failed: "+resp.Status, nil)
	}

	// Baca content ke memory
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "read failed", err)
	}

	// Create temporary file
	tempfile, err := os.CreateTemp("", "download-*")
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "create temp failed", err)
	}
	defer os.Remove(tempfile.Name())

	// Write content ke temp file
	if _, err := tempfile.Write(content); err != nil {
		tempfile.Close()
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "write failed", err)
	}
	tempfile.Close()

	// Buka lagi sebagai file
	file, err := os.Open(tempfile.Name())
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "open failed", err)
	}
	defer file.Close()

	// Get file info
	fileinfo, err := file.Stat()
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "stat failed", err)
	}

	// Create multipart file header
	fileheader := &multipart.FileHeader{
		Filename: input.Filename,
		Size:     fileinfo.Size(),
	}

	// Upload ke storage
	hash, err := storage.Box.Upload(fileheader)
	if err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "upload failed", err)
	}

	// Simpan ke database
	filedata := db.File{
		Name:      input.Filename,
		Hash:      hash,
		UserID:    auth.UserId(ctx),
		Encrypt:   input.Encrypt,
		SourceURL: input.URL,
	}

	if err := db.DB.Create(&filedata).Error; err != nil {
		return internal.Resperr(ctx, fiber.StatusInternalServerError, "save failed", err)
	}

	return ctx.JSON(fiber.Map{
		"id":   filedata.ID,
		"name": filedata.Name,
		"hash": filedata.Hash,
		"url":  input.URL,
	})
}
