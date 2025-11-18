package handlers

import (
	"context"
	"fmt"
	"io"
	"jk-api/internal/config"
	"jk-api/internal/helper"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
)

const (
	MAX_IMAGE_SIZE = 10 * 1024 * 1024  // 10MB
	MAX_VIDEO_SIZE = 100 * 1024 * 1024 // 100MB
	MAX_PDF_SIZE   = 10 * 1024 * 1024  // 10MB
	MAX_XLSX_SIZE  = 10 * 1024 * 1024  // 10MB
)

var (
	ALLOWED_FILE_EXTENSIONS = []string{"png", "jpg", "jpeg", "mp4", "pdf", "xlsx"}
	BUCKET_NAME             string
)

func init() {
	BUCKET_NAME = config.GetBucketName()
	if BUCKET_NAME == "" {
		helper.LogMessage("ERROR", "❌ GCP_BUCKET_NAME not found in environment")
	}
}

func GetFileByNameHandler(c *fiber.Ctx, name string) error {
	ctx := context.Background()
	client := config.GetBucketApp()
	obj := client.Bucket(config.GetBucketName()).Object(name)

	rc, err := obj.NewReader(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "File not found")
	}
	defer rc.Close()

	if attrs, err := obj.Attrs(ctx); err == nil {
		c.Set("Content-Type", attrs.ContentType)
	}

	_, err = io.Copy(c.Response().BodyWriter(), rc)
	if err != nil {
		return err
	}

	return nil
}

func UploadFileHandler(data *multipart.FileHeader, c *fiber.Ctx) error {
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(data.Filename), "."))
	if !isAllowedExtension(ext) {
		return fiber.NewError(fiber.StatusBadRequest, "File type not allowed")
	}

	switch ext {
	case "png", "jpg", "jpeg":
		if data.Size > MAX_IMAGE_SIZE {
			return fiber.NewError(fiber.StatusBadRequest, "Image too large")
		}
	case "mp4":
		if data.Size > MAX_VIDEO_SIZE {
			return fiber.NewError(fiber.StatusBadRequest, "Video too large")
		}
	case "pdf":
		if data.Size > MAX_PDF_SIZE {
			return fiber.NewError(fiber.StatusBadRequest, "PDF too large")
		}
	case "xlsx":
		if data.Size > MAX_XLSX_SIZE {
			return fiber.NewError(fiber.StatusBadRequest, "XLSX too large")
		}
	}

	// open file
	file, err := data.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer file.Close()

	// upload ke GCS
	ctx := context.Background()
	client := config.GetBucketApp()
	bucket := client.Bucket(BUCKET_NAME)

	objectName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), data.Filename)
	wc := bucket.Object(objectName).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if err := wc.Close(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{
		"filename":    data.Filename,
		"size":        data.Size,
		"object_name": objectName,
	})
}

func isAllowedExtension(ext string) bool {
	for _, e := range ALLOWED_FILE_EXTENSIONS {
		if e == ext {
			return true
		}
	}
	return false
}

func generateSignedURL(bucket, object string, expiry time.Duration) (string, error) {
	url, err := storage.SignedURL(bucket, object, &storage.SignedURLOptions{
		GoogleAccessID: config.GetServiceAccountEmail(),
		PrivateKey:     config.GetServiceAccountPrivateKey(),
		Method:         "GET",
		Expires:        time.Now().Add(expiry),
	})
	if err != nil {
		return "", err
	}
	return url, nil
}
