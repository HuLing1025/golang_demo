package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang_demo/com/demo/huling/service"
	"golang_demo/com/demo/huling/static"
)

func Init() {
	DbInit()     // 初始化数据库连接
	RouterInit() // 初始化路由信息
}

// 初始化数据库连接
func DbInit() {
	// 获取配置信息,相对项目根目录路径
	configMap := InitConfig("dbConfig.txt")
	// 拼接数据库连接信息
	args := configMap["user"] + ":" + configMap["password"] + "@(" +
		configMap["host"] + ":" + configMap["port"] + ")/" +
		configMap["dbname"] + "?charset=utf8&parseTime=True&loc=Local"
	// 测试输出连接信息
	fmt.Println("连接信息: " + args)
	// 连接数据库
	var err error
	static.DB, err = gorm.Open("mysql", args)
	if err != nil { // 连接失败,打印错误信息
		fmt.Println("failed: 连接数据库失败!" + err.Error())
		return
	} else {
		fmt.Println("success: 成功连接数据库!")
	}
	// 打印sql日志信息
	static.DB.LogMode(true)
	// 全局禁用表名复数
	static.DB.SingularTable(true)
}

// 初始化路由信息
func RouterInit() {
	// 无中间价启动
	//router := gin.New()
	// 默认引擎启动,包含 Logger 和 Recovery 中间件
	router := gin.Default()

	// 路由分组
	group1 := router.Group("/golang")
	{ // 添加路由信息
		// 登录账号
		group1.POST("/login", service.Login)
		// 注册账号
		group1.POST("/logon", service.Logon)
	}
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("failed: 请检查端口占用情况再试!")
		return
	}
}
