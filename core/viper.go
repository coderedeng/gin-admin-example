package core

import (
	"fmt"
	"github.com/coderedeng/gin-admin-example/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config.yaml") // 指定配置文件路径
	err := v.ReadInConfig()          // 读取配置信息
	if err != nil {                  // 读取配置信息失败
		zap.L().Error("初始化配置文件报错：", zap.Error(err))
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	v.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("配置文件已修改：,%s", in.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
