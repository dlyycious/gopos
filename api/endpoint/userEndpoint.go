package endpoint

import (
	"github.com/dlyycious/gopos/api/handler"
	"github.com/dlyycious/gopos/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserEndpoint(app *fiber.App) {
	router := app.Group("/user")
	router.Use(middleware.AuthMiddleware)
	router.Get("/", handler.GetInfoHandler)
}
