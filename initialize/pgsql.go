package initialize

import (
	"fmt"
	"ginProject/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
)

func Pgsql() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s %s",
		global.GVA_CONFIG.Pgsql.Host,
		global.GVA_CONFIG.Pgsql.Port,
		global.GVA_CONFIG.Pgsql.UserName,
		global.GVA_CONFIG.Pgsql.DbName,
		global.GVA_CONFIG.Pgsql.PassWord,
		global.GVA_CONFIG.Pgsql.Config,
	)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		zap.L().Error("连接PgSQL报错：", zap.Error(err))
	}

	sqlDB := db.DB()
	sqlDB.SetMaxIdleConns(global.GVA_CONFIG.Pgsql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.GVA_CONFIG.Pgsql.MaxOpenConns)

	defer db.Close()
	return db
}
