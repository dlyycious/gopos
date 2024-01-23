package authmodule

import (
	"strings"
	"time"

	"github.com/dlyycious/gopos/config"
	"github.com/dlyycious/gopos/database"
	"github.com/dlyycious/gopos/model"
	"github.com/dlyycious/gopos/pkg/jwtpackage"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Login authenticates a user based on the provided login credentials.
// It returns a UserDtos and an error. If the authentication is successful,
// it returns the user information as UserDtos. Otherwise, it returns an error.
func (dtos *LoginDtos) Login() (LoginDtos, error) {
	db := database.DB
	query := model.User{
		Username: dtos.Username,
		Password: dtos.Password,
	}
	user := model.User{}

	// Query the database to find the user by username
	if err := db.Preload("Role").Model(&query).Where("username = ?", query.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			// User not found in the database
			return LoginDtos{}, gorm.ErrRecordNotFound
		default:
			// Other database error
			return LoginDtos{}, err
		}
	}

	// Check if the provided password matches the stored hashed password
	if !PasswordCheck([]byte(user.Password), []byte(dtos.Password)) {
		// Password does not match
		return LoginDtos{}, fiber.ErrBadRequest
	}

	jwt := jwtpackage.JWTClaims{
		Username: user.Username,
		Goid:     user.Goid,
	}

	token, errJWT := jwt.Generate()

	if errJWT != nil {
		return LoginDtos{}, errJWT
	}

	checkToken := strings.Split(token, ".")[2]
	tokens := model.TokenLog{
		UserId:    user.Id,
		Token:     checkToken,
		ExpiresAt: time.Now().Add(config.JWT_EXP * time.Hour),
	}

	if err := db.Create(&tokens).Error; err != nil {
		return LoginDtos{}, err
	}

	// Map user information to UserDtos for response
	loginDtos := LoginDtos{
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Goid:      user.Goid,
		IsVerfied: user.IsVerfied,
		Token:     token,
		Role:      user.Role.Role,
		CreatedAt: user.CreatedAt,
	}

	return loginDtos, nil
}
