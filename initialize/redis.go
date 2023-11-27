package initialize

import (
	"fmt"
	"ginProject/global"
	"github.com/go-redis/redis"
)

func Redis() *redis.Client {
	Addr := fmt.Sprintf("%s:%d",
		global.GVA_CONFIG.Redis.Host,
		global.GVA_CONFIG.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: global.GVA_CONFIG.Redis.PassWord, // 密码
		DB:       global.GVA_CONFIG.Redis.DB,       // 数据库
		PoolSize: global.GVA_CONFIG.Redis.PoolSize, // 连接池大小
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic(fmt.Errorf("连接Redis报错: %s \n", err))
	}
	return rdb
}
