package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func formatErrorResponse(c *fiber.Ctx, statusCode int, message string, details interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   message,
		"details": details,
	})
}

func formatSuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"data":    data,
		"message": message,
	})
}
