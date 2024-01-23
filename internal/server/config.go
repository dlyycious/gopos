package server

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func fiberConfig() *fiber.App {
	return fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
}
