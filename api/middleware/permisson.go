package middleware

import (
	"github.com/dlyycious/gopos/database"
	"github.com/dlyycious/gopos/internal/constain"
	"github.com/dlyycious/gopos/internal/module/authmodule"
	"github.com/dlyycious/gopos/model"
	"github.com/dlyycious/gopos/pkg/responsepackage"
	"github.com/gofiber/fiber/v2"
)

func HasManageUser(c *fiber.Ctx) error {
	db := database.DB
	locals := c.Locals("user").(authmodule.UserDtos)
	queryRole := model.Role{Role: locals.Role}
	foundRole := model.Role{}
	if err := db.Model(queryRole).Where("role = ?", queryRole.Role).First(&foundRole).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}
	permission := constain.Manage_User
	queryRolePermission := model.RolePermission{
		RoleId:     foundRole.Id,
		Permission: permission,
	}
	foundRolePermission := model.RolePermission{}
	if err := db.Where("permission = ?", queryRolePermission.Permission).Where("role_id = ?", queryRolePermission.RoleId).First(&foundRolePermission).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}

	return c.Next()
}

func HasManageTransaction(c *fiber.Ctx) error {
	db := database.DB
	locals := c.Locals("user").(authmodule.UserDtos)
	queryRole := model.Role{Role: locals.Role}
	foundRole := model.Role{}
	if err := db.Model(queryRole).Where("role = ?", queryRole.Role).First(&foundRole).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}
	permission := constain.Manage_Transaction
	queryRolePermission := model.RolePermission{
		RoleId:     foundRole.Id,
		Permission: permission,
	}
	foundRolePermission := model.RolePermission{}
	if err := db.Where("permission = ?", queryRolePermission.Permission).Where("role_id = ?", queryRolePermission.RoleId).First(&foundRolePermission).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}

	return c.Next()
}

func HasManageTax(c *fiber.Ctx) error {
	db := database.DB
	locals := c.Locals("user").(authmodule.UserDtos)
	queryRole := model.Role{Role: locals.Role}
	foundRole := model.Role{}
	if err := db.Model(queryRole).Where("role = ?", queryRole.Role).First(&foundRole).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}
	permission := constain.Manage_Tax
	queryRolePermission := model.RolePermission{
		RoleId:     foundRole.Id,
		Permission: permission,
	}
	foundRolePermission := model.RolePermission{}
	if err := db.Where("permission = ?", queryRolePermission.Permission).Where("role_id = ?", queryRolePermission.RoleId).First(&foundRolePermission).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}

	return c.Next()
}

func HasManageReport(c *fiber.Ctx) error {
	db := database.DB
	locals := c.Locals("user").(authmodule.UserDtos)
	queryRole := model.Role{Role: locals.Role}
	foundRole := model.Role{}
	if err := db.Model(queryRole).Where("role = ?", queryRole.Role).First(&foundRole).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}
	permission := constain.Manage_Report
	queryRolePermission := model.RolePermission{
		RoleId:     foundRole.Id,
		Permission: permission,
	}
	foundRolePermission := model.RolePermission{}
	if err := db.Where("permission = ?", queryRolePermission.Permission).Where("role_id = ?", queryRolePermission.RoleId).First(&foundRolePermission).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}

	return c.Next()
}

func HasManageProduct(c *fiber.Ctx) error {
	db := database.DB
	locals := c.Locals("user").(authmodule.UserDtos)
	queryRole := model.Role{Role: locals.Role}
	foundRole := model.Role{}
	if err := db.Model(queryRole).Where("role = ?", queryRole.Role).First(&foundRole).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}
	permission := constain.Manage_Product
	queryRolePermission := model.RolePermission{
		RoleId:     foundRole.Id,
		Permission: permission,
	}
	foundRolePermission := model.RolePermission{}
	if err := db.Where("permission = ?", queryRolePermission.Permission).Where("role_id = ?", queryRolePermission.RoleId).First(&foundRolePermission).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}

	return c.Next()
}

func HasManagePayment(c *fiber.Ctx) error {
	db := database.DB
	locals := c.Locals("user").(authmodule.UserDtos)
	queryRole := model.Role{Role: locals.Role}
	foundRole := model.Role{}
	if err := db.Model(queryRole).Where("role = ?", queryRole.Role).First(&foundRole).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}
	permission := constain.Manage_Payment
	queryRolePermission := model.RolePermission{
		RoleId:     foundRole.Id,
		Permission: permission,
	}
	foundRolePermission := model.RolePermission{}
	if err := db.Where("permission = ?", queryRolePermission.Permission).Where("role_id = ?", queryRolePermission.RoleId).First(&foundRolePermission).Error; err != nil {
		return responsepackage.SendJSON(c, "unauthorized", 401, false)
	}

	return c.Next()
}
