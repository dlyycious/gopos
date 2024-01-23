package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint   `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar(25)"`
	Password  string `gorm:"type:text"`
	Fullname  string `gorm:"type:varchar(25)"`
	Email     string `gorm:"type:varchar(50)"`
	Goid      uint   `gorm:"index"`
	RoleId    uint
	IsVerfied bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      Role           `gorm:"foreignKey:RoleId"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	TokenLog  []TokenLog
}
