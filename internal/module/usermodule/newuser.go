package usermodule

import (
	"github.com/dlyycious/gopos/database"
	"github.com/dlyycious/gopos/model"
	"github.com/dlyycious/gopos/pkg/randomgoidpackage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var goidGenerate uint = randomgoidpackage.GenerateRandomNumber(13)

func (u *UserDtos) Create() (UserDtos, error) {
	db := database.DB
	query := model.User{
		Username: u.Username,
		Goid:     goidGenerate,
	}
	found := model.User{}
	if err := db.Where("username = ?", query.Username).First(&found).Error; err != gorm.ErrRecordNotFound {
		return UserDtos{}, fiber.ErrBadRequest
	} else if err := db.Where("goid = ?", query.Goid).First(&found).Error; err != gorm.ErrRecordNotFound {
		goidGenerate = randomgoidpackage.GenerateRandomNumber(13)
	}
	create := model.User{
		Username:  u.Username,
		Password:  u.Password,
		Fullname:  u.Fullname,
		Email:     u.Email,
		Goid:      goidGenerate,
		RoleId:    u.RoleId,
		IsVerfied: u.IsVerfied,
	}
	db.Create(&create)

	queryRole := model.Role{Id: create.RoleId}
	foundrole := model.Role{}
	if err := db.Where("id = ?", queryRole.Id).First(&foundrole).Error; err != nil {
		return UserDtos{}, fiber.ErrBadRequest
	}

	res := UserDtos{
		Username:  u.Username,
		Fullname:  u.Fullname,
		Email:     u.Email,
		Goid:      u.Goid,
		IsVerfied: u.IsVerfied,
		Role:      foundrole.Role,
	}
	return res, nil
}
