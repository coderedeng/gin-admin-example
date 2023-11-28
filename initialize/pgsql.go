package initialize

import (
	"fmt"
	"ginProject/global"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("连接PgSQL报错： %s \n", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.GVA_CONFIG.Pgsql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.GVA_CONFIG.Pgsql.MaxOpenConns)

	//defer db.Close()
	return db
}
