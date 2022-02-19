package config

import "time"

var Config = new(GlobalConfig)

type GlobalConfig struct {
	Server  ServerConfig
	Mysql   MysqlConfig
	Redis   RedisConfig
	Des     DesConfig
	Version string
}
type ServerConfig struct {
	Port int
}
type RedisConfig struct {
	Database int
	Host     string
	Port     int
}
type MysqlConfig struct {
	Dsn             string
	MaxIdleConn     int
	MaxOpenConn     int
	ConnMaxLifetime time.Duration
}

type DesConfig struct {
	Key string
}
