package initial

import (
	"os"
	"path"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"XDCore/src/global"
)

func InitConfig() {
	cfgFilePath := "./src/config-dev.yaml"

	zap.S().Info("初始化配置文件信息")
	v := viper.New()
	v.SetConfigFile(cfgFilePath)
	if err := v.ReadInConfig(); err != nil {
		wd, _ := os.Getwd()
		zap.S().Panicf("读取配置文件出错，检查文件路径 %s 是否存在", path.Join(wd, cfgFilePath))
	}

	if err := v.Unmarshal(global.ServerConfig); err != nil {
		zap.S().Panic("读取配置文件出错，请检查 server config struct")
	}
	zap.S().Debugf("配置文件读取结果：%v", *global.ServerConfig)
}
