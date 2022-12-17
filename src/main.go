package main

import (
	"XDCore/src/initial"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	port := 8021

	// 1. 初始化 logger
	initial.InitLogger()

	// 2. 初始化配置文件读取
	initial.InitConfig()

	// 3. 初始化 routers
	router := initial.InitRouters()

	// 启动
	zap.S().Infof("启动服务器, 端口： %d", port)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		zap.S().Panic("启动失败", err)
	}
}
