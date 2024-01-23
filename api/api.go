package api

import (
	"github.com/dlyycious/gopos/api/endpoint"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	endpoint.AuthEndpoint(app)
	endpoint.UserEndpoint(app)
}
