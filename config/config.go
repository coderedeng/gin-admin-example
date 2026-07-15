package config

type Server struct {
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Pgsql   Pgsql   `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
