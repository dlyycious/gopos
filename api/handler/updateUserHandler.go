package handler

import (
	"github.com/dlyycious/gopos/internal/module/usermodule"
	"github.com/dlyycious/gopos/model"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UpdateUserHandler(c *fiber.Ctx) error {
	field := new(model.User)
	if err := c.BodyParser(field); err != nil {
		return responsepackage.SendJSON(c, err.Error(), 400, false)
	}

	passwd, _ := bcrypt.GenerateFromPassword([]byte(field.Password), bcrypt.DefaultCost)
	dtos := usermodule.UserDtos{
		Password: string(passwd),
		Fullname: field.Fullname,
		Email:    field.Email,
		RoleId:   field.RoleId,
	}

	user, err := dtos.Update(field.Goid)

	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return responsepackage.SendJSON(c, "user not found", 400, false)
		default:
			return responsepackage.SendJSON(c, err.Error(), 400, false)
		}
	}
	return responsepackage.SendSuccessJSON(c, "success update user", user)
}
