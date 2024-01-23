package authmodule

import (
	"time"

	"github.com/dlyycious/gopos/database"
	"github.com/dlyycious/gopos/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func VerifiedToken(token string) (UserDtos, error) {
	db := database.DB

	dtos := model.TokenLog{
		Token: token,
	}
	tokens := model.TokenLog{}
	if err := db.Preload("User").Model(dtos).Where("token = ?", dtos.Token).First(&tokens).Error; err != nil {
		return UserDtos{}, gorm.ErrRecordNotFound
	}

	if tokens.ExpiresAt.Before(time.Now()) {
		return UserDtos{}, fiber.ErrBadRequest
	}

	query := model.User{Id: tokens.UserId}
	user := model.User{}
	if err := db.Preload("Role").Model(query).Where("id = ?", query.Id).First(&user).Error; err != nil {
		return UserDtos{}, gorm.ErrRecordNotFound
	}

	return UserDtos{
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Goid:      user.Goid,
		IsVerfied: user.IsVerfied,
		CreatedAt: user.CreatedAt,
		Role:      user.Role.Role,
	}, nil

}
