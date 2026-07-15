package config

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	PassWord string `mapstructure:"password" json:"password" yaml:"password"` // 数据库密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // 数据库
	PoolSize int    `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize"` // 进程池大小
}
