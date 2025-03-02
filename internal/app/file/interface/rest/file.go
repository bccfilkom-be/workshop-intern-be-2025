package rest

import (
	"github.com/ahmdyaasiin/workshop-intern-be-2025/internal/app/file/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type FileHandler struct {
	usecase usecase.FileUsecaseItf
}

func NewFileHandler(routerGroup fiber.Router, usecase usecase.FileUsecaseItf) {
	FileHandler := FileHandler{
		usecase: usecase,
	}

	routerGroup = routerGroup.Group("/files")

	routerGroup.Post("", FileHandler.Upload)
}

func (h *FileHandler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return fiber.ErrBadRequest
	}

	url, err := h.usecase.Upload(file)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success upload file",
		"url":     url,
	})
}
