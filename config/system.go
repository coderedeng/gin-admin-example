package config

type System struct {
	Name          string `mapstructure:"name" json:"name" yaml:"name"`                               // 环境值
	Port          int    `mapstructure:"port" json:"port" yaml:"port"`                               // 端口值
	Mode          string `mapstructure:"mode" json:"mode" yaml:"mode"`                               // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	RouterPrefix  string `mapstructure:"routerPrefix" json:"routerPrefix" yaml:"routerPrefix"`       // 路由分组前缀
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
}
