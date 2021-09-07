package model

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//
)

// DB connection single unit
var DB *gorm.DB

// InitDatabase init function
func InitDatabase(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	// 开发调试模式可打开查看 gorm 日志
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	//空闲
	sqlDB.SetMaxIdleConns(10)
	//打开
	sqlDB.SetMaxOpenConns(100)
	//超时
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	// 设置自动建立表时的字符集
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")
	DB = db

	DB.AutoMigrate(&User{})
}
