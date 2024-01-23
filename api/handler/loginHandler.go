package handler

import (
	"github.com/dlyycious/gopos/internal/module/authmodule"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginHandlers(c *fiber.Ctx) error {
	dtos := new(authmodule.LoginDtos)

	if err := c.BodyParser(dtos); err != nil {
		return responsepackage.SendJSON(c, err.Error(), 400, false)
	}

	user, err := dtos.Login()

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return responsepackage.SendJSON(c, "username not found", 400, false)
		case fiber.ErrBadRequest:
			return responsepackage.SendJSON(c, "password is invalid", 400, false)
		default:
			return responsepackage.SendJSON(c, err.Error(), 400, false)
		}
	}
	data := fiber.Map{
		"username":    user.Username,
		"fullname":    user.Fullname,
		"email":       user.Email,
		"goid":        user.Goid,
		"is_verified": user.IsVerfied,
		"role":        user.Role,
		"created_at":  user.CreatedAt,
		"token":       user.Token,
	}
	return responsepackage.SendSuccessJSON(c, "login successfully", data)
}
