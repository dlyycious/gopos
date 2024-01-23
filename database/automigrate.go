package database

import (
	"github.com/dlyycious/gopos/model"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Role{},
		&model.RolePermission{},
		&model.User{},
		&model.TokenLog{},
	)
}
