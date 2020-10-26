package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang_demo/src/com/demo/huling/bean"
	"golang_demo/src/com/demo/huling/dao"
	"net/http"
)

func Login(c *gin.Context) {
	// 绑定前端数据 JSON 数据,单次绑定
	var user bean.DbUser
	_ = c.ShouldBindJSON(&user)
	// 查询用户
	var userResult bean.DbUser
	userResult = dao.SelectUser(user)
	if userResult.Username == "" { // 查询结果为空
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "没有该用户!",
			"data":    nil,
		})
	} else {
		// 匹配密码
		if userResult.Password != "" && userResult.Password == user.Password {
			// 查询角色
			var roleResult bean.DbRole
			roleResult = dao.SelectRoleById(bean.DbRole{ID: userResult.RoleId})
			// 给返回结果装配角色
			userResult.Role = roleResult
			// 定义返回数据类型:map
			var data = make(map[string]bean.DbUser)
			data["user"] = userResult
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "登录成功!",
				"data":    data,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "密码错误!",
				"data":    nil,
			})
		}
	}
}

func Logon(c *gin.Context) {
	// 绑定前端数据 JSON 数据
	// 多次绑定参数时,使用ShouldBindBodyWith方法,
	// ShouldBindJson绑定一次后c.request.body失效
	var user bean.DbUser
	_ = c.ShouldBindBodyWith(&user, binding.JSON)
	var role bean.DbRole
	_ = c.ShouldBindBodyWith(&role, binding.JSON)
	// 用户名是否已经注册
	var userQuery bean.DbUser
	userQuery = dao.SelectUser(bean.DbUser{Username: user.Username})
	if userQuery.Username == "" { // 未查询到该用户,未被注册
		// 查询是否存在该角色
		var roleQuery bean.DbRole
		roleQuery = dao.SelectRoleByName(bean.DbRole{Role: role.Role})
		if roleQuery.Role == "" { // 未查询到该角色
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "查询角色失败!",
				"data":    nil,
			})
			return
		}
		// 设置role_id
		user.RoleId = roleQuery.ID
		flag := dao.InsertUser(user)
		if flag == false {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  500,
				"message": "注册失败!",
				"data":    nil,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "注册成功!",
				"data":    nil,
			})
			return
		}
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "当前账户名已被注册!",
			"data":    nil,
		})
		return
	}
}
