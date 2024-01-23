package handler

import (
	"strings"

	"github.com/dlyycious/gopos/internal/module/authmodule"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	tokens := strings.TrimPrefix(authHeader, "Bearer ")
	jwtsplit := strings.Split(tokens, ".")[2]

	_, err := authmodule.Logout(jwtsplit)

	if err != nil {
		return responsepackage.SendJSON(c, err.Error(), 400, false)
	}

	return responsepackage.SendJSON(c, "logout success", 200, true)
}
