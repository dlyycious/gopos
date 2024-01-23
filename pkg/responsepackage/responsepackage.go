package responsepackage

import (
	"github.com/gofiber/fiber/v2"
)

func SendJSON(c *fiber.Ctx, message string, code int, status ...bool) error {
	var stat bool

	// Check if status is provided, otherwise use default value (true)
	if len(status) > 0 {
		stat = status[0]
	} else {
		stat = true
	}
	return c.Status(code).JSON(fiber.Map{
		"status":  stat,
		"message": message,
	})
}

func SendSuccessJSON(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  true,
		"message": message,
		"data":    data,
	})
}

func SendErrorJSON(c *fiber.Ctx, message string, errors interface{}) error {
	return c.Status(400).JSON(fiber.Map{
		"status":  false,
		"message": message,
		"errors":  errors,
	})
}
