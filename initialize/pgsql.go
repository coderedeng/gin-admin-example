package initialize

import (
	"fmt"
	"github.com/coderedeng/gin-admin-example/global"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Pgsql() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s %s",
		global.GPA_CONFIG.Pgsql.Host,
		global.GPA_CONFIG.Pgsql.Port,
		global.GPA_CONFIG.Pgsql.UserName,
		global.GPA_CONFIG.Pgsql.DbName,
		global.GPA_CONFIG.Pgsql.PassWord,
		global.GPA_CONFIG.Pgsql.Config,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("连接PgSQL报错： %s \n", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.GPA_CONFIG.Pgsql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.GPA_CONFIG.Pgsql.MaxOpenConns)

	//defer db.Close()
	return db
}
