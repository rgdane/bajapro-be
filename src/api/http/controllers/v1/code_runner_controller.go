package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"jk-api/api/http/presenters"

	"github.com/gofiber/fiber/v2"
)

type runCodeRequest struct {
	Language string `json:"language"`
	Version  string `json:"version"`
	Code     string `json:"code"`
	Stdin    string `json:"stdin"`
}

// RunCode meneruskan kode ke Piston (sandbox eksekusi) dan mengembalikan hasilnya.
// Piston dijalankan terpisah (Docker), URL-nya lewat env PISTON_URL.
func RunCode(c *fiber.Ctx) error {
	var input runCodeRequest
	if err := c.BodyParser(&input); err != nil {
		return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Invalid request")
	}

	if strings.TrimSpace(input.Code) == "" {
		return presenters.ErrorResponseWithMessage(c, fiber.StatusBadRequest, "Kode tidak boleh kosong")
	}
	if input.Language == "" {
		input.Language = "java"
	}
	if input.Version == "" {
		input.Version = "*"
	}

	fileName := "Main.java"
	if input.Language != "java" {
		fileName = "main.txt"
	}

	payload := map[string]any{
		"language": input.Language,
		"version":  input.Version,
		"files": []map[string]string{
			{"name": fileName, "content": input.Code},
		},
		"stdin": input.Stdin,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	base := os.Getenv("PISTON_URL")
	if base == "" {
		base = "http://localhost:2000"
	}
	endpoint := strings.TrimRight(base, "/") + "/api/v2/execute"

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Post(endpoint, "application/json", bytes.NewReader(body))
	if err != nil {
		return presenters.ErrorResponseWithMessage(
			c,
			fiber.StatusBadGateway,
			"Sandbox tidak dapat dihubungi. Pastikan Piston berjalan.",
		)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return presenters.ErrorResponse(c, fiber.StatusBadGateway, err)
	}

	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		return presenters.ErrorResponseWithMessage(
			c,
			fiber.StatusBadGateway,
			"Respons sandbox tidak valid",
		)
	}

	if resp.StatusCode >= 400 {
		msg := "Eksekusi gagal"
		if m, ok := result["message"].(string); ok && m != "" {
			msg = m
		}
		return presenters.ErrorResponseWithMessage(c, resp.StatusCode, msg)
	}

	return presenters.SuccessResponse(c, result)
}
