package main

import (
	"fmt"

	"go.uber.org/zap"

	"XDCore/src/global"
	"XDCore/src/initial"
)

func main() {
	// 1. 初始化 logger
	initial.InitLogger()

	// 2. 初始化配置文件读取
	initial.InitConfig()

	// 3. 初始化翻译器
	initial.InitTrans("zh")
	initial.InitValidator()

	// 4. 初始化数据库连接
	initial.InitDatabase()

	// 5. 初始化 routers
	router := initial.InitRouters()

	// 启动
	zap.S().Infof("启动服务器, 端口： %d", global.ServerConfig.Port)
	if err := router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err)
	}
}
