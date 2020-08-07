package model

import (
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB connection single unit
var DB *gorm.DB

// InitDatabase init function
func InitDatabase(connString string) {
	db, err := gorm.Open("mysql", connString)
	// 开发调试模式可打开查看 gorm 日志
	//db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(10)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Minute * 30)
	// 设置自动建立表时的字符集
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")
	DB = db

	DB.AutoMigrate(&User{})
}
