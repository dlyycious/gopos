package model

type RolePermission struct {
	Id         uint `gorm:"primaryKey"`
	RoleId     uint
	Permission uint
	Role       Role `gorm:"foreignKey:RoleId"`
}
