package global

import (
	"github.com/coderedeng/gin-admin-example/config"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GPA_CONFIG              config.Server
	GPA_LOG                 *zap.Logger
	GPA_DB                  *gorm.DB
	GPA_VP                  *viper.Viper
	GPA_REDIS               *redis.Client
	GPA_Concurrency_Control = &singleflight.Group{}
	BlackCache              local_cache.Cache
)
