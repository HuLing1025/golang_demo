package main

import (
	"fmt"
	"golang_demo/com/demo/huling/static"
	"golang_demo/com/demo/huling/utils"
)

func main() {
	fmt.Println("System start!")
	// 初始化
	utils.Init()
	// 程序结束后释放连接
	defer static.DB.Close()
}
