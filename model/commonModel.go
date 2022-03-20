package model

import (
	"log"
	"time"

	"github.com/user_server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = newDB()

// 要改
func getDsn() string {
	conf := config.GetConfig()
	dsn := conf.UserName + ":" + conf.PassWord + "@tcp(" + conf.Ip + ")/" + conf.Database + "?charset=utf8mb4&timeout=30s"
	return dsn
}

// 新建db连接
func newDB() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(50)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db
}
