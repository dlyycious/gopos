package model

type Role struct {
	Id             uint `gorm:"primaryKey"`
	Role           string
	RolePermission []RolePermission
}
