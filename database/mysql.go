package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"management-platform/config"
	"time"
)

type IMysql interface {
	Connect() error
	DB() *gorm.DB
}
type Mysql struct {
	Conn *gorm.DB
}

func (m *Mysql) Connect() error {
	var sqlDB *sql.DB
	dsn := config.Config.Mysql.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err = db.DB()
	if err != nil {
		return err
	}
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.Config.Mysql.MaxIdleConn)

	// 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxOpenConn)

	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour * config.Config.Mysql.ConnMaxLifetime)
	m.Conn = db.Debug()
	return nil
}
func (m Mysql) DB() *gorm.DB {
	return m.Conn
}
