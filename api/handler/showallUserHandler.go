package handler

import (
	"github.com/dlyycious/gopos/database"
	"github.com/dlyycious/gopos/internal/module/usermodule"
	"github.com/dlyycious/gopos/model"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
)

func ShowAllUserHandler(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User
	if err := db.Preload("Role").Find(&users).Error; err != nil {
		return responsepackage.SendJSON(c, err.Error(), 400, false)
	}

	var userDtos []usermodule.UserDtos
	for _, u := range users {
		userDto := usermodule.UserDtos{
			Username:  u.Username,
			Fullname:  u.Fullname,
			Email:     u.Email,
			Goid:      u.Goid,
			IsVerfied: u.IsVerfied,
			Role:      u.Role.Role,
		}
		userDtos = append(userDtos, userDto)
	}

	return responsepackage.SendSuccessJSON(c, "success showing all users", userDtos)
}
