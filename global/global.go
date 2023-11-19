package global

import (
	"ginProject/config"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	GVA_CONFIG config.Server
	GVA_LOG    *zap.Logger
	GVA_DB     *gorm.DB
	GVA_VP     *viper.Viper
	GVA_REDIS  *redis.Client
)
