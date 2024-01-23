package endpoint

import (
	"github.com/dlyycious/gopos/api/handler"
	"github.com/dlyycious/gopos/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthEndpoint(app *fiber.App) {
	router := app.Group("/auth")
	router.Post("/", handler.LoginHandlers)
	router.Delete("/", middleware.AuthMiddleware, handler.LogoutHandler)
}
