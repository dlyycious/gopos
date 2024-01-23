package handler

import (
	"github.com/dlyycious/gopos/internal/module/authmodule"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
)

func GetInfoHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(authmodule.UserDtos)
	return responsepackage.SendSuccessJSON(c, "success get info", user)
}
