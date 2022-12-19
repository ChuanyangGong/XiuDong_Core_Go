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
	Port        int            `mapstructure:"port"`
	JwtConfig   JWTConfig      `mapstructure:"jwtConfig"`
}

// jwt config
type JWTConfig struct {
	SigningKey string `mapstructure:"signingKey"`
}
