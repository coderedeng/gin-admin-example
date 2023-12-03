package initialize

import (
	"context"
	"fmt"
	"ginProject/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	Addr := fmt.Sprintf("%s:%d",
		global.GVA_CONFIG.Redis.Host,
		global.GVA_CONFIG.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: global.GVA_CONFIG.Redis.PassWord, // no password set
		DB:       global.GVA_CONFIG.Redis.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	}
	return client
}
