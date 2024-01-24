package usermodule

import (
	"github.com/dlyycious/gopos/database"
	"github.com/dlyycious/gopos/model"
	"gorm.io/gorm"
)

func DeleteUser(goid uint) (UserDtos, error) {
	db := database.DB

	query := model.User{Goid: goid}
	found := model.User{}

	if err := db.Preload("Role").Where("goid = ?", query.Goid).First(&found).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return UserDtos{}, gorm.ErrRecordNotFound
		default:
			return UserDtos{}, err
		}
	}

	if err := db.Delete(&found).Error; err != nil {
		return UserDtos{}, err
	}

	res := UserDtos{
		Username: found.Username,
		Fullname: found.Fullname,
		Email:    found.Email,
		Goid:     found.Goid,
		Role:     found.Role.Role,
	}

	return res, nil
}
