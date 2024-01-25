package handler

import (
	"github.com/dlyycious/gopos/internal/module/usermodule"
	"github.com/dlyycious/gopos/model"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeleteUserHandler(c *fiber.Ctx) error {
	field := new(model.User)
	if err := c.BodyParser(field); err != nil {
		return responsepackage.SendJSON(c, err.Error(), 500, false)
	}

	delete, err := usermodule.DeleteUser(field.Goid)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return responsepackage.SendJSON(c, "User not Found", 400, false)
		default:
			return responsepackage.SendJSON(c, err.Error(), 500, false)
		}
	}

	return responsepackage.SendSuccessJSON(c, "Success Delete User", delete)
}
