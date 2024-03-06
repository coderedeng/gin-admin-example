package initialize

import (
	"context"
	"fmt"
	"github.com/coderedeng/gin-admin-example/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	Addr := fmt.Sprintf("%s:%d",
		global.GPA_CONFIG.Redis.Host,
		global.GPA_CONFIG.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: global.GPA_CONFIG.Redis.PassWord, // no password set
		DB:       global.GPA_CONFIG.Redis.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GPA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GPA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
	}
	return client
}
