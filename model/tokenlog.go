package model

import "time"

type TokenLog struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint
	User      User   `gorm:"foreignKey:UserId"`
	Token     string `gorm:"varchar(300); index"`
	ExpiresAt time.Time
}
