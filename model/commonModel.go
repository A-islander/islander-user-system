package model

import (
	"log"

	"github.com/user_server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = newDB()

// 要改
func getDsn() string {
	conf := config.GetConfig()
	dsn := conf.UserName + ":" + conf.PassWord + "@tcp(" + conf.Ip + ")/" + conf.Database + "?timeout=30s"
	return dsn
}

// 新建db连接
func newDB() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	return db
}
