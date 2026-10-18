package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"jk-api/api/http/presenters"

	"github.com/gofiber/fiber/v2"
)

const uploadMaxBytes = 5 * 1024 * 1024 // 5 MB

var allowedImageExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
	".gif":  true,
}

func uploadDir() string {
	dir := os.Getenv("UPLOAD_DIR")
	if dir == "" {
		dir = "./uploads"
	}
	return dir
}

// UploadFile menyimpan gambar ke disk lokal dan mengembalikan URL relatif.
func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "File wajib diunggah")
	}

	if file.Size > uploadMaxBytes {
		return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Ukuran file maksimal 5 MB")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedImageExt[ext] {
		return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Format harus jpg, jpeg, png, webp, atau gif")
	}

	dir := uploadDir()
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	name := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	if err := c.SaveFile(file, filepath.Join(dir, name)); err != nil {
		return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return presenters.SuccessResponse(c, fiber.Map{
		"url":      "/uploads/" + name,
		"filename": name,
	})
}
