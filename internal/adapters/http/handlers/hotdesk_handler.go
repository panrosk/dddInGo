package handlers

import (
	"coworking/internal/adapters/http/models"
	"coworking/internal/app/usecases/hotdesk"

	"github.com/gofiber/fiber/v2"
)

type HotdeskHandler struct {
	registerCommand *hotdesk.RegisterHotdeskUsecase
}

func NewHotdeskHandler(registerCommand *hotdesk.RegisterHotdeskUsecase) *HotdeskHandler {
	return &HotdeskHandler{registerCommand: registerCommand}
}

func (h *HotdeskHandler) RegisterHotdesk(c *fiber.Ctx) error {
	var req models.HotdeskDTO

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Invalid request body",
		})
	}

	params := hotdesk.RegisterHotdeskParams{
		Number: req.Number,
	}

	hotdesk, err := h.registerCommand.Execute(params)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Failed to register hotdesk",
		})
	}

	return c.Status(200).JSON(hotdesk)
}
