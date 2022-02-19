package db

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"management-platform/config"
	"time"
)

var sqlDB *sql.DB
var Conn *gorm.DB

func InitMysqlConn() {
	dsn := config.Config.Mysql.Dsn
	//var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err = db.DB()
	if err != nil {
		panic(err)
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.Config.Mysql.MaxIdleConn)

	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxOpenConn)

	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour * config.Config.Mysql.ConnMaxLifetime)
	Conn = db.Debug()
}
