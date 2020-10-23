package dao

import (
	"golang_demo/com/demo/huling/bean"
	"golang_demo/com/demo/huling/static"
)

func SelectRoleByName(role bean.DbRole) bean.DbRole {
	// 查询角色
	var roleResult bean.DbRole
	_ = static.DB.Table("db_role").
		Select("*").
		Where("role=?", role.Role).
		Scan(&roleResult).
		Error
	return roleResult
}
func SelectRoleById(role bean.DbRole) bean.DbRole {
	// 查询角色
	var roleResult bean.DbRole
	_ = static.DB.Table("db_role").
		Select("*").
		Where("id=?", role.ID).
		Scan(&roleResult).
		Error
	return roleResult
}

func SelectRoleList() []bean.DbRole {
	// 查询角色列表
	var roleListResult []bean.DbRole
	_ = static.DB.Table("db_role").
		Select("*").
		Scan(&roleListResult).
		Error
	return roleListResult
}

func InsertRole(role bean.DbRole) bool {
	// 插入
	err := static.DB.Table("db_role").
		Create(&role).
		Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func DeleteRole(role bean.DbRole) bool {
	// 删除
	err := static.DB.Table("db_role").
		Delete(&role).
		Error
	if err != nil {
		return false
	} else {
		return true
	}
}
