package handler

import (
	"github.com/dlyycious/gopos/internal/module/usermodule"
	"github.com/dlyycious/gopos/model"
	"github.com/dlyycious/gopos/pkg/randomgoidpackage"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(c *fiber.Ctx) error {
	field := new(model.User)
	if err := c.BodyParser(field); err != nil {
		return responsepackage.SendJSON(c, err.Error(), 400, false)
	}
	passwd, _ := bcrypt.GenerateFromPassword([]byte(field.Password), bcrypt.DefaultCost)
	goid := randomgoidpackage.GenerateRandomNumber(13)
	dtos := usermodule.UserDtos{
		Username:  field.Username,
		Password:  string(passwd),
		Fullname:  field.Fullname,
		Email:     field.Email,
		Goid:      uint(goid),
		IsVerfied: true,
		RoleId:    field.RoleId,
	}
	user, err := dtos.Create()
	if err != nil {
		return responsepackage.SendJSON(c, err.Error(), 400, false)
	}
	return responsepackage.SendSuccessJSON(c, "success create new account", user)
}
