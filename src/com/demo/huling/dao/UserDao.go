package dao

import (
	"golang_demo/src/com/demo/huling/bean"
	"golang_demo/src/com/demo/huling/static"
)

func SelectUser(user bean.DbUser) bean.DbUser {
	var userResult bean.DbUser
	// 查询user
	_ = static.DB.Table("db_user").
		Select("*").
		Where("username=?", user.Username).
		Scan(&userResult).
		Error
	return userResult
}

func SelectUserList() []bean.DbUser {
	var userListResult []bean.DbUser
	// 查询user列表
	_ = static.DB.Table("db_user u,db_role r").
		Select("*").
		Where("u.role_id=r.id").
		Scan(&userListResult).
		Error
	return userListResult
}

func InsertUser(user bean.DbUser) bool {
	// 插入
	err := static.DB.Table("db_user").
		Create(&user).
		Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func UpdateUser(user bean.DbUser) bool {
	// 修改
	err := static.DB.Table("db_user").
		Where("username=?", user.Username).
		Update(user).
		Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func DeleteUser(user bean.DbUser) bool {
	// 删除
	err := static.DB.Delete(&user).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
