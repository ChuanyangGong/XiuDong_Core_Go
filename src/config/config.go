package config

// 数据库配置
type DatabaseConfig struct {
	IP       string `mapstructure:"ip"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// 服务器配置

type ServerConfig struct {
	DatabaseCfg DatabaseConfig `mapstructure:"database"`
}
